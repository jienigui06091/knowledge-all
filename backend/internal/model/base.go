package model

import (
	"time"

	"gorm.io/gorm"
)

// Base 包含所有具有创建、更新时间的表的公共字段。
type Base struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SoftDelete 为数据库中使用 deleted_at 软删除的表提供公共字段。
type SoftDelete struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func Offset(page, pageSize int) int {
	return (page - 1) * pageSize
}
