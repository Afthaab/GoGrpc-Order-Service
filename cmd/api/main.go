package main

import (
	"log"

	"github.com/order/service/pkg/config"
	"github.com/order/service/pkg/di"
)

func main() {
	cfg, cfgerr := config.LoadConfig()
	if cfgerr != nil {
		log.Fatalln("Could not load the config file :", cfgerr)
	}
	server, err := di.InitApi(cfg)
	if err != nil {
		log.Fatalln("Error in initializing the API", err)
	}
	server.Start()

}
