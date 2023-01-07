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

	t.Run("Sum Duplicates in lines splitted by half", func(t *testing.T) {
		got := sumDuplicates(splitByHalf(input))
		want := 157

		if got != want {
			t.Fatalf("expected %d but got %d", want, got)
		}
	})

	t.Run("Sum Duplicates in Chunk of 3 lines", func(t *testing.T) {
		got := sumDuplicates(splitByChunk(input))
		want := 70

		if got != want {
			t.Fatalf("expected %d but got %d", want, got)
		}
	})
}
