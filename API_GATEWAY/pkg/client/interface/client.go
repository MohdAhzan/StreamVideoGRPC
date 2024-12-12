package interfaces

import (
	"context"
	"mime/multipart"

	"github.com/MohdAhzan/StreamVideoGRPC/API_GATEWAY/pkg/pb"
)

type VideoClient interface {
	UploadVideo(context.Context, *multipart.FileHeader) (*pb.UploadVideoResponse, error)
	StreamVideo(context.Context, string, string) (pb.VideoService_StreamVideoClient, error)
	FindAllVideo(context.Context) (*pb.FindAllResponse, error)
}
