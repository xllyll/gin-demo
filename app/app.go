package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

var R = gin.Default()

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
