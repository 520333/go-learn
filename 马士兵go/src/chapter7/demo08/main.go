package main

import "fmt"

func main() {
	var arr [3][3]int = [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	for key, value := range arr {
		for k, v := range value {
			fmt.Printf("arr[%v][%v]=%v \n", key, k, v)
		}
	}
}
