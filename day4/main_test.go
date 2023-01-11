package main

import "testing"

func TestCountContainedRanges(t *testing.T) {
	input := `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	got := countContainedRanges(input)
	want := 2

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}
}
func TestCountOverlappingRanges(t *testing.T) {
	input := `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	got := countOverlappingRanges(input)
	want := 4

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}
}
