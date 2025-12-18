package common

import (
	"ginCms/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var translator ut.Translator

func Translate(err error) gin.H {
	//仅翻译验证消息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}
	// 翻译
	msg := gin.H{}
	for _, err := range errs {
		msg[err.Field()] = err.Translate(translator)
	}
	return msg
}

func translateMessage() {
	universalTranslator := ut.New(zh.New()) //通用翻译器

	// 具体验证引擎
	validate := binding.Validator.Engine().(*validator.Validate)

	// 具体翻译器
	translator, _ = universalTranslator.GetTranslator("zh")

	// 注册为默认翻译器
	if err := zhTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
		utils.Logger().Warn(err.Error())
	}

	// 注册TagName自定义函数 从json中获取tag
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})
}

func init() {
	// 翻译消息
	translateMessage()
}
