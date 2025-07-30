package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"yeam"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{Title: "哥斯拉大战金刚", Year: 2022, Price: 10, Actors: []string{"金刚", "哥吉拉", "人类群演"}}
	//结构体转json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json Marshal error", err)
		return
	}
	fmt.Printf("jsonstr = %s\n", jsonStr)

	//json转结构体
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json--->struct error: ", err)
		return
	}
	fmt.Printf("json ---> struct: %v\n", my_movie)
}
