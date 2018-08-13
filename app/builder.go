package app

import (
	"github.com/aufaitio/plugins/lib/logger"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Builder checks build queue and builds
type Builder struct {
	db     *mongo.Database
	logger logger.Logger
}

// NewBuilder factory for builder
func NewBuilder(db *mongo.Database, logger logger.Logger) *Builder {
	return &Builder{db, logger}
}

// Check check build queue for build
func (builder *Builder) Check() {
	builder.logger.Infof("Checking build queue")
}

// CleanUp Cleans up builder job
func (builder *Builder) CleanUp() {
}
