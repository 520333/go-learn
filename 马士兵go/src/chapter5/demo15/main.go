package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := strings.Replace("golangandjavago", "go", "golang", -1)
	str2 := strings.Replace("golangandjavago", "go", "golang", 1)
	fmt.Println(str1)
	fmt.Println(str2)

	arr := strings.Split("go-python-java", "-")
	fmt.Println(arr)
	lower := strings.ToLower("Go")
	upper := strings.ToUpper("go")
	fmt.Println(lower)
	fmt.Println(upper)
	fmt.Println(strings.TrimSpace(" sa sad   "))
	fmt.Println(strings.Trim("~213~213~2311312~aaa~bbb~", "~"))
	fmt.Println(strings.TrimPrefix("~adsa~", "~"))
	fmt.Println(strings.TrimSuffix("~adsa~", "~"))
	fmt.Println(strings.HasPrefix("aaa.jpg", ".jpg"))
}
