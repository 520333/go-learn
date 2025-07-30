package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(&fp)
	p := new(person)
	//Name是person的第一个字段不用偏移，即可通过指针修改
	pName := (*string)(unsafe.Pointer(p))
	*pName = "dawn"
	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)

	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(uint(0)))
	fmt.Println(unsafe.Sizeof(uint8(0)))
	fmt.Println(unsafe.Sizeof(uint16(0)))
	fmt.Println(unsafe.Sizeof(uint32(0)))
	fmt.Println(unsafe.Sizeof(uint64(0)))
	fmt.Println(unsafe.Sizeof(string("阿宝狂飙123123")))
	fmt.Println(unsafe.Sizeof(float32(3.14)))
	fmt.Println(unsafe.Sizeof(float64(3398.911)))
	fmt.Println(unsafe.Sizeof([]string{"海绵u无情", "张三"}))
	fmt.Println(unsafe.Sizeof(map[string]int{}))

}

type person struct {
	Name string
	Age  int
}
