package app

import (
	"aggregatorProject/cmd/config"
	"aggregatorProject/cmd/server"
	"aggregatorProject/internal/logg"
	"context"
)

type App struct {
	Layers     *Layers
	httpServer server.Server
	cfg        config.Config
	logger     logg.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.InitDeps(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) Run() error {
	return a.httpServer.Run()
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initLayers,
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

func (a *App) initLayers(ctx context.Context) error {
	var err error
	a.Layers, err = NewLayers(ctx, a.cfg, a.logger)
	return err
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
