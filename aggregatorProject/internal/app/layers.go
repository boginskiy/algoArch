package app

import (
	"aggregatorProject/cmd/config"
	"aggregatorProject/internal/handlers"
	"aggregatorProject/internal/logg"
	"aggregatorProject/internal/repository"
	"aggregatorProject/internal/service"
	"context"
)

type Layers struct {
	userRepository repository.UserRepository
	userService    service.UserService
	userHandlers   handlers.UserHandler

	cfg    config.Config
	logger logg.Logger
}

func NewLayers(ctx context.Context, config config.Config, logger logg.Logger) (*Layers, error) {
	l := &Layers{
		cfg:    config,
		logger: logger,
	}

	err := l.InitDeps(ctx)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (l *Layers) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		l.initUserRepository,
		l.initUserService,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Layers) initUserRepository(ctx context.Context) error {
	l.userRepository = repository.NewUserRepo(ctx, l.cfg, l.logger)
	return nil
}

func (l *Layers) initUserService(ctx context.Context) error {
	l.userService = service.NewUserService(ctx, l.cfg, l.logger, l.userRepository)
	return nil
}

func (l *Layers) initUserHandlers(ctx context.Context) error {
	l.userHandlers = 
}
