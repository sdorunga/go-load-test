package main

import (
	"net/http"
	"testing"
)

type TestClient struct {
	requests      []string
	requestsCount int
}

func (this *TestClient) Do(request Request) (*http.Response, error) {
	this.requests = append(this.requests, request.URL)
	this.requestsCount++
	return &http.Response{}, nil
}

func TestHttpRequestDelegatedToClient(t *testing.T) {
	client := TestClient{}
	req := Request{"GET", "http://google.com", &client}
	req.Perform()
	if client.requestsCount != 1 {
		t.Errorf("Expected client to make 1 request, made %d", client.requestsCount)
	}

	if client.requests[0] != "http://google.com" {
		t.Errorf("Expected client to make 1 request, made %d", client.requestsCount)
	}
}
