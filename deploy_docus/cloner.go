package deploy_docus

import (
	"fmt"
	"os"
	"os/exec"
)

type Cloner struct {
	*Repository
}

func (c *Cloner) BuildCmd() *exec.Cmd {
	path, err := exec.LookPath("git")
	if err != nil {
		path = "git"
	}

	return &exec.Cmd{
		Path: path,
		Args: []string{"git", "clone", c.Origin, c.LocalPath()},
		Env:  []string{"GIT_SSH=script/ssh", fmt.Sprintf("PKEY=%s", c.Repository.Rsa.KeyPath())},
	}
}

func (c *Cloner) Fetch() ([]byte, error) {
	err := os.RemoveAll(c.LocalPath())

	err = c.Repository.Rsa.WriteKey()
	if err != nil {
		return nil, err
	}

	command := c.BuildCmd()
	fmt.Println(command)
	output, err := command.Output()
	return output, err
}

func NewCloner(repository *Repository) *Cloner {
	return &Cloner{repository}
}
