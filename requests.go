package main

import (
	"fmt"
	"io/ioutil"
	"strings"
  "net/http"
)

type Request struct {
	Verb string
	URL  string
  client http.Client
}

func (this *Request) Perform() (resp *http.Response, err error) {
  resp, err = this.client.Get(this.URL)
  return
}

func requests(fileName string, httpClient http.Client) []Request {
	requestLines := fileLines(fileName)
	return buildRequests(requestLines, httpClient)
}

func buildRequests(lines []string, httpClient http.Client) []Request {
	requests := make([]Request, 0)
	for _, line := range lines {
		vars := strings.Fields(line)
		if len(vars) == 0 {
			break
		}

		requests = append(requests, Request{vars[0], vars[1], httpClient})
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
