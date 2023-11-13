package main

import (
	"context"
	pb "go-grpc/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

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
	// you get a response
	res, err := stream.CloseAndRecv()
	log.Println("client streaming finished")
	if err != nil {
		log.Fatalf("could not recieve: %v", err)
	}
	log.Printf("%v", res.Messages)

}
