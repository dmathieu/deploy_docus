package server

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

func TestSuccessfulGetHome(t *testing.T) {
	server := NewServer(80, nil)

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	request, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		panic(err)
	}
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, `You might be "a doctor". I am "the doctor".`, fmt.Sprintf("%s", response.Body))
}

func TestSuccessfulDeploy(t *testing.T) {
	channel := make(chan Message)
	server := NewServer(80, channel)

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
	request, err := http.NewRequest("POST", "http://localhost:3000/deploy", content)
	assert.Equal(t, nil, err)

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	retrieveMessage := func(channel chan Message) {
		message := <-channel
		assert.Equal(t, 1, message.Id)
		assert.Equal(t, "a84d88e7554fc1fa21bcbc4efae3c782a70d2b9d", message.Sha)
	}

	go retrieveMessage(channel)
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)
}
