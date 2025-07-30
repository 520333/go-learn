package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[string]int
	fmt.Println(m, m == nil)

	m = make(map[string]int)
	m["a"] = 10
	fmt.Println(m, m == nil)
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	fmt.Println(m1, m1["e"])
	v, ok := m1["d"]
	if ok {
		fmt.Printf("key e in map and value is %d\n", v)
	} else {
		fmt.Printf("key not in map m1")
	}

	if k, ok := m1["e"]; ok {
		fmt.Printf("key e in map and value is %d\n", k)
	} else {
		fmt.Printf("key not in map m1")
	}
	for k := range m1 {
		fmt.Printf("key=%s val=%d\n", k, m1[k])
	}
	fmt.Println("============")
	for k, v := range m1 {
		fmt.Printf("key=%s val=%d\n", k, v)
	}

	// fmt.Println("============")
	// delete(m1, "b")
	// for k, v := range m1 {
	// 	fmt.Printf("key=%s val=%d\n", k, v)
	// }

	fmt.Println("============")
	var keys []string
	for k := range m1 {
		keys = append(keys, k)
	}

	// 排序迭代
	sort.Strings(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Printf("key=%s,val=%d\n", k, m1[k])
	}
}
