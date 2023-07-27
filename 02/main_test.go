package main

import (
	"testing"
)

func TestGetTotalScore(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}

	want := 15
	got := GetTotalScore(input)

	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestGetTotalScorePart2(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}

	want := 12
	got := GetTotalScorePart2(input)

	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}
