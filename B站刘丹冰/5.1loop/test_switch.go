package main

import "fmt"

func main() {
	// switch var1 { case var2: ... case var3: ... default: ...}
	// 中文星期几 输出对应英文
	day := "星期三"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	case "星期三":
		fmt.Println("Wednesday")
	case "星期四":
		fmt.Println("Thursday")
	case "星期五":
		fmt.Println("Friday")
	case "星期六":
		fmt.Println("Saturday")
	case "星期日":
		fmt.Println("Sunday")
	default:
		fmt.Println("这天可能是个星期八")
	}

	sore := 95
	switch {
	case sore < 60:
		fmt.Println("E")
	case sore >= 60 && sore < 70:
		fmt.Println("D")
	case sore >= 70 && sore < 80:
		fmt.Println("C")
	case sore >= 80 && sore < 90:
		fmt.Println("B")
	case sore >= 90 && sore <= 100:
		fmt.Println("A")
	}
	switch sore {
	case 60, 70, 95:
		fmt.Println("wow...")
	default:
		fmt.Println("on~ no")
	}
	var sum int
	for x := 1; x <= 100; x++ {
		sum += x
	}
	fmt.Println(sum)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
