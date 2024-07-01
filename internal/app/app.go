package app

import (
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

}
