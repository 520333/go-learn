package main

import (
	"bufio"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func handleWebSocket(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			log.Println("升级协议", r.Header["User-Agent"])
			return true // 允许跨域
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()
	start := 0
	for {
		// 调用 Jenkins API 获取日志
		resp, err := HttpGetJenkinsBuildLog(start)
		//defer resp.Body.Close()
		if err != nil {
			log.Println("Jenkins API error:", err)
			return
		}
		// 逐行读取日志内容
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if err := conn.WriteMessage(websocket.TextMessage, []byte(line+"\n")); err != nil {
				log.Println("WebSocket write error:", err)
				resp.Body.Close()
				return
			}
		}
		// 更新 start 位置
		start, _ = strconv.Atoi(resp.Header.Get("X-Text-Size"))
		// newSize, _ := strconv.Atoi(resp.Header.Get("X-Text-Size"))

		// 更新读取位置
		// if newSize > start {
		// 	start = newSize
		// }
		// 如果没有更多数据，退出循环
		if resp.Header.Get("X-More-Data") != "true" {
			resp.Body.Close()
			break
		}
		resp.Body.Close()
		// 等待一段时间再继续请求
		time.Sleep(1 * time.Second)
	}
}
func HttpGetJenkinsBuildLog(start int) (res *http.Response, err error) {
	jenkinsURL := "http://10.166.4.109:28080/job/test-fronend/23/logText/progressiveText"
	request, err := http.NewRequest("GET", jenkinsURL, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	//echo -n "admin:your jenkins token" | base64
	request.Header.Set("Authorization", "Basic emstemh1YW5ndGI6MTFmMTdjZjdmNGE2ZDliODUwMzg5NTdlMjFlMmRlZmRiMQ==")
	//发送请求给服务端,实例化一个客户端
	client := &http.Client{}
	res, err = client.Do(request)
	return
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.GET("/ws", handleWebSocket)
	router.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}
