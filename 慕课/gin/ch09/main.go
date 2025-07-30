package main

// gin html模板
import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	router.LoadHTMLGlob("templates/**/*")
	// router.LoadHTMLFiles("templates/index.tmpl", "templates/goods.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "宝哥网",
		})
	})

	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"name": "微服务开发",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/goods/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/list.html", gin.H{"title": "商品"})
	})
	router.GET("/users/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/list.html", gin.H{"title": "用户"})
	})

	router.Run(":8083")

}
