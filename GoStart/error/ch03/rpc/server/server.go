package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mxshop/app/pkg/code"
	"mxshop/pkg/errors"
	"net"

	"google.golang.org/grpc"
	pb "mxshop/cmd/order/rpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	e := errors.WithCode(code.ErrUserNotFound, "user not found")
	return nil, errors.ToGrpcError(e)
	//return nil, status.Error(codes.NotFound, "user not found")
	//return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
