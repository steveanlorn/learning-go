package main

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func sortByKeyGeneric[K constraints.Ordered, V any](m map[K]V) []V {
	keys := make([]K, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	result := make([]V, len(m))
	for i, key := range keys {
		result[i] = m[key]
	}

	return result
}
