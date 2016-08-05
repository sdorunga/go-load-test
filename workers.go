package main

func runRequests(concurrency int, requests []Request) []Request {
	jobs := make(chan []Request, 100)
	results := make(chan Request, 100)

	setupWorkers(concurrency, jobs, results)
	queueJobs(concurrency, jobs, results, requests)
	return gatherTimes(concurrency, len(requests), results)
}

func gatherTimes(count int, requestCount int, results <-chan Request) []Request {
	requests := make([]Request, 0)
	for i := 0; i < count*requestCount; i++ {
		requests = append(requests, <-results)
	}
	return requests
}

func queueJobs(count int, jobs chan<- []Request, results <-chan Request, requests []Request) {
	for i := 0; i < count; i++ {
		jobs <- requests
	}
	close(jobs)
}

func setupWorkers(count int, jobs <-chan []Request, results chan<- Request) {
	for i := 0; i < count; i++ {
		go worker(i, jobs, results)
	}
}

func worker(id int, jobs <-chan []Request, results chan<- Request) {
	for requests := range jobs {
		for _, request := range requests {
      request.Perform()
			results <- request
		}
	}
}
