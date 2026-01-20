package proxy

import (
	"context"
	"io"
	"log"
	"net"
	"time"
)

type TCPReverseProxy struct {
	Addr            string
	DialTimeout     time.Duration                                                     // 拨号超时时间 持续时间
	DialLine        time.Duration                                                     // 拨号截止时间 截止日志
	KeepAlivePeriod time.Duration                                                     // 长连接超时时间
	DialContext     func(ctx context.Context, network, addr string) (net.Conn, error) // 拨号器
	ModifyResponse  func(net.Conn) error
	ErrorHandler    func(net.Conn, error)
}

func NewSingleHostReverseProxy(addr string) *TCPReverseProxy {
	if addr == "" {
		panic("TCP ADDRESS must not be empty!")
	}
	return &TCPReverseProxy{
		Addr:            addr,
		DialTimeout:     10 * time.Second,
		DialLine:        time.Minute,
		KeepAlivePeriod: time.Hour,
	}
}

func (p *TCPReverseProxy) ServeTCP(ctx context.Context, src net.Conn) {
	defer src.Close()
	var cancel context.CancelFunc
	if p.DialTimeout >= 0 {
		ctx, cancel = context.WithTimeout(ctx, p.DialTimeout)
	}
	if p.DialLine >= 0 {
		ctx, cancel = context.WithDeadline(ctx, time.Now().Add(p.DialTimeout))
	}
	if cancel != nil {
		defer cancel()
	}

	if p.DialContext == nil {
		p.DialContext = (&net.Dialer{
			Timeout:   p.DialTimeout,
			Deadline:  time.Now().Add(p.DialLine),
			KeepAlive: p.KeepAlivePeriod,
		}).DialContext
		dst, err := p.DialContext(ctx, "tcp", p.Addr)
		if err != nil {
			p.getErrorHandler()(src, err)
			//src.Close()
			return
		}
		defer dst.Close()
		if !p.modifyResponse(dst) {
			return
		}
		_, err = bytesCopy(src, dst)
		if err != nil {
			p.getErrorHandler()(dst, err)
			dst.Close()
		}
	}

}

func (p *TCPReverseProxy) getErrorHandler() func(net.Conn, error) {
	if p.ErrorHandler == nil {
		return p.defaultErrorHandler
	}
	return p.ErrorHandler
}

func (p *TCPReverseProxy) defaultErrorHandler(conn net.Conn, err error) {
	log.Printf("TCP Proxy: for connection %v error: %v\n", conn.RemoteAddr().String(), err)
}

func (p *TCPReverseProxy) modifyResponse(res net.Conn) bool {
	if p.ModifyResponse == nil {
		return true
	}
	if err := p.ModifyResponse(res); err != nil {
		res.Close()
		p.getErrorHandler()(res, err)
		return false
	}
	return true
}

func bytesCopy(dst, src net.Conn) (len int64, err error) {
	len, err = io.Copy(dst, src)
	return
}
