package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/config"
	"github.com/liankui/solitude/dao"
	"github.com/liankui/solitude/handler"
	"github.com/spf13/viper"
)

func main() {
	configPath := flag.String("c", "config.yaml", "configuration file")
	flag.Parse()
	gin.SetMode(viper.GetString("GIN_MODE"))

	config.InitCfg(*configPath)
	dao.DB = dao.InitTestDB()
	//dao.Redis = dao.NewRedis()

	var h handler.HttpServer
	h.InitRouter()
}
