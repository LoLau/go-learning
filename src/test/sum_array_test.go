package test

import (
	"testing"
)

func TestSumArray(t *testing.T) {
	arr := []int {1, 2, 3, 4}
	got := sumArray(arr)
	want := 10
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func sumArray(arr []int) int {
	sum := 0
	for _, i  := range arr {
		sum += i
	}
	return sum
}
