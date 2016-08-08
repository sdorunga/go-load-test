package main

import (
	"testing"
)

func TestCorrectNumberOfequests(t *testing.T) {
	if countRequests := len(requests("test_support/requests", &TestClient{})); countRequests != 2 {
		t.Errorf("Wrong number of requests built: expected 2, got %d", countRequests)
	}
}

func TestCorrectURLsAreExtracted(t *testing.T) {
	reqs := requests("test_support/requests", &TestClient{})
	if reqs[0].(*Request).URL != "https://google.com" || reqs[1].(*Request).URL != "https://formly.com?q=hello" {
		t.Errorf("Expected https://google.com and https://formly.com?q=hello urls. Got: %s and %s", reqs[0].(*Request).URL, reqs[1].(*Request).URL)
	}
}

func TestCorrectVerbsExtracted(t *testing.T) {
	reqs := requests("test_support/requests", &TestClient{})
	if reqs[0].(*Request).Verb != "GET" || reqs[1].(*Request).Verb != "POST" {
		t.Errorf("Expected GET and POST verbs. Got: %s and %s", reqs[0].(*Request).Verb, reqs[1].(*Request).Verb)
	}
}
