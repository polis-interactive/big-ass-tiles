package grpc

type Config interface {
	GetGrpcPort() int
}
