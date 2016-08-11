package main

import (
	"fmt"
	"time"
  "os"
  "os/signal"
)

type Result struct {
	task     Task
	duration int
}

type Task interface {
	Perform()
}

func runRequests(concurrency int, requests []Task, verbose bool, hammerMode bool) []Result {
	jobs := make(chan []Task, 1)
	results := make(chan Result, 10)
  unGathered := make(chan bool, 10)
	finalResults := make(chan []Result, 1)

	setupWorkers(concurrency, jobs, results, verbose)
  if hammerMode {
    go queueInfiniteJobs(jobs, requests, unGathered)
  } else {
    go queueJobs(concurrency, jobs, requests, unGathered)
  }
	go gatherResults(results, finalResults, unGathered)
  return <-finalResults
}

func gatherResults(results <-chan Result, finalResults chan<- []Result, unGathered <-chan bool) {
	requests := make([]Result, 0)
  for _ = range unGathered {
    requests = append(requests, <-results)
  }
	finalResults <- requests
}

func queueInfiniteJobs(jobs chan<- []Task, requests []Task, unGathered chan<- bool) {
  interrupt := make(chan os.Signal, 1)
  signal.Notify(interrupt, os.Interrupt)
  loop:
	for {
    select {
    case <- interrupt:
      break loop
    default:
      jobs <- requests
      for _, _ = range requests {
        unGathered <- true
      }
    }
	}
	close(jobs)
  close(unGathered)
}

func queueJobs(count int, jobs chan<- []Task, requests []Task, unGathered chan<- bool) {
	for i := 0; i < count; i++ {
		jobs <- requests
    for _, _ = range requests {
      unGathered <- true
    }
	}
	close(jobs)
  close(unGathered)
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
