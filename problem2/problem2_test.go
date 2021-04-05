package main

import (
	"testing"
)

func TestGetSumOfTwoMultiPles(t *testing.T) {
	got := SumOfEvenFibonacciValues(40)

	want := 44

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
