package app

import (
	"time"

	"github.com/Sirupsen/logrus"
	log "github.com/quantumew/plugins/lib/logger"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Scope contains the application-specific information that are carried around in a request.
type Scope interface {
	log.Logger
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
	DB() *mongo.Database
}

type scope struct {
	log.Logger                 // the logger tagged with the current request information
	db         *mongo.Database // the mongo db client
}

// DB retrieve an instance of the data base
func (rs *scope) DB() *mongo.Database {
	return rs.db
}

// newScope creates a new Scope with the current request information.
func newScope(now time.Time, logger *logrus.Logger, db *mongo.Database) *scope {
	l := log.NewLogger(logger, logrus.Fields{})

	return &scope{
		Logger: l,
		db:     db,
	}
}
