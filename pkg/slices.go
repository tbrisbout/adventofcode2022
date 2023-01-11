package pkg

import (
	"strconv"

	"golang.org/x/exp/slices"
)

func MustAtoi(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func Map[A, B any](list []A, f func(a A) B) (ret []B) {
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

func TakeRight[A any](list []A, n int) []A {
	return list[len(list)-n:]
}

// from stack overflow
func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

// Next returns the next element as if the list is a ring
func Next[T comparable](list []T, item T) T {
	if len(list) == 0 {
		// return zero value for T
		var t T
		return t
	}

	i := slices.Index(list, item)
	if i == -1 || i >= len(list)-1 {
		return list[0]
	}

	return list[i+1]
}

// IsNext checks if a is next to b as if the list is a ring
func IsNext[T comparable](list []T, a, b T) bool {
	return Next(list, b) == a
}

// Prev returns the previous element as if the list is a ring
func Prev[T comparable](list []T, item T) T {
	if len(list) == 0 {
		// return zero value for T
		var t T
		return t
	}

	i := slices.Index(list, item)
	if i <= 0 {
		return list[len(list)-1]
	}

	return list[i-1]
}

// IsPrev checks if a is prev to b as if the list is a ring
func IsPrev[T comparable](list []T, a, b T) bool {
	return Prev(list, b) == a
}
