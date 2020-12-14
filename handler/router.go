package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/logic"
	"github.com/spf13/viper"
	"log"
)

type HttpServer struct{}

func (h *HttpServer) InitRouter() {
	r := gin.Default()

	r.GET("/shorten", logic.Shorten)
	r.GET("/expand/:shorten", logic.Expand)
	r.GET("/print", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world1",
		})
	})

	err := r.Run(":"+viper.GetString("Addr"))
	if err != nil {
		log.Fatalf("router running error: %v", err)
	}
}