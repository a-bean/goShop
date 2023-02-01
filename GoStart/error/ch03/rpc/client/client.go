package main

//import (
//	"context"
//	"flag"
//	"fmt"
//	"log"
//	"mxshop/pkg/errors"
//	"time"
//
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	_ "mxshop/app/pkg/code"
//	pb "mxshop/cmd/order/rpc"
//)
//
//const (
//	defaultName = "world"
//)
//
//var (
//	addr = flag.String("addr", "localhost:50052", "the address to connect to")
//	name = flag.String("name", defaultName, "Name to greet")
//)
//
//func main() {
//	flag.Parse()
//	// Set up a connection to the server.
//	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := pb.NewGreeterClient(conn)
//
//	// Contact the server and print out its response.
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
//	if err != nil {
//		s := errors.FromGrpcError(err)
//		coder := errors.ParseCoder(s)
//		fmt.Println(coder.Code())
//	}
//	log.Printf("Greeting: %s", r.GetMessage())
//}
