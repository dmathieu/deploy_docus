package deploy_docus

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	repository := BuildTestRepository()
	assert.Equal(t, "lyonrb_deploy_docus", repository.Name())

	repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "github_hubot", repository.Name())
}

func TestLocalPath(t *testing.T) {
	repository := BuildTestRepository()
	assert.Equal(t, "/tmp/lyonrb_deploy_docus", repository.LocalPath())

	repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "/tmp/github_hubot", repository.LocalPath())
}

func TestSuccessfulFindRepository(t *testing.T) {
	os.Setenv("REPOSITORY_ORIGIN", "git@github.com:dmathieu/deploy_docus.git")
	os.Setenv("REPOSITORY_DESTINATION", "git@heroku.com:deploy_docus.git")
	os.Setenv("REPOSITORY_PKEY", string(pemPrivateKey))

	var repository *Repository
	repository = FindRepository()

	assert.Equal(t, "git@github.com:dmathieu/deploy_docus.git", repository.Origin)
	assert.Equal(t, "git@heroku.com:deploy_docus.git", repository.Destination)
	assert.Equal(t, pemPrivateKey, repository.Rsa.Key)

	assert.NotEqual(t, nil, repository.Rsa)
	assert.Equal(t, repository, repository.Rsa.Repository)
}
