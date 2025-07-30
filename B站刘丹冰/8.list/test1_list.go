package main

import (
	"container/list"
	"fmt"
)

func main() {
	var mylist list.List
	// 从尾部插入
	mylist.PushBack("go")
	mylist.PushBack("grpc")
	mylist.PushBack("mysql")
	fmt.Println(mylist)

	// 遍历打印值 正序从头到尾
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("==================")

	// 从头部插入
	mylist.PushFront("gin")

	//
	i := mylist.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "grpc" {
			break
		}
	}
	mylist.InsertBefore("gin", i)
	// 删除元素
	mylist.Remove(i)
	// 遍历打印值 倒叙从后到前
	for i := mylist.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

}
