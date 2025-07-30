package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	var array = [...]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(array[2])
	array1 := [5]string{1: "b", 3: "d"}
	fmt.Println(array1)
	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引：%d,对应值：%s\n", i, array[i])
	}
	for i, v := range array {
		fmt.Printf("数组索引：%d,对应值：%s\n", i, v)
	}
	slice := array[2:5]
	fmt.Println(slice, reflect.TypeOf(slice))
	slice[1] = "f"
	fmt.Println(slice)
	// var slice1 = make([]string, 4)
	var slice1 = []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice1), cap(slice1))
	slice2 := append(slice1, "f", "g")
	fmt.Println(slice2)
	slice3 := append(slice1, slice...)
	for i, k := range slice3 {
		fmt.Println(i, k)
	}

	nameAgeMap := make(map[string]int)
	nameAgeMap["阿宝"] = 20
	// nameAgeMap = map[string]int{"宝哥": 21}
	fmt.Println(nameAgeMap["阿宝"])
	age, ok := nameAgeMap["阿宝1"]
	if !ok {
		fmt.Println(age)
	}
	delete(nameAgeMap, "阿宝")
	nameAgeMap["阿宝1"] = 21
	nameAgeMap["阿宝2"] = 22
	for k, v := range nameAgeMap {
		fmt.Println("key is ", k, "value is ", v)
	}
	fmt.Println(len(nameAgeMap))
	var s string = "hello 宝哥发大财"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[15])
	fmt.Println(utf8.RuneCountInString(s))
	for i, r := range s {
		fmt.Println(i, r)
	}
	var array2 [2][]string
	array2[0] = []string{"gin", "gorm"}
	array2[1] = []string{"grpc", "gorouter"}
	// array2[1] = [2]string{"3", "4"}
	// array2[2] = [2]string{"3", "4"}
	for i := 0; i < len(array2); i++ {
		for j := 0; j < len(array2[i]); j++ {
			fmt.Println(array2[i][j])
		}
	}

	slice4 := [2][]string{}
	slice4[0] = array[2:5]
	fmt.Println(slice4)

	// var two [1][]byte
	// two[0] = []byte(s)
	// fmt.Println(two)
	// for i := 0; i < len(two); i++ {
	// 	for j := 0; j < len(two[i]); j++ {
	// 		fmt.Println(two[i][j])
	// 	}
	// }
	aa := [3][3]int{}
	aa[0][0] = 1
	aa[0][1] = 2
	aa[0][2] = 3
	aa[1][0] = 1
	aa[1][1] = 2
	aa[1][2] = 3
	aa[2][0] = 1
	aa[2][1] = 2
	aa[2][2] = 3
	for k, v := range aa {
		fmt.Println(k, v)
	}
}
