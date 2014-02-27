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

func TestIsNew(t *testing.T) {
	repository := BuildTestRepository()
	assert.Equal(t, true, repository.IsNew())

	repository.Id = 42
	assert.Equal(t, false, repository.IsNew())
}

func TestSuccessfulSave(t *testing.T) {
	RemoveAllRepositories()

	repository := BuildTestRepository()

	assert.Equal(t, true, repository.IsNew())
	err := repository.Save()
	assert.Equal(t, nil, err)
	assert.Equal(t, false, repository.IsNew())

	tmp, _ := FindRepository(repository.Id)
	assert.Equal(t, repository.Origin, tmp.Origin)
	assert.Equal(t, repository.Destination, tmp.Destination)

	repository.Origin = "git@github.com:lyonrb/biceps.git"
	repository.Destination = "git@heroku.com:biceps.git"
	err = repository.Save()
	assert.Equal(t, nil, err)

	tmp, _ = FindRepository(repository.Id)
	assert.Equal(t, repository.Origin, tmp.Origin)
	assert.Equal(t, repository.Destination, tmp.Destination)
}

func TestSuccessfulFindRepository(t *testing.T) {
	RemoveAllRepositories()
	var repository *Repository

	tmp := BuildTestRepository()
	err := tmp.Save()
	assert.Equal(t, nil, err)

	repository, err = FindRepository(tmp.Id)

	assert.Equal(t, nil, err)
	assert.Equal(t, repositoryOrigin, repository.Origin)
	assert.Equal(t, repositoryDestination, repository.Destination)

	assert.NotEqual(t, nil, repository.Rsa)
	assert.Equal(t, repository, repository.Rsa.Repository)
}

func TestFindMissingRepository(t *testing.T) {
	RemoveAllRepositories()
	repository, err := FindRepository(1)

	assert.Equal(t, (*Repository)(nil), repository)
	assert.NotEqual(t, nil, err)
}

func TestFindAll(t *testing.T) {
	RemoveAllRepositories()

	tmp := BuildTestRepository()
	_ = tmp.Save()
	tmp = BuildTestRepository()
	_ = tmp.Save()

	repositories, err := AllRepositories()

	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(repositories))

	first := repositories[0]
	assert.Equal(t, repositoryOrigin, first.Origin)
	assert.Equal(t, repositoryDestination, first.Destination)

	second := repositories[0]
	assert.Equal(t, repositoryOrigin, second.Origin)
	assert.Equal(t, repositoryDestination, second.Destination)
}
