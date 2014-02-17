package deploy_docus

import (
	"github.com/bmizerany/assert"
	"io/ioutil"
	"os"
	"testing"
)

var pemPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBALKZD0nEffqM1ACuak0bijtqE2QrI/KLADv7l3kK3ppMyCuLKoF0
fd7Ai2KW5ToIwzFofvJcS/STa6HA5gQenRUCAwEAAQJBAIq9amn00aS0h/CrjXqu
/ThglAXJmZhOMPVn4eiu7/ROixi9sex436MaVeMqSNf7Ex9a8fRNfWss7Sqd9eWu
RTUCIQDasvGASLqmjeffBNLTXV2A5g4t+kLVCpsEIZAycV5GswIhANEPLmax0ME/
EO+ZJ79TJKN5yiGBRsv5yvx5UiHxajEXAiAhAol5N4EUyq6I9w1rYdhPMGpLfk7A
IU2snfRJ6Nq2CQIgFrPsWRCkV+gOYcajD17rEqmuLrdIRexpg8N1DOSXoJ8CIGlS
tAboUGBxTDq3ZroNism3DaMIbKPyYrAqhKov1h5V
-----END RSA PRIVATE KEY-----
`

func TestName(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	assert.Equal(t, "dmathieu_deploy_docus", repository.Name())

	repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "github_hubot", repository.Name())
}

func TestLocalPath(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	assert.Equal(t, "/tmp/dmathieu_deploy_docus", repository.LocalPath())

	repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "/tmp/github_hubot", repository.LocalPath())
}

func TestCreatePKeyFile(t *testing.T) {
	os.RemoveAll("/tmp/deploy_docus/keys")
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git", PKey: "private_key_content"}

	_, err := os.Stat(repository.PKeyPath())
	assert.Equal(t, true, os.IsNotExist(err))

	repository.CreatePKey()

	_, err = os.Stat(repository.PKeyPath())
	assert.Equal(t, false, os.IsNotExist(err))

	content, err := ioutil.ReadFile(repository.PKeyPath())
	assert.Equal(t, "private_key_content", string(content))
}

func TestPKeyPath(t *testing.T) {
	repository := &Repository{Origin: "git@github.com:dmathieu/deploy_docus.git"}
	assert.Equal(t, "/tmp/deploy_docus/keys/dmathieu_deploy_docus", repository.PKeyPath())

	repository.Origin = "git@github.com:github/hubot.git"
	assert.Equal(t, "/tmp/deploy_docus/keys/github_hubot", repository.PKeyPath())
}

func TestSuccessfulEncryptToken(t *testing.T) {
	repository := &Repository{PKey: pemPrivateKey, Origin: "git@github.com:dmathieu/deploy_docus.git"}

	token, err := repository.Token()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, token)
}

func TestSuccessfulDecryptToken(t *testing.T) {
	repository := &Repository{PKey: pemPrivateKey, Origin: "git@github.com:dmathieu/deploy_docus.git"}
	token, _ := repository.Token()

	value, err := repository.DecryptToken(token)

	assert.Equal(t, nil, err)
	assert.Equal(t, repository.Name(), string(value))
}

func TestSuccessfulFindRepository(t *testing.T) {
	os.Setenv("REPOSITORY_ORIGIN", "git@github.com:dmathieu/deploy_docus.git")
	os.Setenv("REPOSITORY_DESTINATION", "git@heroku.com:deploy_docus.git")
	os.Setenv("REPOSITORY_PKEY", "private_repository_ssh_key_content")

	var repository *Repository
	repository = FindRepository()

	assert.Equal(t, "git@github.com:dmathieu/deploy_docus.git", repository.Origin)
	assert.Equal(t, "git@heroku.com:deploy_docus.git", repository.Destination)
	assert.Equal(t, "private_repository_ssh_key_content", repository.PKey)
}
