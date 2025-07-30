package main

import (
	"errors"
	"fmt"
	"learn/ch21/file"
)

type Divide struct {
	dividee int // 被除数
	divider int // 除数
}

func (d *Divide) Error() string {
	strFormat := `
	Cannot proceed,the divider is zero.
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, d.dividee)
}

func computeDiv(d *Divide) (result int, err error) {
	if d.divider == 0 {
		err = d
	} else {
		result = d.dividee
	}
	return
}

func main() {
	err := errors.New("a new err object")
	fmt.Printf("%v", err)

	err = fmt.Errorf("a fmt error format object:%s", err.Error())
	fmt.Println(err)

	de := Divide{100, 10}
	if result, err := computeDiv(&de); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	if c, err := file.ReadFile("abc.txt"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(c)
	}

}
