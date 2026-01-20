package proxy

import (
	"context"
	"net"
	"time"
)

type TCPReverseProxy struct {
	Addr            string
	DialTimeout     time.Duration // 拨号超时时间 持续时间
	DialLine        time.Duration // 拨号截止时间 截止日志
	KeepAlivePeriod time.Duration // 长连接超时时间
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

}
