package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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
