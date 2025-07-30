package main

import (
	"encoding/json"
	"fmt"
	helloworld "learn/rpc/hello_rpc/proto"

	"google.golang.org/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := helloworld.HelloRequest{
		Name:    "dawn",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	result, _ := proto.Marshal(&req) //结构体转protoc
	fmt.Println(string(result), "长度:", len(result))

	jsonStruct := Hello{Name: "dawn", Age: 18, Courses: []string{"go", "gin", "微服务"}}
	jsondata, _ := json.Marshal(jsonStruct) //结构体转json
	fmt.Println(string(jsondata), "长度:", len(jsondata))

	newreq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(result, &newreq)
	fmt.Println("反序列化后：", newreq.Name, newreq.Age, newreq.Courses)

}
