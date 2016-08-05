package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
  "time"
)

type Request struct {
	Verb    string
	URL     string
	client  Client
	verbose bool
  timer Timer
}

type Timer interface {
  Start()
  Stop()
  Duration() int
}

type MilliTimer struct {
  startTime int64
  endTime int64
}

type Client interface {
	Do(request Request) (http.Response, error)
}

type httpClient struct {
	client http.Client
}

func (this *httpClient) Do(request Request) (res http.Response, err error) {
  req, err := http.NewRequest(request.Verb, request.URL, nil)
  if err != nil {
		fmt.Println("Error making a " + request.Verb + " request to " + request.URL)
  }
	this.client.Do(req)
  return
}

func (this *Request) Perform() {
  this.timer.Start()
  _, err := this.client.Do(*this)
	if err != nil {
		fmt.Println("Error making a " + this.Verb + " request to " + this.URL)
	}
  this.timer.Stop()
	if this.verbose {
		fmt.Println("Processed "+this.Verb+" for "+this.URL+" in: ", this.Duration(), "ms")
	}
	return
}

func (this *Request) Duration() int {
  return this.timer.Duration()
}

func (this *MilliTimer) Duration() int {
  return int(this.endTime - this.startTime)
}

func (this *MilliTimer) Start() {
  this.startTime = this.nowInMillis()
}

func (this *MilliTimer) Stop() {
  this.endTime = this.nowInMillis()
}

func (this *MilliTimer) nowInMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
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

		requests = append(requests, Request{vars[0], vars[1], httpClient, verbose, &MilliTimer{}})
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
