package deploy_docus

import (
	"github.com/bmizerany/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestSuccessfulEncrypt(t *testing.T) {
	key, _ := BuildPrivateKey([]byte(pemPrivateKey))
	rsa := &Rsa{Private: key}

	encrypted, err := rsa.Encrypt([]byte("hello"))

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, encrypted)
}

func TestSuccessfulDecrypt(t *testing.T) {
	key, _ := BuildPrivateKey([]byte(pemPrivateKey))
	rsa := &Rsa{Private: key}

	encrypted, err := rsa.Encrypt([]byte("hello world"))
	decrypted, err := rsa.Decrypt(encrypted)

	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("hello world"), decrypted)
}

func TestNewRsa(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git", PKey: pemPrivateKey}
	rsa := NewRsa(repository)

	assert.Equal(t, repository, rsa.Repository)
	assert.NotEqual(t, nil, rsa.Private)
}

func TestCreateKeyFile(t *testing.T) {
	os.RemoveAll("/tmp/deploy_docus/keys")

	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git", PKey: pemPrivateKey}
	rsa := NewRsa(repository)

	_, err := os.Stat(rsa.KeyPath())
	assert.Equal(t, true, os.IsNotExist(err))

	rsa.WriteKey()

	_, err = os.Stat(rsa.KeyPath())
	assert.Equal(t, false, os.IsNotExist(err))

	content, err := ioutil.ReadFile(rsa.KeyPath())
	assert.Equal(t, pemPrivateKey, string(content))
}

func TestKeyPath(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	rsa := &Rsa{Repository: repository}
	assert.Equal(t, "/tmp/deploy_docus/keys/dmathieu_deploy_docus", rsa.KeyPath())

	rsa.Repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "/tmp/deploy_docus/keys/github_hubot", rsa.KeyPath())
}
