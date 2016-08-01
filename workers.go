package main

import(
  "net/http"
  "fmt"
  "time"
  "strconv"
)

func run_requests(concurrency int, requests []Request) {
  jobs := make(chan []Request, 100)
  results := make(chan bool, 100)

  setup_workers(concurrency, jobs, results)
  queue_jobs(concurrency, jobs, results, requests)
}

func wait_for_results(count int, request_count int, results <-chan bool) {
  for i:= 0; i < count * request_count; i++ {
    <-results
  }
}

func queue_jobs(count int, jobs chan<- []Request, results <-chan bool, requests []Request) {
  for i := 0; i < count; i++ {
    jobs <- requests
  }
  close(jobs)
  wait_for_results(count, len(requests), results)
}

func setup_workers(count int, jobs <-chan []Request, results chan<- bool) {
  for i := 0; i < count; i++ {
    go worker(i, jobs, results)
  }
}

func worker(id int, jobs <-chan []Request, results chan<- bool) {
  for requests := range jobs {
    for _, request := range requests {
      start := now_in_millis()
      _, err := http.Get(request.Url)
      if err != nil {
        fmt.Println("Error making a " + request.Verb + " request to " + request.Url)
      }
      fmt.Println("Worker " + strconv.Itoa(id) + " processed " + request.Verb + " for " + request.Url + " in: " , now_in_millis() - start, "ms")
      results <- true
    }
  }
}

func now_in_millis() int64 {
  return time.Now().UnixNano() / int64(time.Millisecond)
}
