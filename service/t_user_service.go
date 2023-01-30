package service

import (
	"gin-demo/common"
	"gin-demo/model"
	"gorm.io/gorm"
	"log"
)

var TUserService = tUserService{}

// tUserService 业务层
type tUserService struct {
}

func (t tUserService) db() *gorm.DB {
	return common.Db
}

// List 分页列表
func (t tUserService) List(page, size int, v *model.TUser) map[string]interface{} {
	// 结果
	var lists []model.TUser
	t.db().Model(&v).Where(&v).Order("").Offset((page - 1) * size).Limit(size).Find(&lists)
	// 统计
	var total int64
	t.db().Model(&v).Where(&v).Count(&total)
	data := make(map[string]interface{})
	data["list"] = lists
	data["total"] = total
	return data
}

// One 根据主键Id查询记录
func (t tUserService) One(id interface{}) model.TUser {
	var v model.TUser
	db := t.db().Find(&v, id)
	if db.RowsAffected != 1 {
		log.Println("未找到数据！")
	}
	return v
}

// Update 修改记录 true -> 操作成功
func (t tUserService) Update(v model.TUser) bool {
	tx := t.db().Model(&v).Updates(v)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}

// Insert 插入记录 true -> 操作成功 注: 主键也传递进来的话，那么就会执行更新或插入操作
func (t tUserService) Insert(v model.TUser) bool {
	tx := t.db().Save(&v)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}

// Delete 根据主键删除 true -> 操作成功
func (t tUserService) Delete(ids []int) bool {
	tx := t.db().Delete(model.TUser{}, ids)
	if tx.Error != nil {
		log.Panicln(tx.Error.Error())
		return false
	}
	return true
}
