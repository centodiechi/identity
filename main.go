package main

import (
	"fmt"
	"log"
	"net"

	idt "github.com/centodiechi/identity/internal"
	v1 "github.com/centodiechi/identity/protos/v1"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	v1.RegisterIdentityServer(server, &idt.Identity{})
	server.Serve(lis)
}
