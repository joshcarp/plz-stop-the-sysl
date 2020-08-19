package main

import (
	"context"
	"github.com/anz-bank/sysl-go/common"
	plzserver "github.com/joshcarp/plz-stop-the-sysl/gen/pkg/servers/myserver"
	pb "github.com/joshcarp/plz-stop-the-sysl/api/plzserver"
	depserver "github.com/joshcarp/plz-stop-the-sysl/gen/pkg/servers/myserver/myserverdep"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {

	log := logrus.New()

	var port string
	if p := os.Getenv("PORT"); port != "" {
		port = p
	}

	LoadServices(log, context.Background(), ":"+port)
}

func LoadServices(log *logrus.Logger, ctx context.Context, port string) error {

	/* Service Interface for constructing the behaviour */
	serviceInterface := plzserver.GrpcServiceInterface{Hello: Hello}

	client, err := depserver.NewClient("localhost:8082", time.Second)
	/*  */
	serviceHandler := plzserver.NewGrpcServiceHandler(common.DefaultCallback(), &serviceInterface, &client)
	s := grpc.NewServer()
	reflection.Register(s)
	serviceHandler.RegisterServer(ctx, s)

	/* */
	log.Info("Starting server on ", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

func Hello(ctx context.Context, req *pb.HelloRequest, client plzserver.HelloClient) (*pb.HelloResponse, error){
	return &pb.HelloResponse{Content: "Hello"}, nil
}