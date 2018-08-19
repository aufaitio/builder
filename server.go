package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/quantumew/builder/app"
	"github.com/quantumew/builder/util"
	log "github.com/quantumew/plugins/lib/logger"
	"github.com/docopt/docopt-go"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"time"
)

func main() {
	doc := `Au Fait
Command line interface for starting the builder micro service for Au Fait.

Usage:
	server [--configPath=<path>]
Options:
	-h --help				Show this message
	--version				Show version info
	--configPath=<path>   	Path to app.yaml config file [default: config]`

	arguments, _ := docopt.ParseDoc(doc)
	configPath := arguments["--configPath"].(string)

	// load application configurations
	if err := app.LoadConfig(configPath); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// create the logger
	logger := log.NewLogger(logrus.New(), logrus.Fields{})

	// connect to the database
	client, err := mongo.Connect(context.Background(), buildDBHost(app.Config), nil)

	if err != nil {
		panic(fmt.Errorf("Failed to connect to MongoDB with error message: %s", err))
	}

	db := client.Database(app.Config.DB.Name)

	logger.Infof("Server builder@%v is started\n", app.Version)
	builder := app.NewBuilder(db, logger)
	duration := time.Duration(app.Config.IntervalMinutes) * time.Minute
	clear := util.SetIntervalAsync(builder.Check, duration)

	// Clean up on sigint
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logger.Infof("Server builder@%v is going down\n", app.Version)
			clear <- true
			builder.CleanUp()
			os.Exit(0)
		}
	}()
	<-clear
}

func buildDBHost(config app.AppConfig) string {
	prefix := ""

	if config.DB.Username != "" {
		prefix = fmt.Sprintf("%s:%s@", config.DB.Username, config.DB.Password)
	}

	return fmt.Sprintf("mongodb://%s%s:%d", prefix, config.DB.Host, config.DB.Port)
}
