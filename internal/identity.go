package internal

import (
	"context"
	"fmt"
	"log"

	userpb "github.com/centodiechi/identity/protos/v1"
	"github.com/centodiechi/store"
)

type identity struct {
	store store.Store
	userpb.UnimplementedIdentityServer
}

func NewIdentity() (*identity, error) {
	storeInstance, err := store.InitializeStore("pgsql", store.PgMeta{
		Host:         "localhost",
		Port:         "5432",
		User:         "admin",
		Password:     "admin",
		DatabaseName: "identitydb",
		TableName:    "userdb",
		CronInterval: 300,
		SslMode:      "disable",
	})
	if err != nil {
		log.Printf("Failed to initialize store: %v", err)
		return nil, err
	}
	return &identity{store: storeInstance}, nil
}

func (idt *identity) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	if idt.store == nil {
		return nil, fmt.Errorf("store is not initialized")
	}

	idt.store.Set(ctx, "user/identity/", []byte(req.Email))

	return &userpb.CreateUserResponse{
		Uid:      "1212",
		Username: req.Username,
		Email:    req.Email,
	}, nil
}
