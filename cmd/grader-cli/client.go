package main

import (
	"context"
	"log"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewClient generates a gRPC client with the server address, and the public cert.
func NewClient(certFile string, addr string) pb.GraderClient {
	ctx := context.Background()

	clientCert, err := credentials.NewClientTLSFromFile(certFile, addr)
	if err != nil {
		log.Fatalln("failed to read cert:", err)
	}

	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}

	return pb.NewGraderClient(conn)
}
