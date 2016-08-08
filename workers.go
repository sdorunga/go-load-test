package main

import (
	"fmt"
	"time"
)

type Result struct {
	task     Task
	duration int
}

type Task interface {
	Perform()
}

func runRequests(concurrency int, requests []Task, verbose bool) []Result {
	jobs := make(chan []Task, 100)
	results := make(chan Result, 100)

	setupWorkers(concurrency, jobs, results, verbose)
	queueJobs(concurrency, jobs, results, requests)
	return gatherResults(concurrency, len(requests), results)
}

func gatherResults(count int, requestCount int, results <-chan Result) []Result {
	requests := make([]Result, 0)
	for i := 0; i < count*requestCount; i++ {
		requests = append(requests, <-results)
	}
	return requests
}

func queueJobs(count int, jobs chan<- []Task, results <-chan Result, requests []Task) {
	for i := 0; i < count; i++ {
		jobs <- requests
	}
	close(jobs)
}

func setupWorkers(count int, jobs <-chan []Task, results chan<- Result, verbose bool) {
	for i := 0; i < count; i++ {
		go worker(i, jobs, results, verbose)
	}
}

func worker(id int, jobs <-chan []Task, results chan<- Result, verbose bool) {
	for requests := range jobs {
		for _, request := range requests {
			time := nowInMillis()
			request.Perform()
			duration := nowInMillis() - time
			if verbose {
				fmt.Println("Processed ", request, " in: ", duration, "ms")
			}
			results <- Result{request, duration}
		}
	}
}

func nowInMillis() int {
	return int(time.Now().UnixNano() / int64(time.Millisecond))
}
