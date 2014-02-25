package deploy_docus

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
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
	c.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

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

	c.Get("/repositories", oauth2.LoginRequired, func(r render.Render) {
		repositories, _ := AllRepositories()

		r.HTML(200, "repositories/index", repositories)
	})

	c.Get("/repositories/new", oauth2.LoginRequired, func(r render.Render) {
		repository := &Repository{}
		r.HTML(200, "repositories/new", repository)
	})

	c.Post("/repositories", oauth2.LoginRequired, func(req *http.Request, writer http.ResponseWriter) {
		origin := req.FormValue("repository[origin]")
		destination := req.FormValue("repository[destination]")
		key := []byte(req.FormValue("repository[key]"))
		repository := BuildRepository(0, origin, destination, key)
		repository.Save()

		http.Redirect(writer, req, "/", http.StatusMovedPermanently)
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
