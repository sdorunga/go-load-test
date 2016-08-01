package main

import (
	"flag"
)

func main() {
	concurrency := flag.Int("c", 1, "Number of concurrent workers to schedule")
	requests_file := flag.String("f", "requests", "Name of file to use as the requests list")
	flag.Parse()

	run_requests(*concurrency, requests(*requests_file))
}
