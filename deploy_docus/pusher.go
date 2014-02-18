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

func (p *Pusher) BuildCmd() *exec.Cmd {
	path, err := exec.LookPath("git")
	if err != nil {
		path = "git"
	}

	return &exec.Cmd{
		Path: path,
		Dir:  p.Repository.LocalPath(),
		Args: []string{"git", "push", p.Repository.Destination, p.Ref(), "-f"},
		Env:  []string{"GIT_SSH=script/ssh", fmt.Sprintf("PKEY=%s", p.Repository.Rsa.KeyPath())},
	}
}

func (p *Pusher) Push() error {
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
