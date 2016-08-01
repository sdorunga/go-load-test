package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Request struct {
	Verb string
	Url  string
}

func requests(file_name string) []Request {
	request_lines := file_lines(file_name)
	return build_requests(request_lines)
}

func build_requests(lines []string) []Request {
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

func file_lines(file_name string) []string {
	data, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println("Error reading request file\nMake sure file exists and is in the format:\nVERB https://example.com")
	}
	return strings.Split(string(data), "\n")
}
