package deploy_docus

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSuccessfulCloneBuildCmd(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:lyonrb/deploy_docus.git"}
	cloner := NewCloner(repository)

	command := cloner.BuildCmd()

	assert.Equal(t, "/usr/bin/git", command.Path)
	assert.Equal(t, []string{"git", "clone", "git@github.com:lyonrb/deploy_docus.git", "/tmp/lyonrb_deploy_docus"}, command.Args)
	assert.Equal(t, []string{"GIT_SSH=script/ssh", "PKEY=/tmp/deploy_docus/keys/lyonrb_deploy_docus"}, command.Env)
}
