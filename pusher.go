package deploy_docus

import (
	"fmt"
	"os/exec"
)

type Pusher struct {
	*Message
}

func (p *Pusher) Ref() string {
	return fmt.Sprintf("%s:master", p.Sha)
}

func (p *Pusher) Command() []string {
	return []string{"git", "push", p.Repository.Destination, p.Ref(), "-f"}
}

func (p *Pusher) BuildCmd() *exec.Cmd {
	path, err := exec.LookPath("git")
	if err != nil {
		path = "git"
	}

	return &exec.Cmd{
		Path: path,
		Dir:  p.Repository.LocalPath(),
		Args: p.Command(),
	}
}

func (p *Pusher) Fetch() error {
	command := p.BuildCmd()
	_, err := command.Output()

	if err != nil {
		return err
	}

	return nil
}

func NewPusher(message *Message) *Pusher {
	return &Pusher{message}
}
