package grpc

import (
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcControl "github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/api/v1/go"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	port       int
	c          controller
}

func NewGrpcController(
	c controller,
	config Config,
) (*Server, error) {

	log.Println("GrpcServer, NewServer: creating")

	var options []grpc.ServerOption
	options = append(
		options,
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcCtxTags.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcCtxTags.UnaryServerInterceptor(),
		)),
	)

	grpcSrv := grpc.NewServer(options...)

	controlService, err := newParticipantServer(c)
	if err != nil {
		return nil, err
	}
	grpcControl.RegisterControlServer(grpcSrv, controlService)

	log.Println("GrpcServer, NewServer: created successfully")

	return &Server{
		grpcServer: grpcSrv,
		port:       config.GetGrpcPort(),
		c:          c,
	}, nil
}

func (s *Server) RunMainLoop() {

	log.Println("GrpcServer, Startup: starting")

	addr := fmt.Sprintf("0.0.0.0:%d", s.port)
	log.Println(fmt.Sprintf("GrpcServer, Startup: listening at %s", addr))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("GrpcServer, Startup: Failed to listen: %v", err)
		return
	}

	go func() {
		err := s.grpcServer.Serve(listener)
		if err != nil {
			log.Printf("GrpcServer: reported err %s", err)
		}
	}()

	log.Println("GrpcServer, Startup: started")

	shutdowns := s.c.GetShutdowns()

	for {
		select {
		case _, ok := <-shutdowns:
			if !ok {
				goto Cleanup
			}
		}
	}

Cleanup:

	log.Println("GrpcServer, Startup: stopped")
	wg := s.c.GetWg()
	wg.Done()

}
