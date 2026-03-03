package service

import (
	"aggregatorProject/cmd/config"
	"aggregatorProject/internal/logg"
	"aggregatorProject/internal/model"
	"aggregatorProject/internal/repository"
	"context"
)

type UserService interface {
	Create(ID int) *model.User
	Get(ID int) *model.User
}

type UserServi struct {
	Repo   repository.UserRepository
	cfg    config.Config
	logger logg.Logger
}

func NewUserService(
	ctx context.Context,
	config config.Config,
	logger logg.Logger,
	repo repository.UserRepository) *UserServi {

	return &UserServi{
		Repo:   repo,
		cfg:    config,
		logger: logger,
	}
}

func (u *UserServi) Create(ID int) *model.User {
	return nil
}

func (u *UserServi) Get(ID int) *model.User {
	return nil
}
