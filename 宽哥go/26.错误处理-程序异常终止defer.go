package main

import (
	"errors"
	"fmt"
)

func connectDatabase(address string, port int) (string, error) {
	//如果address和port为空
	if address == "" || port == 0 {
		return "", errors.New("无法连接数据库")
	} else {
		return "连接成功", nil

	}
}

func returnDataToFrontend(msg string) {
	fmt.Println("返回给前端的数据是:", msg)
}

// myfunc main() {
// 	//1.关闭连接池
// 	//2.关闭文件句柄
// 	//3.记录一些异常日志
// 	//defer:是go语言中的一种延迟调用机制，defer里面的内容可以在函数return之前或者是程序panic之前执行
// 	//一般用于资源回收和数据返回，defer也可以用于异常时的恢复
// 	//defer可以有多个，采用先进后出机制
// 	msg := "数据库连接异常"
// 	defer returnDataToFrontend("1")
// 	defer returnDataToFrontend("2")
// 	defer returnDataToFrontend("3")
// 	defer returnDataToFrontend("4")
// 	defer returnDataToFrontend(msg)
// 	_, err := connectDatabase("", 0)
// 	if err != nil {
// 		fmt.Println(err)
// 		msg = "连接失败"
// 		panic(err)
// 	}
// 	//返回数据给前端
// 	// returnDataToFrontend(msg)
// }
