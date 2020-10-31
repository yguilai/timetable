package core

import (
	"ccsu/global"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	v "github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func init() {
	_zh := zh.New()
	uni := ut.New(_zh, _zh)
	validateTrans, _ := uni.GetTranslator("zh")
	validator := v.New()
	// 收集结构体中的comment标签，用于替换英文字段名称，这样返回错误就能展示中文字段名称了
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})
	if err := zh_translations.RegisterDefaultTranslations(validator, validateTrans); err != nil {
		global.Log.Info(fmt.Sprintf("RegisterDefaultTranslations %v", err))
	}
	global.ValidateTrans = validateTrans
	global.Validator = validator
}
