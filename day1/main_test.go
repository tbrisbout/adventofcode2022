package main

import (
	"fmt"
	"testing"
)

func TestCountNTopHolders(t *testing.T) {
	const input = `
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
	tc := []struct{ Number, Want int }{
		{1, 24000},
		{3, 45000},
	}

	for _, test := range tc {
		t.Run(fmt.Sprintf("count for %d returns %d", test.Number, test.Want), func(t *testing.T) {
			got := countNTopHolders(input, test.Number)
			want := test.Want

			if got != want {
				t.Errorf("expected %d but got %d", want, got)
			}
		})
	}
}
