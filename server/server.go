package server

import (
	log "github.com/sirupsen/logrus"

	"github.com/gordonrehling2/certavs/config"
	"github.com/gordonrehling2/certavs/service/router"
)

func Start() {
	configPath := "config.yaml"
	cfg, err := config.ReadFromConfig(&configPath)

	if err != nil {
		log.Fatalf("Couldn't read configuration file %s", err)
	}

	router := router.NewRouter(*cfg)
	router.Run()
}
