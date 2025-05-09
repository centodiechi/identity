package main

import (
	"fmt"
	"log"
	"net"

	"github.com/centodiechi/identity/internal"
	v1 "github.com/centodiechi/identity/protos/v1"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	identityService, err := internal.NewIdentity()
	if err != nil {
		log.Fatalf("Failed to create identity service: %v", err)
	}
	log.Printf("Identity service initialized successfully")
	server := grpc.NewServer()
	v1.RegisterIdentityServer(server, identityService)
	server.Serve(lis)
}
