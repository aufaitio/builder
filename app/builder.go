package app

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/quantumew/data-access/daos"
	"github.com/quantumew/data-access/models"
	"github.com/quantumew/plugins/lib/logger"
)

// Builder checks build queue and builds
type Builder struct {
	db      *mongo.Database
	logger  logger.Logger
	jobDAO  *daos.JobDAO
	repoDAO *daos.RepositoryDAO
}

// NewBuilder factory for builder
func NewBuilder(db *mongo.Database, logger logger.Logger) *Builder {
	jobDAO := daos.NewJobDAO()
	repoDAO := daos.NewRepositoryDAO()

	return &Builder{db, logger, jobDAO, repoDAO}
}

// Check check build queue for build
func (builder *Builder) Check() {
	builder.logger.Infof("Checking build queue")
	jobList, err := builder.jobDAO.Query(builder.db, 0, 1)

	if err != nil {
		panic("failed to query for job")
	}

	if len(jobList) > 0 {
		job := jobList[0]
		repo, err := builder.repoDAO.Get(builder.db, job.Name)

		if err != nil {
			panic(fmt.Sprintf("failed to get repository for given job %s", job.Name))
		}

		err = builder.Spawn(job, repo)

		if err != nil {
			panic(fmt.Sprintf(`failed to spawn job "%s" for repository "%s"`, job.Name, repo.Name))
		}
	}
}

// Spawn spawns a build from queue
func (builder *Builder) Spawn(job *models.Job, repo *models.Repository) error {
	return nil
}

// CleanUp Cleans up builder job
func (builder *Builder) CleanUp() {
}
