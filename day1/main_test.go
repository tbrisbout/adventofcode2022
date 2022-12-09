package main

import "testing"

func TestCountCalories(t *testing.T) {
	input := `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

	got := countCalories(input)
	want := 24000

	if got != want {
		t.Fatalf("expected %d but got %d", want, got)
	}
}
