package main

import (
	"testing"
)

type TestClient struct {
	requests      []string
	requestsCount int
}

func (this *TestClient) Do(request Request) {
	this.requests = append(this.requests, request.URL)
	this.requestsCount++
}

func TestCorrectNumberOfequests(t *testing.T) {
	if countRequests := len(requests("test_support/requests", &TestClient{}, false)); countRequests != 2 {
		t.Errorf("Wrong number of requests built: expected 2, got %d", countRequests)
	}
}

func TestCorrectURLsAreExtracted(t *testing.T) {
	reqs := requests("test_support/requests", &TestClient{}, false)
	if reqs[0].URL != "https://google.com" || reqs[1].URL != "https://formly.com?q=hello" {
		t.Errorf("Expected https://google.com and https://formly.com?q=hello urls. Got: %s and %s", reqs[0].URL, reqs[1].URL)
	}
}

func TestCorrectVerbsExtracted(t *testing.T) {
	reqs := requests("test_support/requests", &TestClient{}, false)
	if reqs[0].Verb != "GET" || reqs[1].Verb != "POST" {
		t.Errorf("Expected GET and POST verbs. Got: %s and %s", reqs[0].Verb, reqs[1].Verb)
	}
}

func TestHTTPRequests(t *testing.T) {
	client := TestClient{}
	reqs := requests("test_support/requests", &client, false)
	reqs[0].Perform()
	reqs[1].Perform()
	if client.requestsCount != 2 {
		t.Errorf("Expected 2 requests, got %d", client.requestsCount)
	}
}
