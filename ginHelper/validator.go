package ginHelper

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

// 格式化返回错误信息

var (
	trans ut.Translator
)

// InitBoostValidator 更强大的验证器
// 错误返回格式化
// 更多的验证格式支持
func InitBoostValidator(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 让错误对应的字段为json标签
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})
		// 注册翻译器
		err = initTranslator(locale, v)
		if err != nil {
			return
		}
		// 注册自定义格式的验证器
		err = registerCustomValidator(v, "mobile", validateMobile, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}必须是有效的手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
		return
	}
	err = errors.New("初始化失败")
	return
}

func initTranslator(locale string, v *validator.Validate) (err error) {
	// 注册翻译器
	zhT := zh.New() //中文翻译器
	enT := en.New() //英文翻译器
	//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
	uni := ut.New(enT, zhT, enT)
	var ok bool
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", locale)
	}

	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return
}

// removeTopStruct 把字段名的前置部分去除掉
// 比如: Resp.Mobile --> Mobile
func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func FormatValidatorError(context *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		Fail(context, http.StatusBadRequest, -1, "参数错误")
		return
	}
	Fail(context, http.StatusBadRequest, -1, removeTopStruct(errs.Translate(trans)))
	return
}

// validateMobile 手机格式验证
func validateMobile(fld validator.FieldLevel) bool {
	mobile := fld.Field().String()
	// 正则验证
	pattern := "^1[3-9]\\d{9}$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobile)
}

func registerCustomValidator(v *validator.Validate, tag string,
	validatorMethod func(fld validator.FieldLevel) bool,
	registerFormat func(ut ut.Translator) error,
	registerTranslation func(ut ut.Translator, fe validator.FieldError) string) error {
	if err := v.RegisterValidation(tag, validatorMethod); err != nil {
		return err
	}
	// 注册翻译器
	if err := v.RegisterTranslation(tag, trans, registerFormat, registerTranslation); err != nil {
		return err
	}
	return nil
}
