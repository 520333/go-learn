package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	var addr = "localhost:8002"
	http.HandleFunc("/wsHandler", wsHandler)
	log.Println("start http server on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{} //默认选项
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	defer conn.Close()
	// 服务器主动向客户端推送消息
	go func() {
		for {
			// TextMessage:1 BinaryMessage:2
			if err := conn.WriteMessage(1, []byte("heart beat 心跳检测")); err != nil {
				return
			}
			time.Sleep(time.Second * 3)
		}
	}()

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		fmt.Printf("receive msg:%s\n", msg)
		newMsg := string(msg) + "haha"
		msg = []byte(newMsg)
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
