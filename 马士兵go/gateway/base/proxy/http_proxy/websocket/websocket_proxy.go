package main

import (
	"fmt"
	"log"
	"net/http"

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
