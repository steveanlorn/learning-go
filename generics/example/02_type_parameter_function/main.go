package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// type parameters:
// make the function generic, enabling it to work with arguments of different types

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
	// To call generic functions, we provide type and ordinary arguments.
	fmt.Println(min[int64]([]int64{5, 4, 3, 2, 1}))
	fmt.Println(min[float64]([]float64{3.2, 1.1, 2, 8}))
	fmt.Println(min[string]([]string{"b", "c", "a"}))
	fmt.Println(min[int32]([]int32{}))

	// Instantiation is providing type argument to the function.
	// It produces a non-generic function.
	// Two steps of instantiation:
	// 1. Substitute type arguments for type parameters.
	//    At compile time, the type parameter stands for a single type,
	//    the type provided as a type argument by the calling code.
	// 2. Check that type arguments implement their constraints.
	//    Type constraint specifies the permissible type arguments
	//    that calling code can use for the respective type parameter.
}
