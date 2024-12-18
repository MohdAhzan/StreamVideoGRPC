package main

import (
	"log"

	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/config"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/di"
)

func main() {

	c, configerr := config.LoadConfig()
    
	if configerr != nil {
		log.Fatal("cannot load config:", configerr)
	}

	server, dierr := di.InitializeAPI(c)
	if dierr != nil {
		log.Fatal("cannot initialize server", dierr)
	}
	server.Start()
}
