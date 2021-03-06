package deploy_docus

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSuccessfulPushBuildCmd(t *testing.T) {
	repository := BuildTestRepository()
	message := &Message{Repository: repository, Sha: "fabfab"}
	pusher := NewPusher(message, `/home/tmp`)

	command := pusher.BuildCmd()

	assert.Equal(t, "/usr/bin/git", command.Path)
	assert.Equal(t, "/tmp/lyonrb_deploy_docus", command.Dir)
	assert.Equal(t, []string{"git", "push", "git@heroku.com:deploy_docus.git", "fabfab:master", "-f"}, command.Args)
	assert.Equal(t, []string{"GIT_SSH=script/ssh", "PKEY=/home/tmp/deploy_docus/keys/lyonrb_deploy_docus"}, command.Env)
}
