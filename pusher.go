package deploy_docus

import (
	"fmt"
	"os/exec"
)

type Pusher struct {
	*Message
}

func (c *Pusher) Ref() string {
	return fmt.Sprintf("%s:master", c.Sha)
}

func (c *Pusher) Command() []string {
	return []string{"push", c.Repository.Destination, c.Ref(), "-f"}
}

func (c *Pusher) Fetch() error {
	_, err := exec.Command("git", c.Command()...).Output()

	if err != nil {
		return err
	}

	return nil
}

func NewPusher(message *Message) *Pusher {
	return &Pusher{message}
}
