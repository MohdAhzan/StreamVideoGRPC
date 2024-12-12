//go:build wireinject
// +build wireinject

package di

import (
	"github.com/MohdAhzan/StreamVideo/VIDEO_STREAM_SVC/pkg/api"
	"github.com/MohdAhzan/StreamVideo/VIDEO_STREAM_SVC/pkg/api/service"
	"github.com/MohdAhzan/StreamVideo/VIDEO_STREAM_SVC/pkg/config"
	"github.com/MohdAhzan/StreamVideo/VIDEO_STREAM_SVC/pkg/db"
	"github.com/MohdAhzan/StreamVideo/VIDEO_STREAM_SVC/pkg/repository"
	"github.com/google/wire"
)

func InitializeService(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb, repository.NewVideoRepo, service.NewVideoServer, api.NewgrpcServe)
	return &api.Server{}, nil
}
