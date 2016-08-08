package main

import (
	"testing"
)

type TestRequest struct {
  isPerformed bool
  performCount int
}

func (this *TestRequest) Perform() {
  this.isPerformed = true
  this.performCount++
}

func (this *TestRequest) Duration() int {
  return 1
}

func TestWorkerCallsPerformOnRequestsOnceWhenNotConcurrent(t *testing.T) {
  request := &TestRequest{}
  requests := []TimeableTask{request}
  results := runRequests(1, requests)
  performedRequest := results[0].(*TestRequest)
	if performedRequest.isPerformed != true {
		t.Errorf("Expected request to be performed, got %b", performedRequest.isPerformed)
	}

	if performedRequest.performCount != 1 {
		t.Errorf("Expected request to be performed once, got %d performs", performedRequest.performCount)
	}
}

func TestWorkerCallsPerformOnRequestsTwiceWithConcurrency(t *testing.T) {
  request := &TestRequest{}
  requests := []TimeableTask{request}
  results := runRequests(2, requests)
  performedRequest := results[0].(*TestRequest)
	if performedRequest.isPerformed != true {
		t.Errorf("Expected request to be performed, got %b", performedRequest.isPerformed)
	}

	if performedRequest.performCount != 2 {
		t.Errorf("Expected request to be performed twice, got %d performs", performedRequest.performCount)
	}
}
