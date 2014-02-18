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

func (c *Cloner) Fetch() error {
	err := os.RemoveAll(c.LocalPath())

	command := c.BuildCmd()
	fmt.Println(command)
	_, err = command.Output()

	if err != nil {
		return err
	}

	return nil
}

func NewCloner(repository *Repository) *Cloner {
	return &Cloner{repository}
}
