package deploy_docus

import (
	"github.com/bmizerany/assert"
	"io/ioutil"
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

func TestCreatePKeyFile(t *testing.T) {
	os.RemoveAll("/tmp/deploy_docus/keys")
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git", PKey: "private_key_content"}

	_, err := os.Stat(repository.PKeyPath())
	assert.Equal(t, true, os.IsNotExist(err))

	repository.CreatePKey()

	_, err = os.Stat(repository.PKeyPath())
	assert.Equal(t, false, os.IsNotExist(err))

	content, err := ioutil.ReadFile(repository.PKeyPath())
	assert.Equal(t, "private_key_content", string(content))
}

func TestPKeyPath(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	assert.Equal(t, "/tmp/deploy_docus/keys/dmathieu_deploy_docus", repository.PKeyPath())

	repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "/tmp/deploy_docus/keys/github_hubot", repository.PKeyPath())
}

func TestSuccessfulFindRepository(t *testing.T) {
	os.Setenv("REPOSITORY_ORIGIN", "git@github.com:dmathieu/deploy_docus.git")
	os.Setenv("REPOSITORY_DESTINATION", "git@heroku.com:deploy_docus.git")
	os.Setenv("REPOSITORY_PKEY", "private_repository_ssh_key_content")

	var repository *Repository
	repository = FindRepository()

	assert.Equal(t, "git@github.com:dmathieu/deploy_docus.git", repository.Origin)
	assert.Equal(t, "git@heroku.com:deploy_docus.git", repository.Destination)
	assert.Equal(t, "private_repository_ssh_key_content", repository.PKey)
}
