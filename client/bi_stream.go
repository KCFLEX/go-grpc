package main

import (
	"context"
	pb "go-grpc/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidrectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Bidirectional Streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names : %v", err)
	}
	// to send stream and recieve a stream we have to use channels
	// to recieve stream
	waitc := make(chan struct{})

	go func() {
		for {
			respStream, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streaming : %v", err)
			}
			log.Println(respStream)
		}
		close(waitc)
	}()
	// to send stream
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending stream : %v", err)
		}
		log.Printf("send request with name : %s", name)
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Println("Bidirectional streaming ended")
}
