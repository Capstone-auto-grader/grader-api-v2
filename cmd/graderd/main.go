// graderd is the entry-point for starting the grader API service.
//
// This file creates a gRPC endpoint and a http endpoint that uses the same given port.
//
// Incoming http GET/POST requests are translated into gRPC calls and reverse-proxied
// to the gRPC handler.
//
// gRPC calls are directly handled by the gRPC handler.
//
// Usage
//
// 		go build -o *.go graderd
//		./graderd
package main

import (
	"context"
	"flag"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/docker-client"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/graderd"
)

var (
	// command-line flags
	grpcAddr      = flag.String("GRPC_ADDR", "localhost", "gRPC server endpoint")
	grpcPort      = flag.String("GRPC_PORT", ":9090", "gRPC server port")
	dockerAddr    = flag.String("DOCKER_ADDR", "http://localhost:2376", "docker host endpoint")
	dockerVersion = flag.String("DOCKER_VERSION", "1.40", "docker host version")
	databaseAddr  = flag.String("DATABASE_ADDR", "localhost:3306", "database endpoint")
	webAddr       = flag.String("WEB_ADDR", "localhost:8080", "web API server endpoint")
	keyFile       = flag.String("KEY", "certs/server.key", "private key")
	certFile      = flag.String("CERT", "certs/server.pem", "public cert")

	// errors
	failedCertCreation    = "failed to create cert from file"
	failedDialServer      = "failed to dial server"
	failedRegisterGateway = "failed to register gateway"
)

// serve creates and runs the gRPC and http server.
func serve() error {
	// create gRPC server
	serverCert, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	if err != nil {
		log.Fatalln(errors.Wrap(err, failedCertCreation))
	}
	grpcServer := grpc.NewServer(grpc.Creds(serverCert))
	graderService := graderd.NewGraderService(docker_client.NewDockerClient(*dockerAddr, *dockerVersion), graderd.NewPGDatabase(*databaseAddr), *webAddr)
	pb.RegisterGraderServer(grpcServer, graderService)

	endpoint := *grpcAddr + *grpcPort
	clientCert, err := credentials.NewClientTLSFromFile(*certFile, endpoint)
	if err != nil {
		log.Fatalln(errors.Wrap(err, failedCertCreation))
	}
	conn, err := grpc.DialContext(
		context.Background(),
		endpoint,
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, failedDialServer))
	}

	router := runtime.NewServeMux()
	if err = pb.RegisterGraderHandler(context.Background(), router, conn); err != nil {
		log.Fatalln(errors.Wrap(err, failedRegisterGateway))
	}

	return http.ListenAndServeTLS(*grpcPort, *certFile, *keyFile, grpcHandlerFunc(grpcServer, router))
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or httpHandler otherwise. This block is copied from cockroachdb.
func grpcHandlerFunc(grpcHandler *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcHandler.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	flag.Parse()

	if err := serve(); err != nil {
		log.Fatal(err)
	}
}
