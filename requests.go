package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Request struct {
	Verb string
	URL  string
}

func requests(fileName string) []Request {
	requestLines := fileLines(fileName)
	return buildRequests(requestLines)
}

func buildRequests(lines []string) []Request {
	requests := make([]Request, 0)
	for _, line := range lines {
		vars := strings.Fields(string(line))
		if len(vars) == 0 {
			break
		}

		requests = append(requests, Request{vars[0], vars[1]})
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
