package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", welcome)
	router.POST("/form_post", formPost)
	router.POST("/post", getPost)
	router.Run(":8083")
}

// http://localhost:8083/welcome?firstname=abao&lastname=asd
func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "dawn") //设置默认值
	lastName := c.DefaultQuery("lastname", "chuang") //设置默认值
	c.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})

}
func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}
func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.DefaultPostForm("message", "信息")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}
