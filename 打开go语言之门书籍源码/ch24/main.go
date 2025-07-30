package main

import (
	"fmt"
)

func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("%d:%d", i, i)
}
func stringify[T fmt.Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

func main() {
	print[int]([]int{1, 2, 3})
	// 实例化
	p := print[interface{}]
	p([]interface{}{"as", 1, true})

	p1 := print[int]
	p1([]int{1, 2, 3})

	fmt.Println(stringify([]MyInt{1, 2, 3}))

	fmt.Println(4 << 1)
}
