package server

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"net/http"
)

type Server struct {
	Port    int
	Deploys chan Message

	*martini.ClassicMartini
}

func (c *Server) mapRoutes() {
	c.Map(&c.Deploys)

	c.Get("/", func() string {
		return `You might be "a doctor". I am "the doctor".`
	})

	c.Post("/deploy", binding.Form(Message{}), func(message Message, channel *chan Message, req *http.Request) (int, string) {
		*channel <- message

		return 201, ""
	})
}

func (c *Server) Start() {
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), c)
}

func NewServer(port int, channel chan Message) *Server {
	martini := martini.Classic()
	server := &Server{port, channel, martini}
	server.mapRoutes()
	return server
}
