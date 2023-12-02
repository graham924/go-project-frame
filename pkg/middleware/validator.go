package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"go-project-frame/pkg/consts"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
)

// Validator 自定义参数验证器（包含验证失败时 的 错误信息国际翻译器）
func Validator() gin.HandlerFunc {
	// 参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
	return func(ctx *gin.Context) {
		// 1、创建一个参数验证器 validate
		validate := validator.New()

		// 2、创建一个国际化翻译器 uni
		// 设置支持的语言
		enLan := en.New()
		zhLan := zh.New()
		uni := ut.New(zhLan, zhLan, enLan)

		// 3、从请求参数中获取当前参数语言，如果没有则返回默认值zh
		locale := ctx.DefaultQuery("locale", "zh")

		// 4、从国际化翻译器中，获取 对应语言的翻译器
		trans, _ := uni.GetTranslator(locale)

		// 5、将翻译器注册到 验证器validator
		switch locale {
		case "en":
			if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
				return
			}
			// 将一个匿名函数注册到验证器实例 val 中。在这个匿名函数中，通过 fld.Tag.Get("en_comment") 的方式获取字段标签中 "en_comment" 的值，并返回给验证器。
			// 这样，在进行数据验证时，验证器会尝试获取字段的标签值，并将其作为错误信息中字段的说明或注释，提高错误信息的可读性和用户体验。
			validate.RegisterTagNameFunc(func(field reflect.StructField) string {
				return field.Tag.Get("en_comment")
			})
			break
		default:
			if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
				return
			}
			// 将一个匿名函数注册到验证器实例 val 中。在这个匿名函数中，通过 fld.Tag.Get("zh_comment") 的方式获取字段标签中 "zh_comment" 的值，并返回给验证器。
			// 这样，在进行数据验证时，验证器会尝试获取字段的标签值，并将其作为错误信息中字段的说明或注释，提高错误信息的可读性和用户体验。
			validate.RegisterTagNameFunc(func(field reflect.StructField) string {
				return field.Tag.Get("zh_comment")
			})
			break
		}
		// 6、将 验证器+翻译器 存到上下文中，方便后续处理的时候使用
		ctx.Set(consts.ValidatorContextKey, validate)
		ctx.Set(consts.TranslatorContextKey, trans)

		// 7、放行
		ctx.Next()
	}
}
