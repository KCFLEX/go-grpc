package main

import (
	pb "go-grpc/proto" // Import the generated gRPC code from the "go-grpc/proto" package.
	"log"
	"net"

	"google.golang.org/grpc"
)

// define consts
const (
	port = ":8080" // This constant represents the network port on which the gRPC server will listen for incoming connections.
)

// this is the struct to be created, pb is imported upstairs
type helloServer struct { //helloServer defines a custom helloServer type, which embeds the pb.GreetServiceServer interface. This is typically used to implement the gRPC service defined in the pb package.
	pb.GreetServiceServer
}

func main() {
	// create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("failed to start server")
		panic(err)
	}
	// create new grpc server
	grpcServer := grpc.NewServer()
	// register the greet service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{}) // Registers the "helloServer" with the gRPC server.
	log.Printf("server started at %v", lis.Addr())            // start server
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil { // Start serving gRPC requests
		log.Println("failed to start")
		panic(err)
	}
}

/*This code essentially sets up and starts a gRPC server
that listens for incoming connections on a specified port.
When a client makes a gRPC request to this server,
it will be handled by the methods defined in the helloServer type.*/
