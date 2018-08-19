package app

import (
	"github.com/quantumew/data-access"
	"github.com/quantumew/data-access/daos"
	"github.com/quantumew/data-access/models"
	"github.com/quantumew/plugins/lib/logger"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Builder checks build queue and builds
type Builder struct {
	db      *mongo.Database
	logger  logger.Logger
	jobDAO  access.JobDAO
	repoDAO access.RepositoryDAO
}

// NewBuilder factory for builder
func NewBuilder(db *mongo.Database, logger logger.Logger) *Builder {
	jobDAO := daos.NewJobDAO()
	repoDAO := daos.NewRepositoryDAO()

	return &Builder{db, logger, jobDAO, repoDAO}
}

// Check check build queue for build
func (builder *Builder) Check() error {
	builder.logger.Infof("Checking build queue")
	jobList, err := builder.jobDAO.Query(builder.db, 0, 1)

	if err != nil {
		return err
	}

	if len(jobList) > 0 {
		job := jobList[0]
		repo, err := builder.repoDAO.GetByName(builder.db, job.Name)

		if err != nil {
			return err
		}

		err = builder.Spawn(job, repo)
	}

	return err
}

// Spawn spawns a build from queue
func (builder *Builder) Spawn(job *models.Job, repo *models.Repository) error {
}

// CleanUp Cleans up builder job
func (builder *Builder) CleanUp() {
}
