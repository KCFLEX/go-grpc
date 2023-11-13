package main

import (
	"context"
	pb "go-grpc/proto"
)

// unary server
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) { // when you call this function its supposed to return a message called hello and this function is called by the client. after the client gets this message its prints it out.
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
