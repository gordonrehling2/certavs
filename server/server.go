package server

import (
	"context"
	"fmt"
	"io"
	"os"

	"database/sql"
	"github.com/sirupsen/logrus"

	"github.com/JSainsburyPLC/issasdk/httplog"
	"github.com/JSainsburyPLC/issasdk/logging"

	"github.com/gordonrehling2/certavs/config"
	"github.com/gordonrehling2/certavs/service"
	"github.com/JSainsburyPLC/api-r10-client/errors"
)

const defaultConfigPath = "config.yaml"

var ctx = context.Background()

func Start() int {
	fmt.Printf("Server certavs starting...\n")

	serverConfigFile := defaultConfigPath

	config := setUp(ctx, &serverConfigFile)
	logger := newHttpLogger(config)
	service.Start(ctx, config.Api.Port, logger)
	return 0
}

// Returns accessLogger
func setUpLogging(config *config.Config) {
	logrus.SetFormatter(&logging.LogstashFormatter{})
	logrus.SetOutput(logwriter(config.Logs.App))
	log_level_unparsed := config.Logs.AppLevel
	if log_level_unparsed == "" {
		log_level_unparsed = "info"
	}
	log_level_parsed, err := logrus.ParseLevel(log_level_unparsed)
	if err != nil {
		fmt.Printf("Setting Logrus log level failed: %v\n", err)
		os.Exit(1)
	}
	logrus.SetLevel(log_level_parsed)
	w := logrus.StandardLogger().Writer()
	defer w.Close()
}

// create logger for access logging
func newHttpLogger(config *config.Config) httplog.Logger {
	logger := httplog.New(logwriter(config.Logs.Access))
	return logger
}

func getDatabaseConnection(ctx context.Context, config *config.Config) *sql.DB {
	dbConnection, err := sql.Open("postgres", config.DB.BuildConnectionURL())
	if err != nil {
		errors.LogError(ctx, errors.ErrDatabaseConnectionError, err)

	}

	return dbConnection
}

func setUp(ctx context.Context, configFile *string) (*config.Config) {
	c, err := config.ReadFromConfig(configFile)
	if err != nil {
		errors.LogFatal(ctx, errors.ErrFailedToReadConfigFile, err)
	}
	setUpLogging(c)
	return c
}

func logwriter(file string) io.Writer {
	w, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logrus.WithField("file", file).Fatal("Unable to open log file for writing")
	}
	return w
}
