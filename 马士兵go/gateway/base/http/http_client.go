package main

import (
	"fmt"
	"io"
	"net/http"
)

func Client() {
	// 客户端
	var client = http.Client{}
	resp, err := client.Get("http://localhost:8080/hello")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	bds, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bds))
}
