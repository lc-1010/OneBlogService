package middleware

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
)

func TranslateSet(uni *ut.UniversalTranslator, v *validator.Validate) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 将翻译器和验证器存储在请求上下文中
		c.Set("uni", uni)
		c.Set("validator", v)
		c.Next()
	}

}

// func Translations() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
// 		locale := c.GetHeader("locale")

// 		trans, _ := uni.GetTranslator(locale)
// 		v, ok := binding.Validator.Engine().(*validator.Validate)

// 		if ok {
// 			switch locale {
// 			case "zh":
// 				_ = zh_translations.RegisterDefaultTranslations(v, trans)
// 			case "en":
// 				_ = en_translations.RegisterDefaultTranslations(v, trans)
// 			default:
// 				_ = zh_translations.RegisterDefaultTranslations(v, trans)
// 			}
// 			c.Set("trans", trans)
// 			c.Set("validator", v)
// 		}
// 		c.Next()
// 	}
// }
