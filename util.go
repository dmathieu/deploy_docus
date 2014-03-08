package main

import (
	"fmt"
	"github.com/dmathieu/deploy_docus/deploy_docus"
	"os"
	"path"
	"path/filepath"
	"runtime"
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

func getPath() string {
	directory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	_, err := os.Stat(path.Join(directory, "templates"))

	if err != nil {
		_, filename, _, _ := runtime.Caller(1)
		directory = path.Join(path.Dir(filename))
	}

	return directory
}

func deploy(message deploy_docus.Message) {
	fmt.Println("I just received a message:", message)

	tmpPath := path.Join(getPath(), "tmp")
	output, err := deploy_docus.NewCloner(message.Repository, tmpPath).Fetch()
	if output != nil {
		fmt.Println(string(output))
	}
	if err != nil {
		fmt.Println("Couldn't clone the repository:", err)
		return
	}

	output, err = deploy_docus.NewPusher(&message, tmpPath).Push()
	if output != nil {
		fmt.Println(string(output))
	}
	if err != nil {
		fmt.Println("Couldn't push the repository:", err)
		return
	}
}
