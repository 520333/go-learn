package main

import (
	"gin/part11/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("part11/templates/**/*")
	var v1 = r.Group("/version1") //访问：http://localhost:8080/version1/userindex
	{
		v1.GET("/")
		v1.GET("/userindex", myfunc.Hello1)
		v1.GET("/toFormBind", myfunc.Hello2)
		v1.GET("/userindex2", myfunc.Hello3)
	}

	var v2 = r.Group("/version2") //访问：http://localhost:8080/version2/userindex4/丽丽/19
	{
		v2.GET("/userindex3", myfunc.Hello4)
		v2.POST("/toajax", myfunc.Hello5)
		v2.GET("/userindex4/:uname/:age", myfunc.Hello6)
	}
	r.Run(":8080")
}
