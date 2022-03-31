package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func ScaleOptimized[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func ScaleAndPrintOptimized(p Point) {
	r := ScaleOptimized(p, 2)
	fmt.Println(r.String())
}
