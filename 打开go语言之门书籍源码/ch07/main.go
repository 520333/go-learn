package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{errorMsg: "a或者b不能为负数", errorCode: 1}
	} else {
		return a + b, nil
	}
}

type commonError struct {
	errorCode int //错误码
	errorMsg  string
}

func (ce *commonError) Error() string {
	return ce.errorMsg
}

type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	return e.err.Error() + e.msg
}

func connectMySQL(ip, username, password string) {
	if ip == "" {
		panic("IP不能为空")
	}
}
func moreDefer() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Three defer")
	fmt.Println("函数自身代码")
}
func main() {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	sum, err := add(5, -1)
	var cm *commonError

	if errors.As(err, &cm) {
		// if cm, ok := err.(*commonError); ok {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

	newErr := MyError{err, "数据上传问题"}
	fmt.Println(newErr)
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
	fmt.Println(errors.Is(w, e))

	moreDefer()
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	connectMySQL("", "root", "123456")
}
