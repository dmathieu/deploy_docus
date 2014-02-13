package deploy_docus

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
)

type Cloner struct {
	Url  *url.URL
	Path string
}

func (c *Cloner) Fetch() error {
	err := os.RemoveAll(c.Path)
	b, err := exec.Command("git", "clone", c.Url.String(), c.Path).Output()

	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func NewCloner(url *url.URL, path string) *Cloner {
	return &Cloner{Url: url, Path: path}
}
