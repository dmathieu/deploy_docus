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
	Martini *martini.ClassicMartini
}

func (c *Server) mapRoutes() {
	c.Martini.Map(&c.Deploys)

	c.Martini.Get("/", func() string {
		return `You might be "a doctor". I am "the doctor".`
	})

	c.Martini.Post("/deploy", binding.Form(Message{}), func(message Message, channel *chan Message, req *http.Request) (int, string) {
		*channel <- message

		return 201, ""
	})
}

func (c *Server) Start() {
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), c.Martini)
}

func NewServer(port int, channel chan Message) *Server {
	server := &Server{Port: port, Deploys: channel, Martini: martini.Classic()}
	server.mapRoutes()
	return server
}
