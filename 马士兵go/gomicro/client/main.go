package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type User struct {
	Name, Pwd  string
	Age, Score int
}
type Goods struct{}
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

func (user *User) GetScore(u User, resp *int) error {
	*resp = u.Age * u.Score
	return nil
}

func main() {
	//conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//resp := 0
	//conn.Call("User.GetScore", User{"xx", "oo", 2, 60}, &resp)
	var respGoods AddGoodsResp
	var reqGoods = AddGoodsReq{Name: "测试商品1"}
	conn.Call("Goods.AddGoods", reqGoods, &respGoods)
	fmt.Println("远程调用查询的返回：", respGoods)

	var queryGoodsReq QueryGoodsReq
	var queryGoodsResp = QueryGoodsResp{}
	conn.Call("Goods.QueryGoods", queryGoodsReq, &queryGoodsResp)
	fmt.Println("远程调用查询的返回：", queryGoodsResp)
}
