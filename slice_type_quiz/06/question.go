package main

import "fmt"

// Question 6
// What is the number of likes for each users in the end of the program?

type user struct {
	likes int
}

func main() {
	users := make([]user, 3)

	shareUser := &users[1]
	shareUser.likes++

	for i := range users {
		fmt.Printf("User: %d, likes: %d\n", i, users[i])
	}

	users = append(users, user{})
	shareUser.likes++

	fmt.Println("***************************")

	for i := range users {
		fmt.Printf("User: %d, likes: %d\n", i, users[i])
	}
}
