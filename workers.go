package main

func runRequests(concurrency int, requests []TimeableTask) []TimeableTask {
	jobs := make(chan []TimeableTask, 100)
	results := make(chan TimeableTask, 100)

	setupWorkers(concurrency, jobs, results)
	queueJobs(concurrency, jobs, results, requests)
	return gatherTimes(concurrency, len(requests), results)
}

func gatherTimes(count int, requestCount int, results <-chan TimeableTask) []TimeableTask {
	requests := make([]TimeableTask, 0)
	for i := 0; i < count*requestCount; i++ {
		requests = append(requests, <-results)
	}
	return requests
}

func queueJobs(count int, jobs chan<- []TimeableTask, results <-chan TimeableTask, requests []TimeableTask) {
	for i := 0; i < count; i++ {
		jobs <- requests
	}
	close(jobs)
}

func setupWorkers(count int, jobs <-chan []TimeableTask, results chan<- TimeableTask) {
	for i := 0; i < count; i++ {
		go worker(i, jobs, results)
	}
}

func worker(id int, jobs <-chan []TimeableTask, results chan<- TimeableTask) {
	for requests := range jobs {
		for _, request := range requests {
			request.Perform()
			results <- request
		}
	}
}
