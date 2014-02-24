package main

import (
	"fmt"
	"github.com/dmathieu/deploy_docus/deploy_docus"
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
		fmt.Println("Couldn't clone the repository:", err)
		return
	}

	err = deploy_docus.NewPusher(&message).Push()
	if err != nil {
		fmt.Println("Couldn't push the repository:", err)
		return
	}
}