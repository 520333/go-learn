package main

import (
	"fmt"
	"os"
)

func main() {
	srcPath := "../demo04/test.txt"
	desPath := "test_des.txt"
	context, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(desPath, context, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
