# go-grpc
This is a demonstration application showcasing the capabilities of gRPC (Google Remote Procedure Call) in a client-server architecture. 
The application is split into a client and a server, utilizing gRPC for communication.

# Overview
The app's primary goal is to demonstrate different types of RPC (Remote Procedure Call) interactions between the client and server. 
The supported communication patterns include:

# Unary API:
The client sends a single request to the server.
The server responds with a greeting for the provided names.

# Server-side Streaming:
The client sends a request to the server with a list of names.
The server responds with a stream of greetings for each name.

# Client-side Streaming:
The client streams a series of names to the server.
The server responds with a single greeting for the received names.

# Bidirectional Streaming:
Both the client and server send a stream of messages to each other concurrently.
The server responds to each name sent by the client with a greeting.
