package api

import (
	"fmt"
	"net"

	"github.com/MohdAhzan/StreamVideoGRPC/VIDEO_STREAM_SVC/pkg/config"
	"github.com/MohdAhzan/StreamVideoGRPC/VIDEO_STREAM_SVC/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewgrpcServe(c *config.Config, service pb.VideoServiceServer) (*Server, error) {
	grpcserver := grpc.NewServer()
	pb.RegisterVideoServiceServer(grpcserver, service)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		return nil, err
	}
	return &Server{
		gs:   grpcserver,
		Lis:  lis,
		Port: c.Port,
	}, nil
}

func (s *Server) Start() error {
	fmt.Println("Video service on:", s.Port)
	return s.gs.Serve(s.Lis)
}
