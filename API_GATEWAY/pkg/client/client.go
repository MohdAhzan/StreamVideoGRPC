package client

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	interfaces "github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/client/interface"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/config"
	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type videoClient struct {
	Client pb.VideoServiceClient
}

func InitVideoClient(c *config.Config) (pb.VideoServiceClient, error) {
    fmt.Println("GRPC CLIENT SERVICE URL",c.VideoService)
  grpcConnection, err := grpc.NewClient(c.VideoService,grpc.WithTransportCredentials(insecure.NewCredentials()))
 
  if err != nil {
    fmt.Println("Could not connect", err)
  }

  return pb.NewVideoServiceClient(grpcConnection),nil

}

func NewVideoServiceClient(client pb.VideoServiceClient) interfaces.VideoClient {
	return &videoClient{
    Client:client ,
	}
}


func (c *videoClient) UploadVideo(ctx context.Context, file *multipart.FileHeader) (*pb.UploadVideoResponse, error) {
	upLoadfile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer upLoadfile.Close()
	stream, _ := c.Client.UploadVideo(ctx)
	chunkSize := 4096 
	buffer := make([]byte, chunkSize)
	for {
		n, err := upLoadfile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := stream.Send(&pb.UploadVideoRequest{
			Filename: file.Filename,
			Data:     buffer[:n],
		}); err != nil {
			return nil, err
		}
	}
	//the final response is recieved and the streaming is closed
	response, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *videoClient) StreamVideo(ctx context.Context, filename, playlist string) (pb.VideoService_StreamVideoClient, error) {
	res, err := c.Client.StreamVideo(ctx, &pb.StreamVideoRequest{
		Videoid:  filename,
		Playlist: playlist,
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *videoClient) FindAllVideo(ctx context.Context) (*pb.FindAllResponse, error) {
	res, err := c.Client.FindAllVideo(ctx, &pb.FindAllRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
