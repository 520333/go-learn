package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Printf("现在时间：%d年%v月%d日 %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	prev := time.Date(2026, 5, 22, 15, 20, 20, 12340, time.UTC)
	fmt.Println(prev)
	fmt.Println(prev.Before(now))
	fmt.Println(prev.After(now))
	fmt.Println(prev.Equal(now))
	fmt.Println(prev.Sub(now))
}
func randomDemo() {
	// 0-100随机数
	fmt.Println(rand.Intn(100))
}

type BaseResponse struct {
	Code int          `json:"code"`
	Data ResponseData `json:"data"`
}
type ResponseData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Msg  string `json:"msg"`
}

func jsonDemo() {
	br := BaseResponse{
		Code: 1,
		Data: ResponseData{
			Name: "宝哥",
			Age:  19,
			Msg:  "你做的很好下次别做了!!!",
		},
	}
	jsonByte, _ := json.Marshal(&br)
	fmt.Println(string(jsonByte))

	var br2 BaseResponse
	_ = json.Unmarshal(jsonByte, &br2)
	fmt.Println(br2)
}

func regexDemo() {
	input := "My email is dawn@g8s.me xxx@g8s.me 123123@g8s.me"
	// exp, _ := regexp.Compile("@g8s.me")
	// fmt.Println(exp.FindString(input))

	exp, _ := regexp.Compile("[a-zA-Z0-9]+@[a-zA-Z0-9]+.[a-zA-Z0-9]+")
	fmt.Println(exp.FindAllString(input, -1))
	for _, subMatch := range exp.FindAllStringSubmatch(input, -1) {
		fmt.Println(subMatch[0])

	}

}
func init() {
	// rand.Seed(time.Now().Unix())
	fmt.Println(rand.New(rand.NewSource(time.Now().Unix())))
}
func main() {
	timeDemo()
	fmt.Println("=========================")
	randomDemo()
	jsonDemo()
	regexDemo()
}
