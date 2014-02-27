package deploy_docus

import (
	"bytes"
	"fmt"
	"github.com/bmizerany/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestUnauthorizedGetRepositoriesList(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/repositories", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusFound, response.Code)
}

func TestSuccessfulGetRepositoriesList(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	request, err := http.NewRequest("GET", "/repositories", nil)
	if err != nil {
		panic(err)
	}
	LoginTest(t, server, request)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUnauthorizedGetRepositoryNew(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/repositories/new", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusFound, response.Code)
}

func TestSuccessfulGetRepositoryNew(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	request, err := http.NewRequest("GET", "/repositories/new", nil)
	if err != nil {
		panic(err)
	}
	LoginTest(t, server, request)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUnauthorizedGetRepository(t *testing.T) {
	RemoveAllRepositories()
	repository := BuildTestRepository()
	repository.Save()
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	url := fmt.Sprintf("/repositories/%d", repository.Id)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusFound, response.Code)
}

func TestSuccessfulGetRepository(t *testing.T) {
	RemoveAllRepositories()
	repository := BuildTestRepository()
	repository.Save()
	server := NewServer(80, nil, ServerPath())

	url := fmt.Sprintf("/repositories/%d", repository.Id)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	LoginTest(t, server, request)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUnauthorizedPostCreateRepository(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/repositories", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusFound, response.Code)
}

func TestSuccessfulPostCreateRepository(t *testing.T) {
	server := NewServer(80, nil, ServerPath())

	payload := url.Values{
		"repository[origin]":      {repositoryOrigin},
		"repository[destination]": {repositoryDestination},
	}
	content := bytes.NewBufferString(payload.Encode())
	request, err := http.NewRequest("POST", "/repositories", content)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	LoginTest(t, server, request)
	response := httptest.NewRecorder()

	count, _ := RepositoriesCount()
	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusMovedPermanently, response.Code)
	newCount, _ := RepositoriesCount()
	assert.Equal(t, count+1, newCount)
}
