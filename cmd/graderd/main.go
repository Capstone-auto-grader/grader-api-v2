package main

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/Capstone-auto-grader/grader-api-v2/certs"
	pb "github.com/Capstone-auto-grader/grader-api-v2/graderpb"
	"github.com/Capstone-auto-grader/grader-api-v2/internal/graderd"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcAddr = flag.String("grpc-addr", "localhost", "gRPC server endpoint")
	grpcPort = flag.String("grpc-port", ":9090", "gRPC server port")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcEndpoint := *grpcAddr + *grpcPort

	// Read certs
	keyPair, certPool, err := parseCert()
	if err != nil {
		return err
	}
	// Create gRPC server
	serverOpts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(certPool, grpcEndpoint))}
	server := grpc.NewServer(serverOpts...)
	pb.RegisterGraderServer(server, &graderd.Service{})

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: grpcEndpoint,
		RootCAs:    certPool,
	})
	opts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	mux := runtime.NewServeMux()
	err = pb.RegisterGraderHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Print("graderd is up")
	//return http.ListenAndServe(":8081", mux)
	conn, err := net.Listen("tcp", *grpcPort)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    grpcEndpoint,
		Handler: grpcHandlerFunc(server),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*keyPair},
			NextProtos:   []string{"h2"},
		},
	}

	return srv.Serve(tls.NewListener(conn, srv.TLSConfig))
}

func grpcHandlerFunc(grpcServer *grpc.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		grpcServer.ServeHTTP(w, r)
	})
}

func parseCert() (*tls.Certificate, *x509.CertPool, error) {
	var err error
	pair, err := tls.X509KeyPair([]byte(certs.Cert), []byte(certs.Key))
	if err != nil {
		return nil, nil, err
	}

	keyPair := &pair
	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM([]byte(certs.Cert))
	if !ok {
		return nil, nil, errors.New("invalid certs")
	}

	return keyPair, certPool, nil
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
