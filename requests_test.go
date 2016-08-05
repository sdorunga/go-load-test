package main

import (
	"testing"
)

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

func TestHTTPRequestsPerformCallsTheClient(t *testing.T) {
	client := TestClient{}
	reqs := requests("test_support/requests", &client, false)
	reqs[0].Perform()
	reqs[1].Perform()
	if client.requestsCount != 2 {
		t.Errorf("Expected 2 requests, got %d", client.requestsCount)
	}
}

func TestHTTPRequestsPerformCallsAreTimed(t *testing.T) {
	client := TestClient{}
	timer := TestTimer{}
	reqs := requests("test_support/requests", &client, false)
	req := reqs[0]
	req.timer = &timer
	req.Perform()
	if req.Duration() != 100 {
		t.Errorf("Expected 100ms duration, got %d", req.Duration())
	}
}
