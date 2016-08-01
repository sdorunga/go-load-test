package main

import(
  "net/http"
  "fmt"
  "time"
  "strconv"
)

/////////
func median(numbers []int) int {
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
  min := numbers[0]
  for _, number := range numbers {
    if min > number {
      min = number
    }
  }
  return min
}

func max(numbers []int) int {
  max := numbers[0]
  for _, number := range numbers {
    if max < number {
      max = number
    }
  }
  return max
}
/////////

func run_requests(concurrency int, requests []Request) {
  jobs := make(chan []Request, 100)
  results := make(chan int, 100)

  setup_workers(concurrency, jobs, results)
  queue_jobs(concurrency, jobs, results, requests)
}

func wait_for_results(count int, request_count int, results <-chan int) {
  times := make([]int, 0)
  for i:= 0; i < count * request_count; i++ {
    times = append(times, <-results)
  }
  fmt.Println("Number of Requests: ", len(times))
  fmt.Println("Average: ", sum(times)/len(times))
  fmt.Println("Median: ", median(times))
  fmt.Println("Min: ", min(times))
  fmt.Println("Max: ", max(times))
}

func queue_jobs(count int, jobs chan<- []Request, results <-chan int, requests []Request) {
  for i := 0; i < count; i++ {
    jobs <- requests
  }
  close(jobs)
  wait_for_results(count, len(requests), results)
}

func setup_workers(count int, jobs <-chan []Request, results chan<- int) {
  for i := 0; i < count; i++ {
    go worker(i, jobs, results)
  }
}

func worker(id int, jobs <-chan []Request, results chan<- int) {
  for requests := range jobs {
    for _, request := range requests {
      time := time_request(request, id)
      results <- time
    }
  }
}

func time_request(request Request, worker_id int) int {
  start := now_in_millis()
  _, err := http.Get(request.Url)
  if err != nil {
    fmt.Println("Error making a " + request.Verb + " request to " + request.Url)
  }
  request_time := now_in_millis() - start
  fmt.Println("Worker " + strconv.Itoa(worker_id) + " processed " + request.Verb + " for " + request.Url + " in: " , request_time, "ms")
  return int(request_time)
}

func now_in_millis() int64 {
  return time.Now().UnixNano() / int64(time.Millisecond)
}
