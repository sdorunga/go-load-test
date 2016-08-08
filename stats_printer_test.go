package main

import (
	"testing"
)

func TestRequestsCount(t *testing.T) {
	stats := StatsPrinter{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 1000}
	if stats.requestCount() != 10 {
		t.Errorf("Expected the 10 requests, got %d", stats.requestCount())
	}
}

func TestSum(t *testing.T) {
	stats := StatsPrinter{[]int{20, 30, 40}, 1000}
	if stats.sum() != 90 {
		t.Errorf("Expected sum to be 90, is %d", stats.sum())
	}
}

func TestAverage(t *testing.T) {
	stats := StatsPrinter{[]int{20, 30, 40}, 1000}
	if stats.average() != 30 {
		t.Errorf("Expected average to be 30, is %d", stats.average())
	}
}

func TestMin(t *testing.T) {
	stats := StatsPrinter{[]int{40, 30, 40, 31, 50}, 1000}
	if stats.min() != 30 {
		t.Errorf("Expected minimum to be 30, is %d", stats.min())
	}
}

func TestMax(t *testing.T) {
	stats := StatsPrinter{[]int{40, 30, 40, 31, 50}, 1000}
	if stats.max() != 50 {
		t.Errorf("Expected max to be 50, is %d", stats.max())
	}
}

func TestMedian(t *testing.T) {
	stats := StatsPrinter{[]int{45, 30, 40, 31, 50}, 1000}
	if stats.median() != 40 {
		t.Errorf("Expected median to be 40, is %d", stats.median())
	}

	stats = StatsPrinter{[]int{42, 30, 48, 50}, 1000}
	if stats.median() != 45 {
		t.Errorf("Expected median to be 45, is %d", stats.median())
	}
}

func TestPercentile(t *testing.T) {
	stats := StatsPrinter{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 1000}
	if stats.percentile(0.70) != 7 {
		t.Errorf("Expected the 70th percentile to return the 7th number, returned %d", stats.percentile(0.70))
	}
}

func TestPercentileWithRounding(t *testing.T) {
	stats := StatsPrinter{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 1000}
	if stats.percentile(0.75) != 8 {
		t.Errorf("Expected the 75th percentile to return the 8th number, returned %d", stats.percentile(0.75))
	}
}

func TestRequestsPerSecond(t *testing.T) {
	stats := StatsPrinter{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 1000}
	if stats.requestPerSecond() != 10 {
		t.Errorf("Expected the 10 requests per second, got %f", stats.requestPerSecond())
	}
}
