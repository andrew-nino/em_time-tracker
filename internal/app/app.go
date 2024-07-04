package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	handler "github.com/andrew-nino/em_time-tracker/internal/controller/http/v1"
	repoPG "github.com/andrew-nino/em_time-tracker/internal/repository/postgresdb"
	service "github.com/andrew-nino/em_time-tracker/internal/service"
	httpserver "github.com/andrew-nino/em_time-tracker/pkg/httpserver"
	"github.com/andrew-nino/em_time-tracker/pkg/postgres"

	_ "github.com/lib/pq"

	"github.com/andrew-nino/em_time-tracker/config"
	log "github.com/sirupsen/logrus"
)

// Initialization and start of critical components.
func Run() {

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Log.Level)

	// Repositories
	log.Info("Initializing postgres...")
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	// Migrates running
	log.Info("Migrates running...")
	m := NewMigration(cfg)
	err = m.Steps(1)
	if err != nil {
		log.Fatalf("failed to migrate db: %s", err.Error())
	}

	// Services dependencies
	log.Info("Initializing services...")
	repos := repoPG.NewPGRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	// HTTP server
	log.Info("Starting http server...")
	srv := new(httpserver.Server)

	go func() {
		if err := srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print(cfg.App.Name + " Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print(cfg.App.Name + " Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err.Error())
	}

}
