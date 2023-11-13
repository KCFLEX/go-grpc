package main

import (
	"context"
	pb "go-grpc/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names) // we call the SayHelloServerStreaming method in the server_stream.go file to send the namesList
	if err != nil {
		log.Fatalf("could not send names : %v", err)
	}
	// the code below is to  process the stream from the server and print it out
	for { // The function then enters into an infinite loop to receive messages from the streaming RPC using stream.Recv(). The Recv function returns a message and an error.
		message, err := stream.Recv() // recieve the stream
		if err == io.EOF {            //if err = EOF means END OF FILE : it means that the stream has ended (no more messages to receive), and the loop breaks.
			break
		}
		if err != nil { // If there is any other error, the function logs the error and terminates
			log.Fatalf("error while streaming: %v", err)
		}
		log.Println(message) // Inside the loop, the received message is logged using
	}
	//After the loop, the function logs a message indicating that the streaming has finished.
	log.Printf("Streaming finished")
}

/*In summary, this function demonstrates the client-side handling of a server-side streaming RPC.
It initiates the streaming RPC, sends request to the server recieves stream from the server  and logs messages from the server until the stream ends,
and then prints a message indicating the completion of the streaming.
*/
