package main

import "fmt"

// Question 5
// What is the output on line 20 and 22?

func main() {
	fruits1 := make([]string, 5, 8)
	fruits1[0] = "apple"
	fruits1[1] = "banana"
	fruits1[2] = "orange"
	fruits1[3] = "grape"
	fruits1[4] = "melon"

	fruits2 := fruits1[2:4]
	fruits2[0] = "coconut"
	fruits2 = append(fruits2, "avocado")

	inspectSlice(fruits1)
	fmt.Println("-----------------")
	inspectSlice(fruits2)
}

func inspectSlice(s []string) {
	for i, value := range s {
		fmt.Printf("[%d] %p - %s\n", i, &s[i], value)
	}
}
