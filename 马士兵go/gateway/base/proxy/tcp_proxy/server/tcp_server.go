package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 1.监听端口
// 2.获取连接
// 3.封装连接对象,设置服务参数
// 4.回调handler(定义接口)

// TCPServer TCP服务核心结构体。监听指定主机，并提供服务
// Addr    必选 主机地址
// Handler 必选 回调函数 处理TCP请求。提供默认实现
type TCPServer struct {
	Addr             string             // 主机地址
	Handler          TCPHandler         // 回调函数 处理TCP请求
	BaseCtx          context.Context    // 上下文，收集取消 终止 错误等信息
	err              error              // TCP Error
	ReadTimeout      time.Duration      // 读超时
	WriteTimeout     time.Duration      // 写超时
	KeepAliveTimeout time.Duration      // 长连接超时
	mu               sync.Mutex         // 连接关闭等关键动作 需要加锁
	doneChan         chan struct{}      // 服务已完成,监听系统信号
	inShutdown       int32              // 服务终止: 0-未关闭 1-已关闭
	l                *onceCloseListener // 服务器监听器 使用完成都要进行关闭
}

type TCPHandler interface {
	ServeTCP(ctx context.Context, conn net.Conn)
}

type tcpHandler struct{}

func (h *tcpHandler) ServeTCP(ctx context.Context, conn net.Conn) {
	conn.Write([]byte("Pong! TCP handler here.\n"))
}

var (
	ErrServerClosed     = errors.New("tcp: server is closed")
	ErrAbortHandler     = errors.New("net/tcp: abort handler")
	ServerContextKey    = &contextKey{"tcp-server"}
	LocalAddrContextKey = &contextKey{"local-addr"}
)

func (ts *TCPServer) ListenAndServe() error {
	if ts.shuttingDown() {
		return ErrServerClosed
	}
	addr := ts.Addr
	if addr == "" {
		return errors.New("tcp server addr is empty")
	}
	if ts.Handler == nil {
		ts.Handler = &tcpHandler{}
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return ts.Serve(ln)
}

func ListenAndServe(addr string, handler TCPHandler) error {
	server := &TCPServer{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

func (ts *TCPServer) Serve(l net.Listener) error {
	ts.l = &onceCloseListener{Listener: l}
	defer l.Close()
	if ts.BaseCtx == nil {
		ts.BaseCtx = context.Background()
	}
	baseCtx := ts.BaseCtx
	ctx := context.WithValue(baseCtx, ServerContextKey, ts)
	for {
		rw, err := l.Accept()
		if err != nil {
			select {
			case <-ts.getDoneChan():
				return ErrServerClosed
			default:
			}
			return err
		}
		c := ts.newConn(rw) // 对TCPConn 二次封装
		go c.serve(ctx)
	}
}

func (ts *TCPServer) newConn(rwc net.Conn) *conn {
	c := &conn{
		server:     ts,
		rwc:        rwc,
		remoteAddr: rwc.RemoteAddr().String(),
	}
	if t := ts.ReadTimeout; t != 0 {
		c.rwc.SetReadDeadline(time.Now().Add(t))
	}
	if t := ts.WriteTimeout; t != 0 {
		c.rwc.SetWriteDeadline(time.Now().Add(t))
	}
	if t := ts.KeepAliveTimeout; t != 0 {
		if tcpConn, ok := c.rwc.(*net.TCPConn); ok {
			tcpConn.SetKeepAlive(true)
			tcpConn.SetKeepAlivePeriod(t)
		}
	}
	return c
}

func (c *conn) serve(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil && err != ErrAbortHandler {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("tcp panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
		c.rwc.Close()
	}()
	ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())
	if c.server.Handler == nil {
		panic("tcp handler empty!")
	}
	c.server.Handler.ServeTCP(ctx, c.rwc)
}

type conn struct {
	server     *TCPServer
	rwc        net.Conn
	remoteAddr string
}
type onceCloseListener struct {
	net.Listener
	once     sync.Once
	closeErr error
}

func (oc *onceCloseListener) Close() error {
	oc.once.Do(oc.close)
	return oc.closeErr
}

func (oc *onceCloseListener) close() {
	oc.closeErr = oc.Listener.Close()
}

// Close TCPServer关闭功能
func (ts *TCPServer) Close() error {
	atomic.StoreInt32(&ts.inShutdown, 1) // 原子操作修改服务器状态字段： 1-关闭
	close(ts.doneChan)                   // 关闭channel
	ts.l.Close()                         // 关闭监听:listener
	return nil
}

// 检查当前服务器是否已关闭
func (ts *TCPServer) shuttingDown() bool {
	return atomic.LoadInt32(&ts.inShutdown) != 0
}

type contextKey struct {
	name string
}

func (ts *TCPServer) getDoneChan() <-chan struct{} {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	if ts.doneChan == nil {
		ts.doneChan = make(chan struct{})
	}
	return ts.doneChan
}
