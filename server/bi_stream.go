package main

import (
	pb "go-grpc/proto"
	"io"
	"log"
	"time"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	//both sides are responsible for streaming and both sides are also responsible for sending the stream and recieving the stream
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name: %v", req.Name)
		// creating response
		resp := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(resp); err != nil {
			log.Fatalf("error encountered while sending stream: %v", err)
		}
		time.Sleep(2 * time.Second)
	}

}
