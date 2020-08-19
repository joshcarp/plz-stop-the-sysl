// Code generated by sysl DO NOT EDIT.
package depserver

import (
	"context"
	"time"

	pb "api.proto"

	"google.golang.org/grpc"
)

// Client implements a client for myserverdep.
type Client struct {
	client pb.myserverdepClient
	addr   string
}

// NewClient creates a new Client.
func NewClient(addr string, connTimeout time.Duration) (*Client, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctxWithTimeout, addr, grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Client{pb.NewmyserverdepClient(conn), addr}, nil
}

// Hello ...
func (s *Client) Hello(ctx context.Context, input *pb.HelloRequest) (*pb.HelloResponse, error) {
	return s.client.Hello(ctx, input)
}
