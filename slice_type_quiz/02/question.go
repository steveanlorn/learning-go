package main

import "fmt"

// Question 2
// Given a slice with length and capacity value of 5,
// what is the value of the slice, length and
// capacity of the slice after append execution?

func main() {
	data := make([]int, 5)

	for record := 1; record <= 5; record++ {
		fmt.Printf("Value: %v, len %d, cap %d\n", data, len(data), cap(data))
		data = append(data, record)
	}

	fmt.Printf("Value: %v, len %d, cap %d\n", data, len(data), cap(data))
}
