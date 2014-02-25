package deploy_docus

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/sessions"
	"net/http"
	"strconv"
)

type Server struct {
	Port    int64
	Deploys chan Message

	*martini.ClassicMartini
}

func (c *Server) mapRoutes() {
	c.Map(&c.Deploys)

	github := BuildGitHub()
	c.Use(sessions.Sessions("deploy_docus", sessions.NewCookieStore([]byte("secret123"))))
	c.Use(oauth2.Github(&oauth2.Options{
		ClientId:     github.OauthKey,
		ClientSecret: github.OauthSecret,
		RedirectURL:  github.OauthRedirectUri,
		Scopes:       []string{""},
	}))

	c.Get("/", func(tokens oauth2.Tokens) string {
		if tokens.IsExpired() {
			return `You might be "a doctor". I am "the doctor".`
		} else {
			return `Come along Pond!`
		}
	})

	c.Post("/deploy/:id", binding.Form(Message{}), func(message Message, channel *chan Message, req *http.Request, params martini.Params) (int, string) {
		id, _ := strconv.ParseInt(params["id"], 0, 0)
		repository, err := FindRepository(id)

		if err != nil {
			return 404, ""
		}

		message.Repository = repository
		*channel <- message

		return 201, ""
	})
}

func (c *Server) Start() {
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), c)
}

func NewServer(port int64, channel chan Message) *Server {
	martini := martini.Classic()
	server := &Server{port, channel, martini}
	server.mapRoutes()
	return server
}
