package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/config"
	"github.com/liankui/solitude/dao"
	"github.com/liankui/solitude/handler"
	"github.com/spf13/viper"
)

// @title Solitude
// @version 1.0
// @description 一个golang开发的短链接项目
// @termsOfService https://github.com/liankui/solitude

// @contact.name Eric
// @contact.url https://github.com/liankui
// @contact.email

// @host localhost:5555
// @BasePath
func main() {
	configPath := flag.String("c", "config.yaml", "configuration file")
	flag.Parse()
	gin.SetMode(viper.GetString("GIN_MODE"))


	config.InitCfg(*configPath)
	dao.DB = dao.InitTestDB()
	dao.Redis = dao.NewRedis()

	var h handler.HttpServer
	h.InitRouter()
}
