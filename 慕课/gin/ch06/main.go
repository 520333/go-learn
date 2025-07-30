package main

import (
	"gin/ch06/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtobuf", returnProto)
	router.Run(":8083")

}
func moreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "dawn"
	msg.Message = "这是一个测试json"
	msg.Number = 20

	c.JSON(http.StatusOK, msg)
}
func returnProto(c *gin.Context) {
	course := []string{"golang", "java", "python"}
	user := &proto.Teacher{
		Name:   "dawn",
		Course: course,
	}

	c.ProtoBuf(http.StatusOK, user)

}
