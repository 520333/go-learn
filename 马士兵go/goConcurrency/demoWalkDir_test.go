package goConcurrency

import (
	"fmt"
	"testing"
)

func TestWalkDir(t *testing.T) {
	dirs := []string{
		`D:\K8S`,
		`D:\Virtual Machines`,
	}
	fmt.Println(WalkDir(dirs...))
}
func TestQuickSortConcurrency(t *testing.T) {
	randArr := []int{
		19, 21, 0, -8, 11, 12, 19, 7, 25, 33, 2, 5, 100, 221, 21, 51,
	}
	sortArr := QuickSortConcurrency(randArr)
	fmt.Println(sortArr)
}

func TestGenerateRandArr(t *testing.T) {
	randArr := GenerateRandArr(10000)
	fmt.Println(randArr)
}
