package core

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Paging 分页信息
type Paging struct {
	Page     int `json:"page" desc:"页码" default:"1"`
	PageSize int `json:"page_size" desc:"每页的数量" default:"10"`
}

// Context 构建上下文
type Context struct {
	*gin.Context

	UserInfo UserInfo `json:"userInfo"`
}

// UserInfo 用户信息
type UserInfo struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
}

type requestHandle struct {
}

// BindParam 添加基础类型的默认值绑定
// 目前只支持int, string
func (handle requestHandle) BindParam(ctx *gin.Context, param interface{}) error {
	if err := ctx.ShouldBindJSON(&param); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		trans, _ := getLocalTrans("zh")

		return removeTopStruct(errs.Translate(trans))
	}

	v := reflect.ValueOf(param).Elem()
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Struct:
			s := v.Field(i)
			for j := 0; j < s.NumField(); j++ {
				switch s.Field(j).Kind() {
				case reflect.String:
					if s.Field(j).String() == "" {
						tagValue, tagExist := s.Type().Field(j).Tag.Lookup("default")
						if tagExist {
							s.Field(j).SetString(tagValue)
						}
					}
				case reflect.Int:
					if s.Field(j).Int() == 0 {
						tagValue, tagExist := s.Type().Field(j).Tag.Lookup("default")
						if tagExist {
							num, _ := strconv.Atoi(tagValue)
							s.Field(j).SetInt(int64(num))
						}
					}
				}
			}
		case reflect.String:
			if v.Field(i).String() == "" {
				tagValue, tagExist := v.Type().Field(i).Tag.Lookup("default")
				if tagExist {
					v.Field(i).SetString(tagValue)
				}
			}
		case reflect.Int:
			if v.Field(i).Int() == 0 {
				tagValue, tagExist := v.Type().Field(i).Tag.Lookup("default")
				if tagExist {
					num, _ := strconv.Atoi(tagValue)
					v.Field(i).SetInt(int64(num))
				}
			}
		}
	}

	return nil
}

// Swap 获取用户信息
func (handle requestHandle) Swap(ctx *gin.Context) *Context {
	// 根据token的信息，从redis中获取用户的信息
	return &Context{
		ctx,
		UserInfo{
			Name: "rabbit",
			Uid:  11,
		},
	}
}
