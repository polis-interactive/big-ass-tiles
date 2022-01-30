package grpc

import (
	"context"
	grpcControl "github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/api/v1/go"
	"log"
)

type controlServer struct {
	grpcControl.UnsafeControlServer
	c controller
}

var _ grpcControl.ControlServer = (*controlServer)(nil)

func newParticipantServer(c controller) (*controlServer, error) {

	log.Println("participantServer, newServer: creating")

	return &controlServer{
		c: c,
	}, nil
}

func (c *controlServer) RequestControl(_ context.Context, req *grpcControl.ControlRequest) (*grpcControl.EmptyResponse, error) {
	c.c.SetInputValue(int(req.ParticipantId), req.Value)
	return &grpcControl.EmptyResponse{}, nil
}
