package model

import (
	"time"
)

type TUser struct {
	//
	Guid int64 `json:"guid" form:"guid" gorm:"primaryKey" `
	//
	Name string `json:"name" form:"name" `
	//
	Age int64 `json:"age" form:"age" `
	//
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime" `
}

// TableName 解决gorm表明映射
func (TUser) TableName() string {
	return "t_user"
}
