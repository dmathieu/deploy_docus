package main

import (
	"fmt"
	"github.com/dmathieu/deploy_docus/deploy_docus"
	"os"
	"path/filepath"
)

func main() {
	port := getPort()
	channel := make(chan deploy_docus.Message)
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	server := deploy_docus.NewServer(port, channel, path)

	fmt.Println("Running the server on port", server.Port)
	go server.Start()

	for {
		message := <-channel
		go deploy(message)
	}
}
