package app

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/quantumew/data-access"
	"github.com/quantumew/data-access/daos"
	"github.com/quantumew/data-access/models"
	"github.com/quantumew/plugins/lib/logger"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"os/exec"
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
func (builder *Builder) Check() {
	builder.logger.Infof("Checking build queue")
	job, err := builder.jobDAO.Claim(builder.db)

	if err != nil {
		panic("failed to query for job")
	}

	repo, err := builder.repoDAO.Get(builder.db, job.Name)

	if err != nil {
		builder.jobDAO.Release(builder.db, job)
		panic(fmt.Sprintf("failed to get repository for given job %s", job.Name))
	}

	err = builder.Spawn(job, repo)

	if err != nil {
		builder.jobDAO.Release(builder.db, job)
		panic(fmt.Sprintf(`failed to spawn job "%s" for repository "%s"`, job.Name, repo.Name))
	}
}

// Spawn spawns a build from queue
func (builder *Builder) Spawn(job *models.Job, repo *models.Repository) error {
	// Clone, update, install, healthcheck, push, open pull request
	dir := os.TempDir()
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      repo.Config.Remote,
		Progress: os.Stdout,
	})

	// npm will respect the semver version, so npm update depOne depTwo ...
	// perhaps we should provide the ability to override via .aufait.yaml...
	// Also perhaps a --registry option
	argList := []string{"update"}

	for _, dep := range job.Dependencies {
		argList = append(argList, dep.Name)
	}
	err = builder.runNpmCommand(dir, argList)

	if err != nil {
		return fmt.Errorf("Failed to run npm update, reason %s", err.Error())
	}

	// Ability to do lock only when a healthcheck is disabled
	err = builder.runNpmCommand(dir, []string{"install"})

	if err != nil {
		return fmt.Errorf("Failed to run npm install, reason %s", err.Error())
	}

	err = builder.runNpmCommand(dir, []string{"test"})

	// Open failure PR
	if err != nil {
	} else {
	}

	return err
}

func (builder *Builder) runNpmCommand(cwd string, argList []string) error {
	cmd := exec.Command("npm", argList...)
	cmd.Path = cwd
	return cmd.Run()
}

// CleanUp Cleans up builder job
func (builder *Builder) CleanUp() {
}
