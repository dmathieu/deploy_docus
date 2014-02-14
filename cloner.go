package deploy_docus

import (
	"os"
	"os/exec"
)

type Cloner struct {
	*Repository
}

func (c *Cloner) Command() []string {
	return []string{"clone", c.Origin, c.LocalPath()}
}

func (c *Cloner) Fetch() error {
	err := os.RemoveAll(c.LocalPath())
	_, err = exec.Command("git", c.Command()...).Output()

	if err != nil {
		return err
	}

	return nil
}

func NewCloner(repository *Repository) *Cloner {
	return &Cloner{repository}
}
