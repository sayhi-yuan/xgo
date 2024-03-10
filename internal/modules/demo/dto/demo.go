package dto

import "xgo/core"

type DemoRequest struct {
	core.Paging
}

type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}
