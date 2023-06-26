package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// BindAndValid 将错误返回到封装好的error中
func BindAndValid(c *gin.Context, v any) (bool, ValidErrors) {
	var errs ValidErrors
	// 绑定校验
	err := c.ShouldBind(v)
	if err != nil {
		locale := c.GetHeader("locale")
		v := c.Value("validator").(*val.Validate)
		uni := c.Value("uni").(*ut.UniversalTranslator)
		if locale == "" {
			locale = "zh"
		}
		trans, _ := uni.GetTranslator(locale)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}
		switch locale {
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, trans)
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, trans)
		default:
			_ = zh_translations.RegisterDefaultTranslations(v, trans)
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
