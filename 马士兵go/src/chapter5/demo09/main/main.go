package main

import (
	"chapter5/demo09/cautils"
	"chapter5/demo09/dbutils"
	"fmt"
)

func main() {
	fmt.Println("main functions load... OK")
	var num1, num2 = 10, 20
	dbutils.ExeChange(&num1, &num2)
	fmt.Println(num1, num2)
	dbutils.GetConnection()
	dbutils.Add()
	cautils.Add()
}
