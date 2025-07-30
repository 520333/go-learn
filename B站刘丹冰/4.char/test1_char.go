package main

/*
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 长度计算
	name := "china golang工程师一日成魔"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	// 转义符 \n换行  \r回车  \t tab制表符  \\反斜线  \'单引号  \"双引号  \?问号

	courseName1 := "go\"工程师\" \\ \r\n"
	courseName2 := `go"工程师"`
	fmt.Println(courseName1, courseName2)
	fmt.Print("hello\t")
	fmt.Print("world\n")

	// 格式化输出
	username := "dawn"
	age := 18
	address := "Fujian"
	mobile := "18112341234"
	fmt.Println("用户名："+username, ",年龄："+strconv.Itoa(age), ",地址："+address, ",手机："+mobile) //性能高
	fmt.Printf("用户名：%s, 年龄：%d, 地址：%s, 手机：%s\r\n", username, age, address, mobile)         //性能一般

	userMsg := fmt.Sprintf("用户名：%s, 年龄：%d, 地址：%s, 手机：%s\r\n", username, age, address, mobile)
	fmt.Println(userMsg)
	// Builder相比Sprintf和Println性能更高
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString(username)
	builder.WriteString("年龄：")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString("地址：")
	builder.WriteString(address)
	builder.WriteString("手机：")
	builder.WriteString(mobile)
	re := builder.String()
	fmt.Println(re)

}
*/
