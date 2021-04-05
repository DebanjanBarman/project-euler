package main

import (
	"testing"
)

func TestGetSumOfTwoMultiPles(t *testing.T) {
	got := GetSumOfTwoMultiPles(3, 5, 10)

	want := 23

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
