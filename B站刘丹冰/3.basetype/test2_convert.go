package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var a int8 = 12
	// var b uint8 = uint8(a)
	// var c float32 = 3.14
	// d := float64(c)
	// var f64 = float64(a)
	// fmt.Println(b, d, f64)

	type IT int
	var B IT = IT(a)
	fmt.Println(B)

	var istr = "12"
	d, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert error", err)
	}
	fmt.Println(d)

	var myi = 12
	fmt.Println(strconv.Itoa(myi), reflect.TypeOf(myi), reflect.TypeOf(strconv.Itoa(myi)))

	// 字符串转bool类型
	float, err := strconv.ParseFloat("3.1415926", 64)
	if err != nil {
		return
	}
	fmt.Println(float)
	// 进制转换: 要转换的字符,进制,数据大小 【将8进制100转换成10进制】
	parsetInt, err := strconv.ParseInt("100", 8, 64)
	if err != nil {
		return
	}
	fmt.Println(parsetInt)
	ParseBool, err := strconv.ParseBool("0")
	if err != nil {
		fmt.Println("ParseBool error")
		return
	}
	fmt.Println(ParseBool)

	// 基本类型bool转字符串
	boolstr := strconv.FormatBool(true)
	fmt.Printf("boolstr=%T  ,%s\r\n", boolstr, boolstr)

	// 基本类型float转字符串 FormatFloat(浮点数,'格式化类型',保留几位自动四舍五入,4字节长度)
	floatstr := strconv.FormatFloat(3.18159, 'f', 1, 64)
	fmt.Println(floatstr)

	// 基本类型转int 将100转换成8进制
	fmt.Println(strconv.FormatInt(100, 8))

}
