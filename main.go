package main

import (
	"github.com/dmathieu/deploy_docus"
	"fmt"
	"os"
	"strconv"
)

func getPort() int64 {
	str := os.Getenv("PORT")

	if str != "" {
		port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 0)
		if err != nil {
			panic(err)
		}

		return port
	} else {
		return 5000
	}
}

func deploy(message deploy_docus.Message) {
  fmt.Println("I just received a message:", message)

  err := deploy_docus.NewCloner(message.Repository).Fetch()
  if err != nil {
    panic(err)
  }

  err = deploy_docus.NewPusher(&message).Push()
  if err != nil {
    panic(err)
  }
}

func main() {
	port := getPort()
	channel := make(chan deploy_docus.Message)
	server := deploy_docus.NewServer(port, channel)

	fmt.Println("Running the server on port", server.Port)
	go server.Start()

	for {
		message := <-channel
    go deploy(message)
	}
}
