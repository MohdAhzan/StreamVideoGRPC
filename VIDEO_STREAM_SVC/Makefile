run:
	go run cmd/main.go

wire: 
	cd pkg/di && wire
proto:
	 protoc --go_out=. --go-grpc_out=. ./pkg/pb/video.proto