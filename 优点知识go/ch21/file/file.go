package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(f string) (string, error) {
	file, err := os.Open(f)
	defer fmt.Println("first defer function")
	defer file.Close()
	defer fmt.Println("second defer function")

	if err != nil {
		// return "", err
		panic(err.Error())
	}
	bts, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err

	}
	return string(bts), nil

}
