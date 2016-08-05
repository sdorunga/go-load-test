package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	stats := StatsPrinter{[]int{20, 30, 40}}
	if stats.sum() != 90 {
		t.Errorf("Expected sum to be 90, is %d", stats.sum())
	}
}

func TestAverage(t *testing.T) {
	stats := StatsPrinter{[]int{20, 30, 40}}
	if stats.average() != 30 {
		t.Errorf("Expected average to be 30, is %d", stats.average())
	}
}

func TestMin(t *testing.T) {
	stats := StatsPrinter{[]int{40, 30, 40, 31, 50}}
	if stats.min() != 30 {
		t.Errorf("Expected minimum to be 30, is %d", stats.min())
	}
}

func TestMax(t *testing.T) {
	stats := StatsPrinter{[]int{40, 30, 40, 31, 50}}
	if stats.max() != 50 {
		t.Errorf("Expected max to be 50, is %d", stats.max())
	}
}

func TestMedian(t *testing.T) {
	stats := StatsPrinter{[]int{45, 30, 40, 31, 50}}
	if stats.median() != 40 {
		t.Errorf("Expected median to be 40, is %d", stats.median())
	}

	stats = StatsPrinter{[]int{42, 30, 48, 50}}
	if stats.median() != 45 {
		t.Errorf("Expected median to be 45, is %d", stats.median())
	}
}