package main

import (
	"errors"
	"fmt"
)

// 实现除法
func division(i1, i2 float64) (res float64, err error) {
	fmt.Println("需要计算的数字是:", i1, i2)
	if i2 == 0 {
		return 0, errors.New("输入的分母不能为0")
	} else {
		res = i1 / i2
		return res, nil
	}
}

// myfunc main() {
// 	// f, err := ioutil.ReadFile("./text.txt")
// 	// if err != nil {
// 	// 	fmt.Println("读取文件失败:", err.Error())
// 	// } else {
// 	// 	fmt.Println(string(f))
// 	// }
// 	//自定义ERR
// 	err := errors.New("自定义错误")
// 	fmt.Println(err) //stdout: 自定义错误
// 	err2 := fmt.Errorf("这是一个自定义的错误: %s,它是使用fmt生成的", "这是错误内容")
// 	fmt.Println(err2)                            //stdout: 这是一个自定义的错误: 这是错误内容,它是使用fmt生成的
// 	fmt.Println("这是一个使用fmt定义的错误:", err2.Error()) // stdout: 这是一个使用fmt定义的错误: 这是一个自定义的错误: 这是错误内容,它是使用fmt生成的

// 	res, err3 := division(2, 0)
// 	if err3 != nil {
// 		fmt.Println("计算错误", err3.Error())
// 	} else {
// 		fmt.Println("计算结果:", res)
// 	}
// }
