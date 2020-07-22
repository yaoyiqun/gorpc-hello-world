package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-hello-world/proto"
	"log"
)

func Client() {
	creds, err := credentials.NewClientTLSFromFile("./certs/server.pem", "grpc.abc")
	if err != nil {
		log.Printf("Failed to create TLS credentials %v\n", err)
	}
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creds))
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := pb.NewHelloWorldClient(conn)
	context := context.Background()
	body := &pb.HelloWorldRequest{
		Referer: "Grpc",
	}

	r, err := c.SayHelloWorld(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Message)
}
