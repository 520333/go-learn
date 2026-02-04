package main

import (
	"context"
	"flag"
	"gateway/base/proxy/grpc_proxy/proto"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	var port = flag.Int("port", 8085, "the port to server on")
	flag.Parse()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// UnknownServiceHandler 返回一个 ServerOption 允许添加自定义的未知服务处理程序
	servOption := grpc.UnknownServiceHandler(handler)
	s := grpc.NewServer(servOption)
	_ = s.Serve(listener)
}

// 0.过滤非RPC请求
// 1.构建下游连接器
// 创建下游连接：往下游真实服务器创建连接
// 封装下游客户端流实例
// 2.上游与下游数据拷贝
// 3.关闭双向流
func handler(srv interface{}, pxyServerStream grpc.ServerStream) error {
	// 0.过滤非RPC请求  /service/method
	methodName, ok := grpc.MethodFromServerStream(pxyServerStream)
	if !ok {
		return status.Errorf(codes.Internal, "There is no PRC-Request in this context")
	}
	// 不处理内部请求
	if strings.HasPrefix(methodName, "/com.example.internal") {
		return status.Errorf(codes.Unimplemented, "Unimplemented method")
	}

	// 1.构建下游连接器
	ctx := pxyServerStream.Context()
	// 负载均衡器算法获取下游服务器地址
	pxyClientConn, err := grpc.DialContext(ctx, "localhost:8005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer pxyClientConn.Close()

	// 从上游请求上下文中获取元数据
	md, _ := metadata.FromIncomingContext(ctx)
	outCtx, clientCancel := context.WithCancel(ctx)
	outCtx = metadata.NewOutgoingContext(outCtx, md)
	// 1.2 封装下游客户端流实例
	pxyStreamDesc := &grpc.StreamDesc{
		ServerStreams: true,
		ClientStreams: true,
	}
	pxyClientStream, err := grpc.NewClientStream(outCtx, pxyStreamDesc, pxyClientConn, methodName)
	if err != nil {
		return err
	}

	// 2.上游与校友数据拷贝
	// 把上游服务器消息 发送给下游真实服务器
	s2cErrChan := serverToClient(pxyClientStream, pxyServerStream)
	// 把下游响应消息 发回上游客户端
	c2sErrChan := clientToServer(pxyServerStream, pxyClientStream)

	// 3.关闭流
	for i := 0; i < 2; i++ {
		select {
		case s2cErr := <-s2cErrChan: // 向下游发消息
			if s2cErr == io.EOF {
				pxyClientStream.CloseSend()
			} else {
				if clientCancel != nil {
					clientCancel()
				}
				return status.Errorf(codes.Internal, "failed proxying server to client:%v", s2cErr)
			}
		case c2sErr := <-c2sErrChan: // 往上游回写消息
			pxyServerStream.SetTrailer(pxyClientStream.Trailer())
			if c2sErr != io.EOF {
				return c2sErr
			}
			return nil
		}
	}

	return nil
}

func clientToServer(dst grpc.ServerStream, src grpc.ClientStream) chan error {
	res := make(chan error, 1)
	go func() {
		msg := &proto.EchoResponse{}
		//for i := 0; ; i++ {
		//	if i == 0 {
		for {
			// response header 进行处理
			// 客户端读取响应时，会先读取响应头，然后在作出相应的处理
			// 所以有必要设置响应头
			md, err := src.Header()
			if err != nil {
				res <- err
				break
			}
			if err = dst.SendHeader(md); err != nil {
				res <- err
				break
			}

			if err = src.RecvMsg(msg); err != nil {
				res <- err
				break
			}
			if err = dst.SendMsg(msg); err != nil {
				res <- err
				break
			}
			//	}
		}
	}()
	return res
}

func serverToClient(dst grpc.ClientStream, src grpc.ServerStream) chan error {
	res := make(chan error, 1)
	go func() {
		msg := &proto.EchoRequest{}
		for {
			// 客户端请求头，拷贝到下游
			// X-Forward-For
			// clientHeaderToServer(dst grpc.ClientStream, src grpc.ServerStream)
			if err := src.RecvMsg(msg); err != nil {
				res <- err
				break
			}
			if err := dst.SendMsg(msg); err != nil {
				res <- err
				break
			}
		}
	}()
	return res
}
