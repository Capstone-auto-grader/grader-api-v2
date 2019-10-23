// client.go generates a gRPC client using the protobuf stub,
// the server address, and the public cert.
package main

import (
	"context"
	"log"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewClient(certFile, addr string) pb.GraderClient {
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
