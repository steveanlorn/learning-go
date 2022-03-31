package main

import (
	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		return T(0)
	}

	min := s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

func main() {
	min[int64]([]int64{5, 4, 3, 2, 1})
	min([]float64{3.2, 1.1, 2, 8})
}
