package xhttp

import (
	CN_ZH "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"reflect"
)

var Validate *validator.Validate

// Validate/v10 全局验证器
var (
	trans ut.Translator
	found bool
)

type CustomValidator struct{}

func (v CustomValidator) Validate(r *http.Request, data any) error {
	if err := Validate.Struct(data); err != nil {
		return err
	}
	return nil
}

// 初始化Validate/v10国际化
func init() {
	zh := CN_ZH.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator(zh.Locale())

	Validate = validator.New()
	//通过label标签返回自定义错误内容
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})

	Validate.RegisterTranslation("required", trans, requiredOverrideRegister, requiredOverrideTranslation)
	zhTranslations.RegisterDefaultTranslations(Validate, trans)
	httpx.SetValidator(CustomValidator{})
}

func requiredOverrideRegister(ut ut.Translator) error { //这个函数的作用是注册翻译模板
	return ut.Add("required", "{0}是一个必须填写的字段", true) // {}是占位符 true代表了是否重写已有的模板
}

func requiredOverrideTranslation(ut ut.Translator, fe validator.FieldError) string { // 这个函数的作用是负责翻译内容
	t, _ := ut.T("required", fe.Field()) //参数可以有多个，取决于注册对应Tag的模板有多少个占位符
	return t
}
