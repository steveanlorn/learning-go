package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Scale multiplies each element in slice s with c.
func Scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	return fmt.Sprintf("%d", p[0])
}

func ScaleAndPrint(p Point) {
	r := Scale(p, 2)
	fmt.Println(r.String())
}

func main() {

	// Practice:
	// Make this code able to compile without changing the ScaleAndPrint functions.
	//
	// Hints:
	// It is not compiled because the Scale return slice of E,
	// where E is the type argument of slice.
	// so r is type []int32 & it does not have method

	p := Point{1, 2, 3}
	ScaleAndPrint(p)
}
