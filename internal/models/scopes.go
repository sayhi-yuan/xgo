package models

import (
	"xgo/core"

	"gorm.io/gorm"
)

// 分页
func Paginate(paging core.Paging) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((paging.Page - 1) * paging.PageSize).Limit(paging.PageSize)
	}
}
