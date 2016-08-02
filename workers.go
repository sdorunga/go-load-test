package main

import (
	"fmt"
	"strconv"
	"time"
)

/////////
func average(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return sum(numbers) / len(numbers)
}

func median(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func sum(numbers []int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return total
}

func min(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
	}
	return min
}

func max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}
	return max
}

/////////

func runRequests(concurrency int, requests []Request) {
	jobs := make(chan []Request, 100)
	results := make(chan int, 100)

	setupWorkers(concurrency, jobs, results)
	queueJobs(concurrency, jobs, results, requests)
}

func waitForResults(count int, requestCount int, results <-chan int) {
	times := make([]int, 0)
	for i := 0; i < count*requestCount; i++ {
		times = append(times, <-results)
	}
	fmt.Println("Number of Requests: ", len(times))
	fmt.Println("Average: ", average(times))
	fmt.Println("Median: ", median(times))
	fmt.Println("Min: ", min(times))
	fmt.Println("Max: ", max(times))
}

func queueJobs(count int, jobs chan<- []Request, results <-chan int, requests []Request) {
	for i := 0; i < count; i++ {
		jobs <- requests
	}
	close(jobs)
	waitForResults(count, len(requests), results)
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
	_, err := request.Perform()
	if err != nil {
		fmt.Println("Error making a " + request.Verb + " request to " + request.URL)
	}
	requestTime := nowInMillis() - start
	fmt.Println("Worker "+strconv.Itoa(workerID)+" processed "+request.Verb+" for "+request.URL+" in: ", requestTime, "ms")
	return int(requestTime)
}

func nowInMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
