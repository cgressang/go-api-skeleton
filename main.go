package main

import (
	delivery "github.com/cgressang/go-api-skeleton/delivery/http"
	controller "github.com/cgressang/go-api-skeleton/delivery/http/handler"
	"github.com/cgressang/go-api-skeleton/delivery/http/interactor"
	cnf "github.com/cgressang/go-api-skeleton/infrastructure/config"
	"github.com/cgressang/go-api-skeleton/infrastructure/database"
	log "github.com/cgressang/go-api-skeleton/infrastructure/logger/zerolog"
	rt "github.com/cgressang/go-api-skeleton/infrastructure/router"
	"github.com/cgressang/go-api-skeleton/infrastructure/server"
)

func main() {
	config, err := cnf.New()
	if err != nil {
		panic(err.Error())
	}

	logger := log.New(config.Log.Debug)

	mux := rt.New()

	app := delivery.New(config, logger, mux)

	dbConn, err := database.NewMysql(
		config.Database.Host,
		config.Database.Port,
		config.Database.Net,
		config.Database.DbName,
		config.Database.User,
		config.Database.Password)

	if err != nil {
		logger.Fatalf("database failed to connect: %v", err)
	}

	defer dbConn.Close()

	app.Database = dbConn

	homeIntractor := interactor.NewHomeInteractor()

	app.Handlers["home"] = controller.NewHomeHandler(app.Logger, homeIntractor)

	for _, handler := range app.Handlers {
		handler.SetUpRoutes(app.Router)
	}

	srv := server.New(app.Router, app.Config.Server.Port)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
