package core

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DataList struct {
	List     interface{} `json:"list"`
	PageInfo PageInfo    `json:"page_info"`
}

func (p *DataList) Make(data any, paging Paging) *DataList {
	v := reflect.ValueOf(data)

	// 检测一定是数组或切片
	if !(v.Kind() == reflect.Array || v.Kind() == reflect.Slice) {
		panic("分页数据必须为数组或者切片")
	}

	// 获取数据总量
	total := v.Len()
	dataList := []any{}
	start := (paging.Page - 1) * paging.PageSize
	if start <= total-1 {
		for ; start < total; start++ {
			dataList = append(dataList, v.Index(start).Interface())
			if len(dataList) == paging.PageSize {
				break
			}
		}
	}

	p.List = dataList
	p.PageInfo = (&PageInfo{}).Out(paging, int64(total))
	return p
}

type DataListTotal struct {
	DataList
	TotalData interface{} `json:"total_data"`
}

// PageInfo 分页返回信息
type PageInfo struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

func (p *PageInfo) Out(paging Paging, total int64) PageInfo {
	p.Page = paging.Page
	p.PageSize = paging.PageSize
	p.Total = int(total)
	p.TotalPage = int(math.Ceil(float64(total) / float64(paging.PageSize)))

	return *p
}

// 常量

const RequestIDKey = "x-request-id"

type responseBase struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Timestamp int64  `json:"timestamp"`
}

type responseData struct {
	responseBase

	Data interface{} `json:"data"`
}

type responseHandle struct {
}

// Success 成功返回
func (handle responseHandle) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &responseBase{
		Code:      SuccessCode.Index(),
		Message:   SuccessCode.String(),
		RequestID: handle.getRequestId(ctx),
		Timestamp: time.Now().Unix(),
	})
}

// SuccessMessage 自定义设置message
func (handle responseHandle) SuccessMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, &responseBase{
		Code:      SuccessCode.Index(),
		Message:   message,
		RequestID: handle.getRequestId(ctx),
		Timestamp: time.Now().Unix(),
	})
}

// SuccessData 返回数据
func (handle responseHandle) SuccessData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &responseData{
		responseBase: responseBase{
			Code:      SuccessCode.Index(),
			Message:   SuccessCode.String(),
			RequestID: handle.getRequestId(ctx),
			Timestamp: time.Now().Unix(),
		},
		Data: data,
	})
}

// SuccessDataList 返回数据
func (handle responseHandle) SuccessDataList(ctx *gin.Context, data DataList) {
	ctx.JSON(http.StatusOK, &responseData{
		responseBase: responseBase{
			Code:      SuccessCode.Index(),
			Message:   SuccessCode.String(),
			RequestID: handle.getRequestId(ctx),
			Timestamp: time.Now().Unix(),
		},
		Data: data,
	})
}

// Fail 失败返回
func (handle responseHandle) Fail(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &responseBase{
		Code:      FailCode.Index(),
		Message:   FailCode.String(),
		RequestID: handle.getRequestId(ctx),
		Timestamp: time.Now().Unix(),
	})
}

// FailCode 根据code返回失败
func (handle responseHandle) FailCode(ctx *gin.Context, code responseCode) {
	ctx.JSON(http.StatusOK, &responseBase{
		Code:      code.Index(),
		Message:   code.String(),
		RequestID: handle.getRequestId(ctx),
		Timestamp: time.Now().Unix(),
	})
}

// FailMessage 自定义设置message
func (handle responseHandle) FailMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, &responseBase{
		Code:      FailCode.Index(),
		Message:   message,
		RequestID: handle.getRequestId(ctx),
		Timestamp: time.Now().Unix(),
	})
}

// FailError 根据error返回失败
func (handle responseHandle) FailError(ctx *gin.Context, err error) {
	if err == gorm.ErrRecordNotFound {
		err = fmt.Errorf("未查询到该记录信息")
	}

	ctx.JSON(http.StatusOK, &responseBase{
		Code:      FailCode.Index(),
		Message:   err.Error(),
		RequestID: handle.getRequestId(ctx),
		Timestamp: time.Now().Unix(),
	})
}

// 获取request-id
func (handle responseHandle) getRequestId(ctx *gin.Context) string {
	id, _ := ctx.Get(RequestIDKey)
	return id.(string)
}

// 返回码
type responseCode int

const (
	SuccessCode responseCode = 1  // 成功
	FailCode    responseCode = -1 // 失败
)

func (code responseCode) String() string {
	return map[responseCode]string{
		SuccessCode: "success",
		FailCode:    "fail",
	}[code]
}

func (code responseCode) Index() int {
	return int(code)
}
