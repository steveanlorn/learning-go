package main

import "fmt"

// Question 4
// We iterate over data using range.
// On line 14, we create a new slice from the data slice from index 0 up to 1.
// What is the output of line 15 and line 18?

func main() {
	data := []int{1,2,3,4,5}

	for _, value := range data {
		data = data[:2]
		fmt.Println(value)
	}

	fmt.Printf("Data: %v\n", data)
}
