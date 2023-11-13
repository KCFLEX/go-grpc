package main

import (
	pb "go-grpc/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error { // This is a method named SayHelloServerStreaming belonging to the helloServer type.
	//The method takes two parameters:
	//req: A parameter of type *pb.NamesList, which is a protocol buffer (protobuf) message type containing a list of names.
	//stream: A parameter of type pb.GreetService_SayHelloServerStreamingServer, which represents the server-side streaming RPC. The server uses this to send multiple responses back to the client.
	log.Printf("got request with names : %v", req.Names) // from the client this is the request that we recieved

	for _, name := range req.Names { //loop to iterate over the names received in the req parameter.
		res := &pb.HelloResponse{ // For each name, it creates a new pb.HelloResponse message, where the Message field is set to a greeting message that concatenates "Hello" with the current name.
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil { // stream.Send(res) function sends this response back to the client. If there is an error sending the response, it returns that error.
			return err
		}
		time.Sleep(2 * time.Second) // takes a 2 second pause after sending each response
	}
	return nil // Finally, after sending all the responses for each name in the loop, the function returns nil to indicate that the server-side streaming RPC has completed successfully.
}

/*this function is part of a gRPC server and implements server-side streaming.
It recieves a request with a list of names, sends a series of greeting messages back to the client for each name,
 introduces a delay between responses, and completes the streaming RPC when all responses have been sent.
 The server logs the received names before processing.*/
