package service

import "github.com/gin-gonic/gin"

type HttpServer struct {}

func (h *HttpServer) InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}