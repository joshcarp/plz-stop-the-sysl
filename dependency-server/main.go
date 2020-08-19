// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"github.com/anz-bank/conf-demo/dependency-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

)

// server is used to implement helloworld.GreeterServer.
type server struct {
	proto.UnimplementedItProjectServer
}

func (s *server) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("hello func ")
	return &proto.HelloResponse{Content: "Hello World"}, nil
}
var port = ":443"
func main() {
	if p := os.Getenv("PORT"); p !=""{
		port = p
	}
	if port[0] != ':'{
		port = ":"+port
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	proto.RegisterItProjectServer(s, &server{})
	fmt.Println("Starting grpc server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}