package mysql

import (
	"apiserver/internal/model"
	"apiserver/internal/repo"
	"context"

	"gorm.io/gorm"
)

var _ repo.UserRepo = (*userRepo)(nil)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.Where("name = ?", name).Find(user).Error
	return user, err
}

func (ur *userRepo) GetUserById(ctx context.Context, uid int64) (*model.User, error) {
	user := &model.User{}
	err := ur.db.Where("id = ?", uid).Find(user).Error
	return user, err
}

func (ur *userRepo) GetUserByMobile(ctx context.Context, mobile string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.
		Where("mobile = ?", mobile).
		Where("enabled_status = 1").
		First(user).Error
	return user, err
}
