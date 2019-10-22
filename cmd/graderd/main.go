package main

import (
	"context"
	"flag"
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
	grpcAddr = flag.String("addr", "localhost", "gRPC server endpoint")
	grpcPort = flag.String("port", ":9090", "gRPC server port")
	keyFile  = flag.String("key", "certs/server.key", "private key")
	certFile = flag.String("cert", "certs/server.pem", "public cert")

	// errors
	failedCertCreation    = "failed to create cert from file"
	failedDialServer      = "failed to dial server"
	failedRegisterGateway = "failed to register gateway"
)

func serve() error {
	serverCert, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	if err != nil {
		log.Fatalln(errors.Wrap(err, failedCertCreation))
	}
	grpcServer := grpc.NewServer(grpc.Creds(serverCert))
	pb.RegisterGraderServer(grpcServer, &graderd.Service{})

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
// connections or otherHandler otherwise. Copied from cockroachdb.
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
