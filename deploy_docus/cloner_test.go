package deploy_docus

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSuccessfulCloneBuildCmd(t *testing.T) {
	repository := BuildTestRepository()
	cloner := NewCloner(repository, `/home/tmp`)

	command := cloner.BuildCmd()

	assert.Equal(t, "/usr/bin/git", command.Path)
	assert.Equal(t, []string{"git", "clone", "git@github.com:lyonrb/deploy_docus.git", "/tmp/lyonrb_deploy_docus"}, command.Args)
	assert.Equal(t, []string{"GIT_SSH=script/ssh", "PKEY=/home/tmp/deploy_docus/keys/lyonrb_deploy_docus"}, command.Env)
}
