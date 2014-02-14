package deploy_docus

import (
	"os"
	"os/exec"
)

type Cloner struct {
	*Repository
}

func (c *Cloner) Path() string {
	return "/tmp/deploy_docus"
}

func (c *Cloner) Fetch() error {
	err := os.RemoveAll(c.Path())
	_, err = exec.Command("git", "clone", c.Origin, c.Path()).Output()

	if err != nil {
		return err
	}

	return nil
}

func NewCloner(repository *Repository) *Cloner {
	return &Cloner{repository}
}
