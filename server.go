package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/aufaitio/builder/app"
	"github.com/aufaitio/builder/util"
	"github.com/docopt/docopt-go"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
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
	logger := logrus.New()

	// connect to the database
	client, err := mongo.Connect(context.Background(), buildDBHost(app.Config), nil)

	if err != nil {
		panic(fmt.Errorf("Failed to connect to MongoDB with error message: %s", err))
	}

	db := client.Database(app.Config.DB.Name)

	logger.Infof("server %v is started\n", app.Version)
	builder := app.NewBuilder(db)
	_ = util.SetIntervalAsync(builder.Check, int(app.Config.Interval))
}

func buildDBHost(config app.AppConfig) string {
	prefix := ""

	if config.DB.Username != "" {
		prefix = fmt.Sprintf("%s:%s@", config.DB.Username, config.DB.Password)
	}

	return fmt.Sprintf("mongodb://%s%s:%d", prefix, config.DB.Host, config.DB.Port)
}
