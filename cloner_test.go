package deploy_docus

import (
	"github.com/bmizerany/assert"
	"net/url"
	"testing"
)

func TestSuccessfulFetch(t *testing.T) {
	url, _ := url.Parse("https://github.com/lyonrb/deploy_docus.git")
	cloner := NewCloner(url, "/tmp/deploy_docus")

	err := cloner.Fetch()

	assert.Equal(t, nil, err)
}
