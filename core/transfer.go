package core

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func init() {
	if err := initTrans(); err != nil {
		panic(err)
	}
}

var Uni *ut.UniversalTranslator

func initTrans() error {
	// 修改gin框架中的 validator 引擎属性, 实现自定制
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法，返回错误字段使用 json tag 字段，而不是结构体字段名
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		enT := en.New()             // 英文翻译器
		zhT := zh.New()             // 中文翻译器
		Uni = ut.New(enT, zhT, enT) // 第一个参数是备用语言环境，后面的参数是应该支持的语言环境

		// 注册中文翻译器
		if trans, ok := Uni.GetTranslator("zh"); ok {
			if err := zhTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
				return err
			}
		} else {
			return fmt.Errorf(`Uni.GetTranslator("%s") error`, "zh")
		}

		// 注册英文翻译器
		if trans, ok := Uni.GetTranslator("en"); ok {
			if err := enTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
				return err
			}
		} else {
			return fmt.Errorf(`Uni.GetTranslator("%s") error`, "en")
		}
		return nil
	} else {
		return fmt.Errorf("")
	}
}

func getLocalTrans(locale string) (ut.Translator, error) {
	// 根据不同语言获取相应的翻译器
	if trans, ok := Uni.GetTranslator(locale); !ok {
		return nil, fmt.Errorf(`Uni.GetTranslator("%s") error`, locale)
	} else {
		return trans, nil
	}
}

func removeTopStruct(fields map[string]string) error {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}

	errStr := ""
	for _, item := range res {
		errStr += fmt.Sprintf("%s;", item)
	}

	return fmt.Errorf(errStr)
}
