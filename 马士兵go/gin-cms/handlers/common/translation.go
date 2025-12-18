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

// 自定义错误消息
var customMsg = map[string]string{
	"roleTitleUnique": "{0}对应的角色名称已经存在",
	"roleKeyUnique":   "{0}对应的角色键已经存在",
}

func Translate(err error) gin.H {
	//仅翻译验证消息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return gin.H{}
	}
	// 翻译
	msg := gin.H{}
	for _, err := range errs {
		msg[err.Field()] = err.Translate(translator)
	}
	return msg
}

var translator ut.Translator

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

	// 注册自定义消息
	translateFn := func(ut ut.Translator, fe validator.FieldError) string {
		msg, err := ut.T(fe.Tag(), fe.Field())
		if err != nil {
			utils.Logger().Warn(err.Error())
			return ""
		}
		return msg
	}
	for tag, text := range customMsg {
		if err := validate.RegisterTranslation(tag, translator, func(ut ut.Translator) error {
			if err := ut.Add(tag, text, false); err != nil {
				return err
			}
			return nil
		}, translateFn); err != nil {
			utils.Logger().Warn(err.Error())
		}
	}
}

func init() {
	// 翻译消息
	translateMessage()
}
