package main

import "fmt"

// Vertical bar symbol denotes union of types
// ~ symbol defines all type with underlying type is int32, for example.

// inline constraint
func add[T interface{ float64 | int64 | ~int32 }](x, y T) T {
	return x + y
}

// omitting interface
func min[T float64 | int64 | ~int](x, y T) T {
	return x - y
}

// interface still abe to set method or embed another interface
type numericConstraint interface {
	float64 | int64 | ~int32
}

func multiply[T numericConstraint](x, y T) T {
	return x * y
}

// type constraint can reference other type parameters -------
func printSlice[S ~[]E, E interface{}](s S) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type sliceConstraint[E interface{}] interface {
	~[]E
}

func firstElem[S sliceConstraint[E], E interface{}](s S) E {
	return s[0]
}

// any is a new built-in type constraint, an alias for interface{}
// attention: make sure operations in function body supported by any type.
func print[T any](s T) {
	fmt.Println(s)
}

// comparable is a new built-in type constraint, any types whose values may be used
// as an operand of the comparison operators == and !=
func isEqual[T comparable](a, b T) bool {
	return a == b
}

func main() {
	fmt.Println(add[float64](1.1, 2.2))
	fmt.Println(min[int](5, 2))

	type myInt32 int32
	var i myInt32 = 2
	fmt.Println(multiply[myInt32](1, i))

	printSlice[[]string]([]string{"Hi", "aloha"})
	printSlice[[]int]([]int{1, 2})

	fmt.Println(firstElem[[]bool]([]bool{false, true, true}))

	print[string]("Hi")

	fmt.Println(isEqual[int](1, 2))
	fmt.Println(isEqual[string]("a", "a"))
}
