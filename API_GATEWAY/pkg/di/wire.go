//go:build wireinject
// +build wireinject

package di

import (
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/api"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/api/handler"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/client"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/config"
	"github.com/google/wire"
)

func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(client.InitVideoClient, client.NewVideoServiceClient, handler.NewVideoHandler, api.NewServeHTTP)
	return &api.Server{}, nil
}
