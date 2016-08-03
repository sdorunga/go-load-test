package main

import (
	"flag"
  "net/http"
)

func main() {
	concurrency := flag.Int("c", 1, "Number of concurrent workers to schedule")
	requestsFile := flag.String("f", "requests", "Name of file to use as the requests list")
	flag.Parse()

  client := httpClient{http.Client{}}
	runRequests(*concurrency, requests(*requestsFile, &client))
}
