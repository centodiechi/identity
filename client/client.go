package client

import (
	"fmt"

	v1 "github.com/centodiechi/identity/protos/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type identityClient struct {
	conn   *grpc.ClientConn
	client v1.IdentityClient
}

func GetIdtClient() (*identityClient, error) {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client := v1.NewIdentityClient(conn)

	return &identityClient{
		conn:   conn,
		client: client,
	}, nil
}
