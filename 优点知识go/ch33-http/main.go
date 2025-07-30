package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
)

func sendHttpRequest(url string) {
	// 设置header头
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User_Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1")
	req.Header.Add("Host", "www.baidu.com")

	resp, err := http.DefaultClient.Do(req)

	// http客户端
	// resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(resp)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"int"`
}

func main() {
	// sendHttpRequest("https://cn.bing.com")
	// http服务端 监听端口
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "golang pong\n")
	})

	// json数据
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Name: "宝哥",
			Age:  19,
		}
		userJSON, err := json.Marshal(user) //结构体转json
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(userJSON)
	})

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		image := path.Join("ch33-http", "900.gif")
		http.ServeFile(w, r, image)
	})

	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Name: "宝哥",
			Age:  19,
		}
		httpFile := path.Join("ch33-http", "index.html")
		tmpl, err := template.ParseFiles(httpFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := tmpl.Execute(w, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	log.Fatal(http.ListenAndServe(":12345", nil))
}
