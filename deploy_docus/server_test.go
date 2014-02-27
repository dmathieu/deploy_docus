package deploy_docus

import (
	"bytes"
	"fmt"
	"github.com/bmizerany/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessfulGetHome(t *testing.T) {
	server := NewServer(80, nil)

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `You might be "a doctor". I am "the doctor".`, fmt.Sprintf("%s", response.Body))
}

func TestSuccessfulGetLogin(t *testing.T) {
	server := NewServer(80, nil)

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusFound, response.Code)
}
