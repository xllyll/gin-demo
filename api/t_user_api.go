package api

import (
	"gin-demo/app"
	"gin-demo/common"
	"gin-demo/model"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func InitUser() {
	c := &TUserApi{}
	c.init(app.R) //这里需要引入你的gin框架的实例
}

func (t TUserApi) init(g *gin.Engine) {
	// 依次: 分页列表，单条，插入，修改，删除
	group := g.Group("/t/user")
	group.GET("/list", t.list) //不设置限制条件的画默认查询所有
	group.GET("/one", t.one)
	group.POST("/insert", t.insert)
	group.POST("/update", t.update)
	group.POST("/delete", t.delete)
}

// TUserApi 控制器
type TUserApi struct {
	Page int   `form:"page"`
	Size int   `form:"size"`
	Ids  []int `form:"ids"`
}

func (t TUserApi) db() *gorm.DB {
	return common.Db
}

// 分页列表
func (t TUserApi) list(c *gin.Context) {
	_ = c.Bind(&t)
	v := model.TUser{}
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.TUserService.List(t.Page, t.Size, &v)))
}

// 根据主键Id查询记录
func (t TUserApi) one(c *gin.Context) {
	var v model.TUser
	_ = c.Bind(&v)
	c.JSON(http.StatusOK, model.OkData(service.TUserService.One(v.Guid)))
}

// 修改记录
func (t TUserApi) update(c *gin.Context) {
	var v model.TUser
	_ = c.ShouldBindJSON(&v)
	service.TUserService.Update(v)
	c.JSON(http.StatusOK, model.OkMsg("修改成功！"))
}

// 插入记录
func (t TUserApi) insert(c *gin.Context) {
	var v model.TUser
	_ = c.ShouldBindJSON(&v)
	service.TUserService.Insert(v)
	c.JSON(http.StatusOK, model.OkMsg("插入成功！"))
}

// 根据主键删除
func (t TUserApi) delete(c *gin.Context) {
	_ = c.Bind(&t)
	service.TUserService.Delete(t.Ids)
	c.JSON(http.StatusOK, model.OkMsg("删除成功！"))
}
