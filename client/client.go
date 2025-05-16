package client

import (
	"fmt"

	v1 "github.com/centodiechi/identity/protos/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IdentityClient struct {
	conn   *grpc.ClientConn
	client v1.IdentityClient
}

func GetIdtClient() (*IdentityClient, error) {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client := v1.NewIdentityClient(conn)

	return &IdentityClient{
		conn:   conn,
		client: client,
	}, nil
}

func (c *IdentityClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
