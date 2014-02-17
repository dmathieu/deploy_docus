package deploy_docus

import (
	"os"
	"os/exec"
)

type Cloner struct {
	*Repository
}

func (c *Cloner) Command() []string {
	return []string{"git", "clone", c.Origin, c.LocalPath()}
}

func (c *Cloner) BuildCmd() *exec.Cmd {
	path, err := exec.LookPath("git")
	if err != nil {
		path = "git"
	}

	return &exec.Cmd{
		Path: path,
		Args: c.Command(),
	}
}

func (c *Cloner) Fetch() error {
	err := os.RemoveAll(c.LocalPath())

	command := c.BuildCmd()
	_, err = command.Output()

	if err != nil {
		return err
	}

	return nil
}

func NewCloner(repository *Repository) *Cloner {
	return &Cloner{repository}
}
