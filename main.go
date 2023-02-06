package main

import (
	"gin-demo/api"
	"gin-demo/app"
	"gin-demo/common"
	"github.com/spf13/viper"
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
	port := viper.GetString("server.port")
	var portStr = ":8080"
	if port != "" && len(port) > 0 {
		portStr = ":" + port
	}
	app.R.Run(portStr) // 监听并在 0.0.0.0:8080 上启动服务
}
