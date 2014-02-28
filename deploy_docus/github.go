package deploy_docus

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/google/go-github/github"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/sessions"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type GitHub struct {
	OauthKey         string
	OauthSecret      string
	OauthRedirectUri string
	OauthAllowedId   int64
	SessionToken     []byte
}

func (g *GitHub) GetClient(token string) *github.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	return github.NewClient(t.Client())
}

func BuildGitHub() *GitHub {
	allowed_id, _ := strconv.ParseInt(os.Getenv("GITHUB_OAUTH_ALLOWED_ID"), 0, 0)

	return &GitHub{
		OauthKey:         os.Getenv("GITHUB_OAUTH_KEY"),
		OauthSecret:      os.Getenv("GITHUB_OAUTH_SECRET"),
		OauthRedirectUri: os.Getenv("GITHUB_OAUTH_REDIRECT_URI"),
		OauthAllowedId:   allowed_id,
		SessionToken:     []byte(os.Getenv("SECRET_SESSION_TOKEN")),
	}
}

var GitHubLoginRequired martini.Handler = func() martini.Handler {
	failedLogin := func(req *http.Request, writer http.ResponseWriter, s sessions.Session) {
		s.Delete("oauth2_token")
		s.Delete("isAllowed")
		next := url.QueryEscape(req.URL.RequestURI())
		http.Redirect(writer, req, oauth2.PathLogin+"?next="+next, 302)
	}

	return func(s sessions.Session, token oauth2.Tokens, c martini.Context, writer http.ResponseWriter, req *http.Request, rend render.Render) {
		isAllowed := s.Get("isAllowed")

		if token == nil || token.IsExpired() {
			failedLogin(req, writer, s)
		} else if isAllowed != true {
			hub := BuildGitHub()
			client := hub.GetClient(token.Access())
			user, _, _ := client.Users.Get("")

			if user == nil {
				failedLogin(req, writer, s)
			} else if int64(*user.ID) == hub.OauthAllowedId {
				s.Set("isAllowed", true)
			} else {
				orgs, _, _ := client.Organizations.List("", &github.ListOptions{})
				for _, org := range orgs {
					if int64(*org.ID) == hub.OauthAllowedId {
						return
					}
				}

				rend.HTML(401, "not_allowed", nil)
			}
		}
	}
}()
