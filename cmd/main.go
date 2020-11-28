package main

import (
	"flag"
	"github.com/liankui/solitude/cmd/service"
	"log"
)

func main() {
	configPath := flag.String("c", "../config.yaml", "configuration file")
	flag.Parse()
	log.Println("configPath:", *configPath)

	var h service.HttpServer
	r := h.InitRouter()
	r.Run()
}
