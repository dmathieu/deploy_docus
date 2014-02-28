package main

import (
	"fmt"
	"github.com/dmathieu/deploy_docus/deploy_docus"
)

func main() {
	port := getPort()
	channel := make(chan deploy_docus.Message)
	path := getPath()
	server := deploy_docus.NewServer(port, channel, path)

	fmt.Println("Running the server on port", server.Port)
	go server.Start()

	for {
		message := <-channel
		go deploy(message)
	}
}
