package main

import (
	"context"
	pb "go-grpc/proto"
	"log"
	"time"
)

// unary api calls
func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.NoParam{}) // when calling client.SayHello you are sending a request which is NoParam to the server. a response is expected from the server and saved in the resp variable
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", resp.Message) // then the resp is printed here
}
