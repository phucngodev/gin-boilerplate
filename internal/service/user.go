package service

import (
	"apiserver/internal/model"
	"apiserver/internal/repo"
	"context"
	"errors"
)

var _ UserService = (*userService)(nil)

type UserService interface {
	GetByName(ctx context.Context, name string) (*model.User, error)
	GetById(ctx context.Context, uid int64) (*model.User, error)
	GetByMobile(ctx context.Context, ID string) (*model.User, error)
}

type userService struct {
	ur repo.UserRepo
}

func NewUserService(_ur repo.UserRepo) *userService {
	return &userService{

		ur: _ur,
	}
}

func (us *userService) GetByName(ctx context.Context, name string) (*model.User, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid name")
	}
	return us.ur.GetUserByName(ctx, name)
}

func (us *userService) GetById(ctx context.Context, uid int64) (*model.User, error) {
	return us.ur.GetUserById(ctx, uid)
}

func (us *userService) GetByMobile(ctx context.Context, mobile string) (*model.User, error) {
	return us.ur.GetUserByMobile(ctx, mobile)
}
