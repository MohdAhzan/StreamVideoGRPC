package interfaces

import "github.com/MohdAhzan/StreamVideoGRPC/VIDEO_STREAM_SVC/pkg/pb"


type VideoRepo interface {
	CreateVideoid(string) error
	FindAllVideo() ([]*pb.VideoID, error)
}
