package pkg

import (
	"strconv"
)

func MustAtoi(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func Map[A, B any](list []A, f func(a A) B) []B {
	ret := make([]B, len(list))
	for _, n := range list {
		ret = append(ret, f(n))
	}

	return ret
}

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Sum[A number](list []A) A {
	var sum A
	for _, a := range list {
		sum += a
	}
	return sum
}
