package testutils

import "fmt"

var Age int
var Sex string
var Name string

func init() {
	fmt.Println("testutils init function execution.")
	Age = 19
	Sex = "男"
	Name = "李四"
}
