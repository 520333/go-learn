package main

import (
	"fmt"
)

func main() {

}
func StringPlus() string {
	var s string
	s += "昵称" + ":" + "宝哥无情" + "\n"
	s += "博客" + ":" + "https://blog.g8s.me" + "\n"
	s += "微信公众号" + ":" + "g8s_me"
	return s
}

func StringFmt() string {
	return fmt.Sprint("昵称", ":", "宝哥无情", "\n", "博客", ":", "https://blog.g8s.me", "\n", "微信公众号", ":", "g8s_me")
}
