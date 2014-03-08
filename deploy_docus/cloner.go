package deploy_docus

import (
	"fmt"
	"os"
	"os/exec"
)

type Cloner struct {
	Repository *Repository
	Path       string
}

func (c *Cloner) BuildCmd() *exec.Cmd {
	path, err := exec.LookPath("git")
	if err != nil {
		path = "git"
	}

	return &exec.Cmd{
		Path: path,
		Args: []string{"git", "clone", c.Repository.Origin, c.Repository.LocalPath()},
		Env:  []string{"GIT_SSH=script/ssh", fmt.Sprintf("PKEY=%s", c.Repository.Rsa.KeyPath(c.Path))},
	}
}

func (c *Cloner) Fetch() ([]byte, error) {
	err := os.RemoveAll(c.Repository.LocalPath())

	err = c.Repository.Rsa.WriteKey(c.Path)
	if err != nil {
		return nil, err
	}

	command := c.BuildCmd()
	fmt.Println(command)
	output, err := command.Output()
	return output, err
}

func NewCloner(repository *Repository, path string) *Cloner {
	return &Cloner{repository, path}
}
