package deploy_docus

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSuccessfulFetch(t *testing.T) {
	repository := &Repository{Origin: "https://github.com/lyonrb/deploy_docus.git"}
	cloner := NewCloner(repository)

	err := cloner.Fetch()

	assert.Equal(t, nil, err)
}
