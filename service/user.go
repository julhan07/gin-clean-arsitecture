package service

import (
	"clean-code/entity"
	"clean-code/repository"
	"context"
)

type UserService interface {
	GetAll(ctx context.Context) ([]*entity.User, error)
}

type service struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &service{
		repo: repo,
	}
}

func (v *service) GetAll(ctx context.Context) ([]*entity.User, error) {

	users, err := v.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
