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
