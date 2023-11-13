package main

import (
	pb "go-grpc/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string // messages is a slice of strings. This slice is used to accumulate messages received from the client during the streaming process.
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{
				Messages: messages,
			})
			/*&pb.MessagesList{Messages: messages}: This part creates a new instance of the pb.MessagesList protobuf message. pb.MessagesList is a type generated from a Protocol Buffers definition file (usually with a .proto extension).
			The Messages field in pb.MessagesList is likely a slice of strings (based on the previous context in your code). It's used to store a list of messages that the server wants to send back to the client.
			stream.SendAndClose(...): The SendAndClose method is part of the gRPC stream interface (pb.GreetService_SayHelloClientStreamingServer). It's used to send a single response to the client and then close the stream.
			The parameter passed to SendAndClose is the message that the server wants to send to the client. In this case, it's the protobuf message &pb.MessagesList{Messages: messages}.*/
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Printf("Got request with names: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}

/*In summary, this function implements the server-side handling of a client-side streaming RPC.
It continuously receives stream(messages) from the client, accumulates the names in a slice,
and responds to the client with the accumulated list when the client indicates the end of streaming.
If any error occurs during the streaming process, it logs a fatal error and terminates the program.*/
