package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/logic"
	"log"
)

type HttpServer struct{}

func (h *HttpServer) InitRouter() {
	r := gin.Default()

	r.GET("/shorten", logic.Shorten)
	r.GET("/expand", logic.Expand)

	err := r.Run(":5555")
	if err != nil {
		log.Fatalf("router running error: %v", err)
	}
}