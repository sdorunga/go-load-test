package main

import (
	"time"
)

func runRequests(concurrency int, requests []Request) {
	jobs := make(chan []Request, 100)
	results := make(chan int, 100)

	setupWorkers(concurrency, jobs, results)
	queueJobs(concurrency, jobs, results, requests)
	times := gatherTimes(concurrency, len(requests), results)

	stats := StatsPrinter{times}
	stats.Print()
}

func gatherTimes(count int, requestCount int, results <-chan int) []int {
	times := make([]int, 0)
	for i := 0; i < count*requestCount; i++ {
		times = append(times, <-results)
	}
	return times
}

func queueJobs(count int, jobs chan<- []Request, results <-chan int, requests []Request) {
	for i := 0; i < count; i++ {
		jobs <- requests
	}
	close(jobs)
}

func setupWorkers(count int, jobs <-chan []Request, results chan<- int) {
	for i := 0; i < count; i++ {
		go worker(i, jobs, results)
	}
}

func worker(id int, jobs <-chan []Request, results chan<- int) {
	for requests := range jobs {
		for _, request := range requests {
			time := timeRequest(request, id)
			results <- time
		}
	}
}

func timeRequest(request Request, workerID int) int {
	start := nowInMillis()
	request.Perform()
	requestTime := nowInMillis() - start
	return int(requestTime)
}

func nowInMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
