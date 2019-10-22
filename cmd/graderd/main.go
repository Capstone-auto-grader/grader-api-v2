package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/graderd"
)

var (
	// command-line flags:
	grpcAddr = flag.String("addr", "localhost", "gRPC server endpoint")
	grpcPort = flag.String("port", ":9090", "gRPC server port")
	keyFile  = flag.String("key", "certs/server.key", "private key")
	certFile = flag.String("cert", "certs/server.pem", "public cert")
)

func serve() error {
	serverCert, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	if err != nil {
		log.Fatalln("failed to create cert:", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(serverCert))
	pb.RegisterGraderServer(grpcServer, &graderd.Service{})

	endpoint := *grpcAddr + *grpcPort
	clientCert, err := credentials.NewClientTLSFromFile(*certFile, endpoint)
	if err != nil {
		log.Fatalln("failed to create cert:", err)
	}
	conn, err := grpc.DialContext(
		context.Background(),
		endpoint,
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	router := runtime.NewServeMux()
	if err = pb.RegisterGraderHandler(context.Background(), router, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	return http.ListenAndServeTLS(*grpcPort, *certFile, *keyFile, grpcHandlerFunc(grpcServer, router))
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	flag.Parse()

	if err := serve(); err != nil {
		log.Fatal(err)
	}
}
