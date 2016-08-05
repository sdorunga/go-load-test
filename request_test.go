package main

import (
	"net/http"
	"testing"
)

type TestClient struct {
	requests      []string
	requestsCount int
}

type TestTimer struct {
	startTime int64
	endTime   int64
}

func (this *TestTimer) Start() {
	this.startTime = 100
}

func (this *TestTimer) Stop() {
	this.endTime = 200
}

func (this *TestTimer) Duration() int {
	return int(this.endTime - this.startTime)
}

func (this *TestClient) Do(request Request) (http.Response, error) {
	this.requests = append(this.requests, request.URL)
	this.requestsCount++
	return http.Response{}, nil
}

func TestDurationDelegatedToTimer(t *testing.T) {
	client := TestClient{}
	timer := TestTimer{}
	req := Request{"GET", "http://google.com", &client, false, &timer}
	req.Perform()
	if req.Duration() != 100 {
		t.Errorf("Expected request to take 100ms, took %d", req.Duration())
	}
}

func TestHttpRequestDelegatedToClient(t *testing.T) {
	client := TestClient{}
	timer := TestTimer{}
	req := Request{"GET", "http://google.com", &client, false, &timer}
	req.Perform()
	if client.requestsCount != 1 {
		t.Errorf("Expected client to make 1 request, made %d", client.requestsCount)
	}

	if client.requests[0] != "http://google.com" {
		t.Errorf("Expected client to make 1 request, made %d", client.requestsCount)
	}
}
