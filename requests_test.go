package main

import "testing"

func TestCorrectNumberOfRequests(t *testing.T) {
	if countRequests := len(requests("test_support/requests")); countRequests != 2 {
		t.Errorf("Wrong number of requests built: expected 2, got %d", countRequests)
	}
}

func TestCorrectURLsAreExtracted(t *testing.T) {
	reqs := requests("test_support/requests")
	if reqs[0].URL != "https://google.com" || reqs[1].URL != "https://formly.com?q=hello" {
		t.Errorf("Expected https://google.com and https://formly.com?q=hello urls. Got: %s and %s", reqs[0].URL, reqs[1].URL)
	}
}

func TestCorrectVerbsExtracted(t *testing.T) {
	reqs := requests("test_support/requests")
	if reqs[0].Verb != "GET" || reqs[1].Verb != "POST" {
		t.Errorf("Expected GET and POST verbs. Got: %s and %s", reqs[0].Verb, reqs[1].Verb)
	}
}
