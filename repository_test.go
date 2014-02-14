package deploy_docus

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	assert.Equal(t, "dmathieu_deploy_docus", repository.Name())

	repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "github_hubot", repository.Name())
}

func TestLocalPath(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	assert.Equal(t, "/tmp/dmathieu_deploy_docus", repository.LocalPath())

	repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "/tmp/github_hubot", repository.LocalPath())
}

func TestSuccessfulFindRepository(t *testing.T) {
	os.Setenv("REPOSITORY_ORIGIN", "git@github.com:dmathieu/deploy_docus.git")
	os.Setenv("REPOSITORY_DESTINATION", "git@heroku.com:deploy_docus.git")

	var repository *Repository
	repository = FindRepository()

	assert.Equal(t, "git@github.com:dmathieu/deploy_docus.git", repository.Origin)
	assert.Equal(t, "git@heroku.com:deploy_docus.git", repository.Destination)
}
