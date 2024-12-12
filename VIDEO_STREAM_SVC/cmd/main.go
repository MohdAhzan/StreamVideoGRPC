package main

import (
	"log"

	"github.com/MohdAhzan/StreamVideoGRPC/VIDEO_STREAM_SVC/pkg/config"
	"github.com/MohdAhzan/StreamVideoGRPC/VIDEO_STREAM_SVC/pkg/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config", err.Error())
	}
	server, err1 := di.InitializeService(cfg)
	if err1 != nil {
		log.Fatal("failed to init server", err1.Error())
	}
	if err := server.Start(); err != nil {
		log.Fatal("couldn't start the server")
	}
}
