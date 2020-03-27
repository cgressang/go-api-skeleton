package http

import (
	"database/sql"

	cnf "github.com/cgressang/go-api-skeleton/infrastructure/config"
	"github.com/cgressang/go-api-skeleton/usecase/handler"
	"github.com/cgressang/go-api-skeleton/usecase/log"
	"github.com/go-chi/chi"
)

type App struct {
	Config   *cnf.Config
	Handlers map[string]handler.Base
	Database *sql.DB
	Logger   log.StdLogger
	Router   *chi.Mux
}

func New(
	config *cnf.Config,
	logger log.StdLogger,
	router *chi.Mux,
) *App {
	return &App{
		Config:   config,
		Handlers: make(map[string]handler.Base),
		Logger:   logger,
		Router:   router,
	}
}
