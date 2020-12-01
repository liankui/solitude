package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/logic"
)

type HttpServer struct {}

func (h *HttpServer) InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/shorten", logic.Shorten)
	r.GET("/expand", logic.Expand)

	return r
}