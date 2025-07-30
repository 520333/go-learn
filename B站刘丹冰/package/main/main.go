package main

import (
	"fmt"
	"net/http"
	u "package/user"

	// . "learn/package/user"  //.将包导入到当前目录
	_ "package/user2" //_匿名导入 不使用

	"github.com/gin-gonic/gin"
)

func main() {
	c := u.Course{
		Name: "go",
	}
	fmt.Println(u.GetCourse(c))

	// c2 := u2.Course{
	// 	Name: "python",
	// }
	// fmt.Printf(u.GetCourse(c2))
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
