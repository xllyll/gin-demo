package model

import (
	"time"
)

type TSmsCode struct {
	//
	Guid int64 `json:"guid" form:"guid" gorm:"primaryKey" `
	//
	UserGuid int64 `json:"userGuid" form:"userGuid" `
	//
	Code string `json:"code" form:"code" `
	//
	Phone string `json:"phone" form:"phone" `
	//
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"autoCreateTime" `
}

// TableName 解决gorm表明映射
func (TSmsCode) TableName() string {
	return "t_sms_code"
}
