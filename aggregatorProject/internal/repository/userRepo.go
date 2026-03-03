package repository

import (
	"aggregatorProject/cmd/config"
	"aggregatorProject/internal/logg"
	"aggregatorProject/internal/model"
	"context"
	"errors"
	"sync"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int, error)
	Get(ctx context.Context, ID int) (*model.User, error)
}

type UserRepo struct {
	Store  map[int]*model.User
	mx     sync.RWMutex
	cfg    config.Config
	logger logg.Logger
}

func NewUserRepo(ctx context.Context, config config.Config, logger logg.Logger) *UserRepo {
	return &UserRepo{
		Store:  make(map[int]*model.User, 10),
		cfg:    config,
		logger: logger,
	}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) (int, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.Store[user.ID] = user
	return user.ID, nil
}

func (r *UserRepo) Get(ctx context.Context, ID int) (*model.User, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	if user, ok := r.Store[ID]; ok {
		return user, nil
	}
	return nil, errors.New("user does not exist")
}
