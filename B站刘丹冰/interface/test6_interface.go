package main

import "fmt"

func mPrint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}

type myinfo struct{}

func (mi *myinfo) Error() string {
	return "我不是error"
}
func main() {
	var data1 = []interface{}{"dawn", 18, 1.88}
	mPrint(data1...)

	var data2 = []string{"chuang", "20", "1.72"}
	var datai []interface{}
	for _, value := range data2 {
		datai = append(datai, value)
	}
	mPrint(datai...)

	err := &myinfo{}
	fmt.Println(err)
}
