package deploy_docus

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestGetClient(t *testing.T) {
	github := BuildGitHub()
	client := github.GetClient("hello world")
	assert.NotEqual(t, nil, client)
}

func TestBuildGitHub(t *testing.T) {
	os.Setenv("GITHUB_OAUTH_KEY", "hello")
	os.Setenv("GITHUB_OAUTH_SECRET", "world")
	os.Setenv("GITHUB_OAUTH_REDIRECT_URI", "http://example.com")
	os.Setenv("GITHUB_OAUTH_ALLOWED_ID", "9347")
	os.Setenv("SECRET_SESSION_TOKEN", "1234")

	github := BuildGitHub()

	assert.Equal(t, "hello", github.OauthKey)
	assert.Equal(t, "world", github.OauthSecret)
	assert.Equal(t, "http://example.com", github.OauthRedirectUri)
	assert.Equal(t, int64(9347), github.OauthAllowedId)
	assert.Equal(t, []byte("1234"), github.SessionToken)
}
