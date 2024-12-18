// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/api"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/api/handler"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/client"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/config"
)

// Injectors from wire.go:

func InitializeAPI(c *config.Config) (*api.Server, error) {
	videoServiceClient, err := client.InitVideoClient(c)
	if err != nil {
		return nil, err
	}
	videoClient := client.NewVideoServiceClient(videoServiceClient)
	videoHandler := handler.NewVideoHandler(videoClient)
	server, err := api.NewServeHTTP(c, videoHandler)
	if err != nil {
		return nil, err
	}
	return server, nil
}
