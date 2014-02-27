package deploy_docus

import (
	"os"
)

type GitHub struct {
	OauthKey         string
	OauthSecret      string
	OauthRedirectUri string
	SessionToken     []byte
}

func BuildGitHub() *GitHub {
	return &GitHub{
		OauthKey:         os.Getenv("GITHUB_OAUTH_KEY"),
		OauthSecret:      os.Getenv("GITHUB_OAUTH_SECRET"),
		OauthRedirectUri: os.Getenv("GITHUB_OAUTH_REDIRECT_URI"),
		SessionToken:     []byte(os.Getenv("SECRET_SESSION_TOKEN")),
	}
}
