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

func TestSuccessfulDeploy(t *testing.T) {
	RemoveAllRepositories()
	tmp := BuildTestRepository()
	tmp.Save()

	channel := make(chan Message)
	server := NewServer(80, channel, ServerPath())

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	payload := url.Values{
		"id":             {"1"},
		"sha":            {"a84d88e7554fc1fa21bcbc4efae3c782a70d2b9d"},
		"url":            {"https://api.github.com/repos/octocat/example/deployments/1"},
		"creator[login]": {"octocat"},
		"creator[id]":    {"1"},
		"creator[type]":  {"User"},
		"payload":        {"{\"environment\":\"production\"}"},
		"created_at":     {"2012-07-20T01:19:13Z"},
		"updated_at":     {"2012-07-20T01:19:13Z"},
		"description":    {"Deploy request from hubot"},
		"statuses_url":   {"https://api.github.com/repos/octocat/example/deployments/1/statuses"},
	}
	content := bytes.NewBufferString(payload.Encode())
	url := fmt.Sprintf("/deploy/%d?token=%s", tmp.Id, tmp.Token())
	request, err := http.NewRequest("POST", url, content)
	assert.Equal(t, nil, err)

	request.Header.Add("X-GitHub-Event", "deployment")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	retrieveMessage := func(channel chan Message) {
		message := <-channel
		assert.Equal(t, 1, message.Id)
		assert.Equal(t, "a84d88e7554fc1fa21bcbc4efae3c782a70d2b9d", message.Sha)

		assert.NotEqual(t, nil, message.Repository)
		assert.Equal(t, tmp.Id, message.Repository.Id)
	}

	go retrieveMessage(channel)
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestDeployMissingRepository(t *testing.T) {
	RemoveAllRepositories()
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("POST", "/deploy/42", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("X-GitHub-Event", "deployment")
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestDeployInvalidToken(t *testing.T) {
	RemoveAllRepositories()
	tmp := BuildTestRepository()
	tmp.Save()
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	url := fmt.Sprintf("/deploy/%d", tmp.Id)
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("X-GitHub-Event", "deployment")
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestDeployOtherEvent(t *testing.T) {
	RemoveAllRepositories()
	server := NewServer(80, nil, ServerPath())

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("POST", "/deploy/42", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("X-GitHub-Event", "push")
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}
