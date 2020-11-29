package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/logic"
)

type HttpServer struct {}

func (h *HttpServer) InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/shorturl", logic.Shorturl)

	return r
}