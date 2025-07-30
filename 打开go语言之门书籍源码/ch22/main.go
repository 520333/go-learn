package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("初始化一些功能...")
}
func main() {
	fmt.Println("先导入fmt包才能使用")
	router := gin.Default()
	router.Run()

}
