package myfunc

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	context.HTML(http.StatusOK, "demo01/hello.html", nil)
}

func Hello2(context *gin.Context) {
	// 获取前端传入的文件
	file, _ := context.FormFile("myfile")
	fmt.Println(file.Filename)
	time_int := time.Now().Unix()
	time_str := strconv.FormatInt(time_int, 10)
	err := context.SaveUploadedFile(file, "./"+time_str+file.Filename)
	if err != nil {
		log.Println("save file err:", err)
	}

	// 响应一个字符串
	context.String(200, "文件上传成功")

}

func Hello3(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["myfile"]
	for _, file := range files {
		time_int := time.Now().Unix()
		time_str := strconv.FormatInt(time_int, 10)
		err := context.SaveUploadedFile(file, "./"+time_str+file.Filename)
		if err != nil {
			log.Println("save file err:", err)
		}
	}

	// 响应一个字符串
	context.String(200, "文件上传成功")
}

func Hello4(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		time_int := time.Now().Unix()
		time_str := strconv.FormatInt(time_int, 10)
		err := context.SaveUploadedFile(file, "./"+time_str+file.Filename)
		if err != nil {
			log.Println("save file err:", err)
		}
	}

	// 响应一个字符串
	context.String(200, "文件上传成功")
}
