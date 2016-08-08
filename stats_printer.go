package main

import (
	"fmt"
	"sort"
)

type StatsPrinter struct {
	times     []int
	totalTime int
}

func (this *StatsPrinter) Print() {
	fmt.Println("Number of Requests: ", this.requestCount())
	fmt.Println("Average: ", this.average(), "ms")
	fmt.Println("Median: ", this.median(), "ms")
	fmt.Println("Min: ", this.min(), "ms")
	fmt.Println("Max: ", this.max(), "ms")
	fmt.Println("90th Percentile: ", this.percentile(0.90), "ms")
	fmt.Println("95th Percentile: ", this.percentile(0.95), "ms")
	fmt.Println("99th Percentile: ", this.percentile(0.99), "ms")
	fmt.Printf("Requests/s %.2f\n", this.requestPerSecond())
	fmt.Println("Duration", this.totalTime, "ms")
}

func (this *StatsPrinter) average() int {
	if this.requestCount() == 0 {
		return 0
	}

	return this.sum() / this.requestCount()
}

func (this *StatsPrinter) median() int {
	if this.requestCount() == 0 {
		return 0
	}
	sort.Ints(this.times)

	middle := this.requestCount() / 2
	result := this.times[middle]
	if this.requestCount()%2 == 0 {
		result = (result + this.times[middle-1]) / 2
	}
	return result
}

func (this *StatsPrinter) sum() (total int) {
	for _, number := range this.times {
		total += number
	}
	return total
}

func (this *StatsPrinter) min() int {
	if this.requestCount() == 0 {
		return 0
	}

	min := this.times[0]
	for _, number := range this.times {
		if min > number {
			min = number
		}
	}
	return min
}

func (this *StatsPrinter) max() int {
	if this.requestCount() == 0 {
		return 0
	}

	max := this.times[0]
	for _, number := range this.times {
		if max < number {
			max = number
		}
	}
	return max
}

func (this *StatsPrinter) percentile(rank float32) int {
	if this.requestCount() == 0 {
		return 0
	}

	sort.Ints(this.times)
	index := rank * (float32(this.requestCount()))
	// The 0.5 is used to make int round up to the next integer
	// Does not work for negative numbers which times should never be
	return this.times[int(index-1+0.5)]
}

func (this *StatsPrinter) requestPerSecond() float32 {
	return float32(this.requestCount()) / (float32(this.totalTime) / 1000)
}

func (this *StatsPrinter) requestCount() int {
	return (len(this.times))
}
