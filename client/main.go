package main

import (
	"log"

	pb "go-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// it's almost like the client is making functional requests on the server-side that means they both need to be interacting on the same port
const (
	port = ":8080"
)

func main() {
	//client establishing connection
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) /*the grpc.Dial function is used to create a connection. The connection is assigned to the variable conn, and any potential error is assigned to the variable err.
	The second argument to grpc.Dial includes grpc.WithTransportCredentials(insecure.NewCredentials()), which specifies that the connection should use insecure credentials. In a production environment, you would typically use secure credentials for encryption, but here, it seems to be using insecure credentials, possibly for testing purposes.*/
	if err != nil {
		log.Println("didn't connect")
		panic(err)
	}
	defer conn.Close() // after making all the resquests the client connection closes

	client := pb.NewGreetServiceClient(conn)
	// these names are passed either as a regular request or as a stream request
	names := &pb.NamesList{ // we are sending this Namelist to the server
		Names: []string{"Akhil", "Alice", "Bob"},
	}
	// unary api call
	//callSayHello(client)

	// server streaming api
	//callSayHelloServerStream(client, names) // here the client makes a request to the server and the server sends a stream of responses

	// client streaming api
	//callSayHelloClientStream(client, names)

	// bidirectional streaming api
	callSayHelloBidrectionalStream(client, names)
}

/*the goal of this code is to call the grpc server with the help of grpc.dial on the same port which the grpc server is running */
