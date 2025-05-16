package internal

import (
	"context"
	"fmt"
	"log"

	userpb "github.com/centodiechi/identity/protos/v1"
	"github.com/centodiechi/identity/utils"
	"github.com/centodiechi/store"
)

type identity struct {
	store map[string]store.Store
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
	storeInstanceDB, err := store.InitializeStore("redis", store.RedisMeta{
		Host:     "localhost",
		Port:     "6379",
		Password: "",
		DB:       0,
	})
	if err != nil {
		log.Printf("Failed to initialize store: %v", err)
		return nil, err
	}
	return &identity{store: map[string]store.Store{
		"pgsql": storeInstance,
		"redis": storeInstanceDB,
	},
	}, nil
}

func (idt *identity) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	if idt.store == nil {
		return nil, fmt.Errorf("store is not initialized")
	}

	uid, err := idt.store["redis"].(*store.RedisStore).GetNextID(ctx)
	if err != nil {
		log.Printf("Failed to get next ID: %v", err)
		return nil, err
	}

	hashedPassword, err := utils.GetPasswordHash(req.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, err
	}

	idt.store["pgsql"].Set(ctx, fmt.Sprintf("identity/user/%s/email", uid), []byte(req.Email))
	idt.store["pgsql"].Set(ctx, fmt.Sprintf("identity/user/%s/password", uid), []byte(hashedPassword))
	idt.store["pgsql"].Set(ctx, fmt.Sprintf("identity/user/%s/username", uid), []byte(req.Username))

	return &userpb.CreateUserResponse{
		Uid:      uid,
		Username: req.Username,
		Email:    req.Email,
	}, nil
}
