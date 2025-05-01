package internal

import (
	"context"
	"time"

	userpb "github.com/centodiechi/identity/protos/v1"
	"github.com/centodiechi/store"
)

type Identity struct {
	store store.Store
	userpb.UnimplementedIdentityServer
}

func (idt *Identity) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	if idt.store == nil {
		sp, err := store.InitializeStore("pgsql", store.PgMeta{
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
			return nil, err
		}
		idt.store = sp
	}

	idt.store.SetWithTTL(ctx, "user/1212/email", []byte(req.Email), 3*time.Hour)

	return &userpb.CreateUserResponse{
		Uid:      "1212",
		Username: req.Username,
		Email:    req.Email,
	}, nil
}
