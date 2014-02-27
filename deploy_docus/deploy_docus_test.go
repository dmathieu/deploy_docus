package deploy_docus

import (
	"github.com/bmizerany/assert"
	"github.com/martini-contrib/sessions"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	pemPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBALKZD0nEffqM1ACuak0bijtqE2QrI/KLADv7l3kK3ppMyCuLKoF0fd7Ai2KW5ToIwzFofvJcS/STa6HA5gQenRUCAwEAAQJBAIq9amn00aS0h/CrjXqu/ThglAXJmZhOMPVn4eiu7/ROixi9sex436MaVeMqSNf7Ex9a8fRNfWss7Sqd9eWuRTUCIQDasvGASLqmjeffBNLTXV2A5g4t+kLVCpsEIZAycV5GswIhANEPLmax0ME/EO+ZJ79TJKN5yiGBRsv5yvx5UiHxajEXAiAhAol5N4EUyq6I9w1rYdhPMGpLfk7AIU2snfRJ6Nq2CQIgFrPsWRCkV+gOYcajD17rEqmuLrdIRexpg8N1DOSXoJ8CIGlStAboUGBxTDq3ZroNism3DaMIbKPyYrAqhKov1h5V
-----END RSA PRIVATE KEY-----
`)

	repositoryOrigin      = "git@github.com:lyonrb/deploy_docus.git"
	repositoryDestination = "git@heroku.com:deploy_docus.git"
)

func BuildTestRepository() *Repository {
	repository := &Repository{Origin: repositoryOrigin, Destination: repositoryDestination}
	repository.Rsa = BuildRsa(repository, pemPrivateKey)

	return repository
}

func RemoveAllRepositories() error {
	db, err := GetDbConnection()
	if err != nil {
		return err
	}
	_, err = db.Exec(`TRUNCATE TABLE repositories;`)
	if err != nil {
		return err
	}
	return nil
}

func LoginTest(t *testing.T, server *Server, request *http.Request) {
	server.Get("/test_login", func(s sessions.Session) string {
		s.Set("oauth2_token", []byte("{\"AccessToken\":\"\",\"RefreshToken\":\"\",\"Expiry\":\"2025-01-01T00:00:00Z\",\"Extra\":null}"))
		return "OK"
	})

	login, err := http.NewRequest("GET", "/test_login", nil)
	if err != nil {
		panic(err)
	}

	loginResponse := httptest.NewRecorder()
	server.ServeHTTP(loginResponse, login)
	assert.Equal(t, http.StatusOK, loginResponse.Code)
	request.Header.Set("Cookie", loginResponse.Header().Get("Set-Cookie"))
}
