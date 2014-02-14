package deploy_docus

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestSuccessfulFindRepository(t *testing.T) {
	os.Setenv("REPOSITORY_ORIGIN", "git@github.com:dmathieu/deploy_docus.git")
	os.Setenv("REPOSITORY_DESTINATION", "git@heroku.com:deploy_docus.git")

	var repository *Repository
	repository = FindRepository()

	assert.Equal(t, "git@github.com:dmathieu/deploy_docus.git", repository.Origin)
	assert.Equal(t, "git@heroku.com:deploy_docus.git", repository.Destination)
}
