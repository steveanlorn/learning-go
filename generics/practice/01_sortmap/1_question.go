package main

import (
	"fmt"
	"sort"
)

// sortByKey sorts value from the map based on the key
// in ascending order.
func sortByKey(m map[int64]string) []string {
	keys := make([]int64, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	result := make([]string, len(m))
	for i, key := range keys {
		result[i] = m[key]
	}

	return result
}

func main() {
	map1 := map[int64]string{
		3: "three",
		1: "one",
		2: "two",
	}

	// Practice:
	// Create a generic version of function sortByKey
	// so map2 & map3 can be sorted too.

	// map2 := map[string]int64{
	// 	"c": 3,
	// 	"a": 1,
	// 	"b": 2,
	// }

	// map3 := map[float64]bool{
	// 	3.1: true,
	// 	1.2: false,
	// 	2.4: false,
	// }

	fmt.Println(sortByKey(map1))
}
