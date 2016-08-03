package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	Verb    string
	URL     string
	client  Client
	verbose bool
}

type Client interface {
	Do(request Request)
}

type httpClient struct {
	client http.Client
}

func (this *httpClient) Do(request Request) {
  req, err := http.NewRequest(request.Verb, request.URL, nil)
  if err != nil {
		fmt.Println("Error making a " + request.Verb + " request to " + request.URL)
  }
	this.client.Do(req)
}

func (this *Request) Perform() (resp *http.Response, err error) {
	start := nowInMillis()
	this.client.Do(*this)
	if err != nil {
		fmt.Println("Error making a " + this.Verb + " request to " + this.URL)
	}
	requestTime := nowInMillis() - start
	if this.verbose {
		fmt.Println("Processed "+this.Verb+" for "+this.URL+" in: ", requestTime, "ms")
	}
	return
}

func requests(fileName string, httpClient Client, verbose bool) []Request {
	requestLines := fileLines(fileName)
	return buildRequests(requestLines, httpClient, verbose)
}

func buildRequests(lines []string, httpClient Client, verbose bool) []Request {
	requests := make([]Request, 0)
	for _, line := range lines {
		vars := strings.Fields(line)
		if len(vars) == 0 {
			break
		}

		requests = append(requests, Request{vars[0], vars[1], httpClient, verbose})
	}
	return requests
}

func fileLines(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading request file\nMake sure file exists and is in the format:\nVERB https://example.com")
	}
	return strings.Split(string(data), "\n")
}
