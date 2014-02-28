package deploy_docus

import (
	"bytes"
	"github.com/bmizerany/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessfulGetHome(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusMovedPermanently, response.Code)
	assert.Equal(t, "/login", response.Header().Get("Location"))
}

func TestSuccessfulGetHomeSignedIn(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("GET", "/", nil)
	LoginTest(t, server, request)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusMovedPermanently, response.Code)
	assert.Equal(t, "/repositories", response.Header().Get("Location"))
}

func TestSuccessfulGetLogin(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusFound, response.Code)
}
