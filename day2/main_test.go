package main

import "testing"

func TestComputeScore(t *testing.T) {
	input := `
A Y
B X
C Z
`

	got := computeScore(input)
	want := 15

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}
}

func TestComputeScoreWithRequestedEnding(t *testing.T) {
	input := `
A Y
B X
C Z
`

	got := computeScoreWithEnding(input)
	want := 12

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}

}
