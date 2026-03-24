package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type User struct {
	Name, Pwd  string
	Age, Score int
}

func (user *User) GetScore(u User, resp *int) error {
	*resp = u.Age * u.Score
	return nil
}

type Goods struct{}

// AddGoodsReq 请求结构体
type AddGoodsReq struct {
	Name string
}
type AddGoodsResp struct {
	Success bool
	Message string
}

type QueryGoodsReq struct {
	Id   int
	Name string
}
type QueryGoodsResp struct {
	Success bool
	Goods   Goods
	Message string
}

func (g *Goods) AddGoods(req QueryGoodsReq, resp *QueryGoodsResp) error {
	fmt.Println(req.Name, "GORM 入库操作！！！")
	*resp = QueryGoodsResp{true, Goods{}, "新增商品成功！"}
	return nil
}
func (g *Goods) QueryGoods(req QueryGoodsReq, resp *QueryGoodsResp) error {
	fmt.Println(req.Name, "GORM 入库操作！！！")
	*resp = QueryGoodsResp{true, Goods{}, "查询商品成功！"}
	return nil
}

func main() {
	//user := new(User)
	goods := new(Goods)
	//rpc.Register(user)
	rpc.Register(goods)
	rpc.HandleHTTP()
	//err := http.ListenAndServe(":8081", nil)
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	defer func() {
		_ = listen.Close()
	}()

	for {
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("accept error:", err)
		}
		rpc.ServeConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
