package app

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Builder checks build queue and builds
type Builder struct {
	db *mongo.Database
}

// NewBuilder factory for builder
func NewBuilder(db *mongo.Database) *Builder {
	return &Builder{db}
}

// Check check build queue for build
func (builder *Builder) Check() {

}
