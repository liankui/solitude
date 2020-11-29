package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/liankui/solitude/handler"
	"github.com/spf13/viper"
	"log"
)

func main() {
	configPath := flag.String("c", "../config.yaml", "configuration file")
	flag.Parse()
	log.Println("configPath:", *configPath)

	gin.SetMode(viper.GetString("GIN_MODE"))

	var h handler.HttpServer
	r := h.InitRouter()
	if err := r.Run(); err != nil {
		log.Fatalf("gin run error: %v", err)
	}
}
