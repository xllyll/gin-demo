package main

import (
	"gin-demo/api"
	"gin-demo/app"
	"gin-demo/common"
	"log"
)

func main() {

	log.Print("======== begin =============")
	Run()

}
func Run() {
	app.InitConfig()
	common.InitDB()
	api.InitUser()
	api.InitSmsCode()
	app.R.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
