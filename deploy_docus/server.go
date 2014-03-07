package deploy_docus

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/sessions"
	"net/http"
	"path"
	"strconv"
)

type Server struct {
	Port    int64
	Deploys chan Message
	Path    string

	*martini.ClassicMartini
}

func (c *Server) mapRoutes() {
	c.Map(&c.Deploys)

	c.Use(render.Renderer(render.Options{
		Layout:    "layout",
		Directory: path.Join(c.Path, "templates"),
	}))

	github := BuildGitHub()
	c.Use(sessions.Sessions("deploy_docus", sessions.NewCookieStore(github.SessionToken)))
	c.Use(oauth2.Github(&oauth2.Options{
		ClientId:     github.OauthKey,
		ClientSecret: github.OauthSecret,
		RedirectURL:  github.OauthRedirectUri,
		Scopes:       []string{"repo_deployment", "read:org"},
	}))

	c.Get("/", func(tokens oauth2.Tokens, writer http.ResponseWriter, req *http.Request) {
		if tokens.IsExpired() {
			http.Redirect(writer, req, "/login", http.StatusMovedPermanently)
		} else {
			http.Redirect(writer, req, "/repositories", http.StatusMovedPermanently)
		}
	})

	c.Post("/deploy/:id", binding.Form(Message{}), func(message Message, channel *chan Message, req *http.Request, params martini.Params) (int, string) {
		eventType := req.Header.Get("X-GitHub-Event")

		if eventType == "deployment" {
			id, _ := strconv.ParseInt(params["id"], 0, 0)
			token := req.URL.Query().Get("token")
			repository, err := FindRepository(id)

			if err != nil || repository.Token() != token {
				return 404, ""
			}

			message.Repository = repository
			*channel <- message

			return 201, ""
		} else {
			return 400, "This webhook endpoint accepts only deployments"
		}
	})

	c.Get("/repositories", GitHubLoginRequired, func(r render.Render) {
		repositories, err := AllRepositories()

		if err != nil {
			r.HTML(500, "error", err)
		} else {
			r.HTML(200, "repositories/index", repositories)
		}
	})

	c.Get("/repositories/new", GitHubLoginRequired, func(r render.Render) {
		repository := &Repository{}
		r.HTML(200, "repositories/new", repository)
	})

	c.Get("/repositories/:id", GitHubLoginRequired, func(r render.Render, params martini.Params, req *http.Request) {
		id, _ := strconv.ParseInt(params["id"], 0, 0)
		repository, err := FindRepository(id)

		if err != nil {
			r.HTML(500, "error", err)
		} else {
			values := struct {
				Repository *Repository
				Host       string
			}{repository, req.Host}

			r.HTML(200, "repositories/show", values)
		}
	})

	c.Post("/repositories", GitHubLoginRequired, func(req *http.Request, writer http.ResponseWriter, r render.Render) {
		origin := req.FormValue("repository[origin]")
		destination := req.FormValue("repository[destination]")
		repository := BuildRepository(0, origin, destination, nil)
		err := repository.Save()

		if err != nil {
			r.HTML(500, "error", err)
		} else {
			http.Redirect(writer, req, "/repositories", http.StatusMovedPermanently)
		}
	})
}

func (c *Server) Start() {
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), c)
}

func NewServer(port int64, channel chan Message, path string) *Server {
	martini := martini.Classic()
	server := &Server{port, channel, path, martini}
	server.mapRoutes()
	return server
}
