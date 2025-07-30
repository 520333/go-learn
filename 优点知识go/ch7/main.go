package main

import "fmt"

func operatorDemo() {
	var a, b = 21, 10
	fmt.Printf("%d+%d=%d\n", a, b, a+b)
	fmt.Printf("%d-%d=%d\n", a, b, a-b)
	fmt.Printf("%d*%d=%d\n", a, b, a*b)

	fmt.Printf("%d/%d=%f\n", a, b, float32(a/b))
	fmt.Printf("%d%%%d=%d\n", a, b, a%b)
	a++
	fmt.Println(a)
	a = 21
	a--
	fmt.Println(a)
}

func relationDemo() {
	var a int = 21
	var b int = 10
	if a == b {
		fmt.Println("a==b")
	} else {
		fmt.Println("a!=b")
	}

	if a > b {
		fmt.Println("a>b")
	} else {
		fmt.Println("a<=b")
	}
	if a < b {
		fmt.Println("a<b")
	} else {
		fmt.Println("a<=b")
	}
}

func logicDemo() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("a和b都是真")
	} else {
		fmt.Println("a和b有一个为假")
	}

	if a || b {
		fmt.Println("a和b有一个为真")
	} else {
		fmt.Println("a和b都为假")
	}
	if !a {
		fmt.Println("a为假")
	} else {
		fmt.Println("a为真")
	}
}

func byteComputeDemo() {
	var a uint = 60
	fmt.Printf("%b\n", a)
	var b uint = 13
	fmt.Printf("%b\n", b)

	// 与运算
	//a=111100
	//b=001101
	//c=001100
	c := a & b
	fmt.Printf("与预算:%b %d\n", c, c)

	// 或运输
	//a=111100
	//b=001101
	//c=111101
	c = a | b
	fmt.Printf("%b %d\n", c, c)

	// 亦或运输
	//a=111100
	//b=001101
	//c=110001
	c = a ^ b
	fmt.Printf("%b %d\n", c, c)

	c = a << 2 // 乘法
	fmt.Printf("%b %d\n", c, c)

	c = a >> 2 // 除法
	fmt.Printf("%b %d\n", c, c)

}
func main() {
	operatorDemo()
	relationDemo()
	logicDemo()
	byteComputeDemo()
}
