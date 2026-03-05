package app

import (
	"aggregatorProject/cmd/config"
	"aggregatorProject/cmd/server"
	"aggregatorProject/internal/converter"
	"aggregatorProject/internal/handlers"
	"aggregatorProject/internal/logg"
	"aggregatorProject/internal/repository"
	"aggregatorProject/internal/response"
	"aggregatorProject/internal/service"
	"aggregatorProject/pkg/router"
	"context"
)

type App struct {
	httpServer server.Server

	ctx    context.Context
	cfg    config.Config
	logger logg.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{ctx: ctx}

	err := app.InitDeps(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) Run() error {
	// Repository
	userRepo := repository.NewUserRepo(a.cfg, a.logger)

	// Service
	userService := service.NewUserServi(a.ctx, a.cfg, a.logger, userRepo)

	// Converter
	userConverter := converter.NewUserConvert()

	// Response
	response := response.NewResp()

	// Handlers
	userHandlers := handlers.NewUserHandle(userService, userConverter, response)

	// Router
	router := a.initRoutes(router.NewRoute(), userHandlers)

	// Server
	return a.httpServer.Run(router)
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	a.httpServer = server.NewHTTPServer(ctx, a.cfg, a.logger)
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	envConf, err := config.NewEnvConf("")
	if err != nil {
		return err
	}
	a.cfg = config.NewConf(envConf)
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	a.logger = logg.NewLogg()
	return nil
}

func (a *App) initRoutes(r router.Router, userH handlers.UserHandler) router.Router {
	r.Handle("GET", "/user", userH.Read)
	r.Handle("POST", "/user", userH.Create)
	return r
}
