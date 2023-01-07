package main

import "testing"

func TestSumDuplicates(t *testing.T) {
	input := `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	got := sumDuplicates(input)
	want := 157

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}
}

func TestSumDuplicatesInGroups(t *testing.T) {
	input := `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	got := sumDuplicatesInGroups(input)
	want := 70

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}
}
