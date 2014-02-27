package deploy_docus

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestBuildGitHub(t *testing.T) {
	os.Setenv("GITHUB_OAUTH_KEY", "hello")
	os.Setenv("GITHUB_OAUTH_SECRET", "world")
	os.Setenv("GITHUB_OAUTH_REDIRECT_URI", "http://example.com")
	os.Setenv("SECRET_SESSION_TOKEN", "1234")

	github := BuildGitHub()

	assert.Equal(t, "hello", github.OauthKey)
	assert.Equal(t, "world", github.OauthSecret)
	assert.Equal(t, "http://example.com", github.OauthRedirectUri)
	assert.Equal(t, []byte("1234"), github.SessionToken)
}
