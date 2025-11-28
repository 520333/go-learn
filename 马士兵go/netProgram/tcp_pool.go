package netProgram

import (
	"errors"
	"net"
	"sync"
	"time"
)

type Pool interface {
	Get() (net.Conn, error)  //获取池子中空闲连接
	Put(conn net.Conn) error //放回连接，非关闭
	Release() error          //释放连接池
	Len() int                //有效连接的长度
}

// PoolConfig 连接池配置
type PoolConfig struct {
	MinConnNum  int           //最小连接数，至少保持多少个有效连接
	MaxConnNum  int           //最大连接数，池中最多支持多少连接
	MaxIdleNum  int           //空闲连接数，池中最多有多少可用的连接
	IdleTimeout time.Duration //空闲连接超时时间，多久后空闲连接后会被释放
	Factory     ConnFactory   //连接工厂

}

// IdleConn 空闲连接类型
type IdleConn struct {
	conn    net.Conn  //连接本身
	putTime time.Time //放回池子时间，用于判断是否空闲超时
}

// TcpPool 连接池结构
type TcpPool struct {
	config        *PoolConfig    //配置信息
	openingConNum int            //开放使用的连接数量
	idleList      chan *IdleConn //空闲的连接列表
	mu            sync.RWMutex   //并发安全锁
}

// ConnFactory 连接工厂接口
type ConnFactory interface {
	Factory(addr string) (net.Conn, error) //创建生产连接
	Close(net.Conn) error                  //关闭连接
	Ping(net.Conn) error                   //Ping测试
}

// TcpConnFactory TCP连接工厂类型
type TcpConnFactory struct{}

func (f *TcpConnFactory) Factory(addr string) (net.Conn, error) {
	//校验参数的合理性
	if addr == "" {
		return nil, errors.New("addr is empty")
	}
	//建立连接
	conn, err := net.DialTimeout("tcp", addr, time.Second*5)
	if err != nil {
		return nil, err
	}
	//返回
	return conn, nil
}
func (f *TcpConnFactory) Close(conn net.Conn) error {
	return conn.Close()
}
func (f *TcpConnFactory) Ping(conn net.Conn) error {
	return nil
}

// Get TcpPool实现 Pool 接口
func (*TcpPool) Get(net.Conn) (net.Conn, error) {
	return nil, nil
}
func (*TcpPool) Put(conn net.Conn) error {
	return nil
}
func (f *TcpPool) Release() error {
	return nil
}
func (f *TcpPool) Len() int {
	return 0
}
