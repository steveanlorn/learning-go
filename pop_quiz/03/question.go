package main

import "fmt"

// Question 3
// We increment each value in the slice by one.
// What is the output of line 16?

func main() {
	data := []int{1,2,3,4,5}
	for _, value := range data {
		value++
	}

	for i, value := range data {
		fmt.Printf("Index %d - %d\n", i, value)
	}
}
