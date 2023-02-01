package server

import (
	"context"
	"fmt"

	hpb "GoStart/api"
)

type helloServer struct {
	hpb.UnimplementedGreeterServer
}

func NewHelloServer() *helloServer {
	return &helloServer{}
}

func (h *helloServer) SayHello(ctx context.Context, request *hpb.HelloRequest) (*hpb.HelloReply, error) {
	return &hpb.HelloReply{
		Message: fmt.Sprintf("Hello %s", request.Name),
	}, nil
}

func (h *helloServer) SayHelloAgain(ctx context.Context, request *hpb.HelloRequest) (*hpb.HelloReply, error) {
	return &hpb.HelloReply{
		Message: fmt.Sprintf("Hello %s again", request.Name),
	}, nil
}

var _ hpb.GreeterServer = &helloServer{}
