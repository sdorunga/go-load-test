package main

import (
	"flag"
	"net/http"
)

func main() {
	concurrency := flag.Int("c", 1, "Number of concurrent workers to schedule")
	requestsFile := flag.String("f", "requests", "Name of file to use as the requests list")
	hammerMode := flag.Bool("h", false, "Whether the request sequence should loop until cancelled from the command line")
	verbose := flag.Bool("v", false, "Whether to print each request as it happens")
	flag.Parse()

	client := httpClient{http.Client{}}

	time := nowInMillis()
	results := runRequests(*concurrency, requests(*requestsFile, &client), *verbose, *hammerMode)
	totalTime := nowInMillis() - time

	var times []int
	for _, result := range results {
		times = append(times, result.duration)
	}
	stats := StatsPrinter{times, totalTime}
	stats.Print()
}
