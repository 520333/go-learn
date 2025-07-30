package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

type LoginForm struct {
	User     string `json:"user" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required"`
}

type SignUpForm struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// POST响应体只显示.后面的字段
func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for filed, err := range fileds {
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	return rsp
}

// I18N国际化
func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法 用于发起POST请求时响应体返回struct中定义的json tag
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()              //中文翻译器
		enT := en.New()              //英文翻译器
		uni := ut.New(enT, zhT, enT) // 第一个参数是备用语言环境 后面的参数是支持的语言
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func main() {
	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}
	router := gin.Default()
	router.POST("/loginJSON", func(c *gin.Context) {
		var loginForm LoginForm
		if err := c.ShouldBind(&loginForm); err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": removeTopStruct(errs.Translate(trans)),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
		})
	})

	router.POST("/sigup", func(ctx *gin.Context) {
		var signUpForm SignUpForm
		if err := ctx.ShouldBind(&signUpForm); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	router.Run(":8083")
}
