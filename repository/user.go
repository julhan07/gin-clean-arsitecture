package repository

import (
	model "clean-code/database"
	"clean-code/entity"
	"context"

	"gorm.io/gorm"
)

type userGorm struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userGorm{
		db: db,
	}
}

func (v *userGorm) GetAll(ctx context.Context) ([]*entity.User, error) {

	var users []*entity.User

	err := v.db.Model(&model.User{}).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
