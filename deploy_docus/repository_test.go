package deploy_docus

import (
	"github.com/bmizerany/assert"
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
	RemoveAllRepositories()
	var repository *Repository
	tmp, err := CreateRepository(repositoryOrigin, repositoryDestination, pemPrivateKey)
	assert.Equal(t, nil, err)

	repository, err = FindRepository(tmp.Id)

	assert.Equal(t, nil, err)
	assert.Equal(t, repositoryOrigin, repository.Origin)
	assert.Equal(t, repositoryDestination, repository.Destination)
	assert.Equal(t, pemPrivateKey, repository.Rsa.Key)

	assert.NotEqual(t, nil, repository.Rsa)
	assert.Equal(t, repository, repository.Rsa.Repository)
}

func TestFindMissingRepository(t *testing.T) {
	RemoveAllRepositories()
	repository, err := FindRepository(1)

	assert.Equal(t, (*Repository)(nil), repository)
	assert.NotEqual(t, nil, err)
}
