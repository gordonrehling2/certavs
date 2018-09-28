package server

import (
	log "github.com/sirupsen/logrus"

	"github.com/gordonrehling2/certavs/config"
	"github.com/gordonrehling2/certavs/service/router"
	"github.com/gordonrehling2/certavs/server/db"
	"github.com/gordonrehling2/certavs/service"
)

func Start() {
	configPath := "config.yaml"
	cfg, err := config.ReadFromConfig(&configPath)

	if err != nil {
		log.Fatalf("Couldn't read configuration file %s", err)
	}

	db := db.NewPostgresDB(*cfg)
	db.Connect()

	rfeService := service.NewRfeService(db)

	router := router.NewRouter(*cfg)
	router.Run(rfeService)
}
