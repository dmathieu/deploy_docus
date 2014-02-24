package deploy_docus

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
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

	c.Get("/", func() string {
		return `You might be "a doctor". I am "the doctor".`
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
