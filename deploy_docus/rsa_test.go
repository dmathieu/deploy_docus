package deploy_docus

import (
	"github.com/bmizerany/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestSuccessfulEncrypt(t *testing.T) {
	key, _ := BuildPrivateKey(pemPrivateKey)
	rsa := &Rsa{Private: key}

	encrypted, err := rsa.Encrypt([]byte("hello"))

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, encrypted)
}

func TestNewRsa(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	rsa := NewRsa(repository, pemPrivateKey)

	assert.Equal(t, repository, rsa.Repository)
	assert.NotEqual(t, nil, rsa.Private)
}

func TestCreateKeyFile(t *testing.T) {
	os.RemoveAll("/tmp/deploy_docus/keys")

	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	rsa := &Rsa{Repository: repository, Key: pemPrivateKey}

	_, err := os.Stat(rsa.KeyPath())
	assert.Equal(t, true, os.IsNotExist(err))

	rsa.WriteKey()

	_, err = os.Stat(rsa.KeyPath())
	assert.Equal(t, false, os.IsNotExist(err))

	content, err := ioutil.ReadFile(rsa.KeyPath())
	assert.Equal(t, pemPrivateKey, content)
}

func TestKeyPath(t *testing.T) {
	repository := BuildTestRepository()
	rsa := repository.Rsa

	assert.Equal(t, "/tmp/deploy_docus/keys/lyonrb_deploy_docus", rsa.KeyPath())

	rsa.Repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "/tmp/deploy_docus/keys/github_hubot", rsa.KeyPath())
}

func TestPublicKey(t *testing.T) {
	repository := BuildTestRepository()
	rsa := repository.Rsa

	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAQQCymQ9JxH36jNQArmpNG4o7ahNkKyPyiwA7+5d5Ct6aTMgriyqBdH3ewItiluU6CMMxaH7yXEv0k2uhwOYEHp0V\n", rsa.PublicKey())
}
