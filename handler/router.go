package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/liankui/solitude/docs"
	"github.com/liankui/solitude/logic"
	"github.com/spf13/viper"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

type HttpServer struct{}

func (h *HttpServer) InitRouter() {
	r := gin.Default()
	// 使用http://localhost:5555/swagger/index.html进去
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.GET("/shorten", logic.Shorten)
	r.GET("/expand/:shorten", logic.Expand)
	r.GET("/print", logic.Print)	// 打印最近的一次记录

	err := r.Run(":"+viper.GetString("Addr"))
	if err != nil {
		log.Fatalf("router running error: %v", err)
	}
}