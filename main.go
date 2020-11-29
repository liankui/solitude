package main

import (
	"flag"
	"github.com/liankui/solitude/handler"
	"log"
)

func main() {
	configPath := flag.String("c", "../config.yaml", "configuration file")
	flag.Parse()
	log.Println("configPath:", *configPath)

	var h handler.HttpServer
	r := h.InitRouter()
	r.Run()
}
