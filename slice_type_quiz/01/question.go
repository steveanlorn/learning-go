package main

import "fmt"

// Question 1
// Given two ways of declaring a slice (line 9 & 10),
// what is the output of each slice's (slice a and slice b) len, cap, address, and nil or not nil?

func main() {
	var a []string
	b := []string{}

	fmt.Printf("len %d, cap %d, address %p, %t\n", len(a), cap(a), a, a == nil)
	fmt.Printf("len %d, cap %d, address %p, %t\n", len(b), cap(b), b, b == nil)
}
