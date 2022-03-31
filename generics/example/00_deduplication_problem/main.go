package main

import (
	"fmt"
)

// Explicit functions
// cons:
// - lots of manual labour
// - code duplication lower maintainability

func minInt64(s []int64) int64 {
	if len(s) == 0 {
		return 0
	}

	min := s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

func minFloat64(s []float64) float64 {
	if len(s) == 0 {
		return 0
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
	fmt.Println(minInt64([]int64{5, 4, 3, 2, 1}))
	fmt.Println(minFloat64([]float64{5.1, 4.2, 3.2, 2.1, 1.1}))
}
