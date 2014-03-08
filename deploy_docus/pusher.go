package deploy_docus

import (
	"fmt"
	"os/exec"
)

type Pusher struct {
	Message *Message
	Path    string
}

func (p *Pusher) Ref() string {
	return fmt.Sprintf("%s:master", p.Message.Sha)
}

func (p *Pusher) BuildCmd() *exec.Cmd {
	path, err := exec.LookPath("git")
	if err != nil {
		path = "git"
	}

	return &exec.Cmd{
		Path: path,
		Dir:  p.Message.Repository.LocalPath(),
		Args: []string{"git", "push", p.Message.Repository.Destination, p.Ref(), "-f"},
		Env:  []string{"GIT_SSH=script/ssh", fmt.Sprintf("PKEY=%s", p.Message.Repository.Rsa.KeyPath(p.Path))},
	}
}

func (p *Pusher) Push() ([]byte, error) {

	err := p.Message.Repository.Rsa.WriteKey(p.Path)
	if err != nil {
		return nil, err
	}

	command := p.BuildCmd()
	fmt.Println(command)
	output, err := command.Output()
	return output, err
}

func NewPusher(message *Message, path string) *Pusher {
	return &Pusher{message, path}
}
