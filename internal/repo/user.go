package repo

import (
	"apiserver/internal/model"
	"context"
)

type UserRepo interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserById(ctx context.Context, uid int64) (*model.User, error)
	GetUserByMobile(ctx context.Context, mobile string) (*model.User, error)
}
