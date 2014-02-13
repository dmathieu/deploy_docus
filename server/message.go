package server

type Message struct {
	Id  int    `form:"id"`
	Sha string `form:"sha"`
}
