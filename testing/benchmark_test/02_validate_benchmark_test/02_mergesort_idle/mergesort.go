package mergesort

// Codes from
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/testing/benchmarks/validate/validate_test.go

import (
	"math"
	"runtime"
	"sync"
)

// Single uses a single goroutine to perform the merge sort.
func Single(n []int) []int {

	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Sort the left side.
	l := Single(n[:i])

	// Sort the right side.
	r := Single(n[i:])

	// Place things in order and merge ordered lists.
	return merge(l, r)
}

// Unlimited uses a goroutine for every split to perform the merge sort.
func Unlimited(n []int) []int {

	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Maintain the ordered left and right side lists.
	var l, r []int

	// For each split we will have 2 goroutines.
	var wg sync.WaitGroup
	wg.Add(2)

	// Sort the left side concurrently.
	go func() {
		l = Unlimited(n[:i])
		wg.Done()
	}()

	// Sort the right side concurrently.
	go func() {
		r = Unlimited(n[i:])
		wg.Done()
	}()

	// Wait for the splitting to end.
	wg.Wait()

	// Place things in order and merge ordered lists.
	return merge(l, r)
}

// NumCPU uses the same number of goroutines that we have cores
// to perform the merge sort.
func NumCPU(n []int, lvl int) []int {

	// Once we have a list of one we can begin to merge values.
	if len(n) <= 1 {
		return n
	}

	// Split the list in half.
	i := len(n) / 2

	// Maintain the ordered left and right side lists.
	var l, r []int

	// Cacluate how many levels deep we can create goroutines.
	// On an 8 core machine we can keep creating goroutines until level 4.
	// 		Lvl 0		1  Lists		1  Goroutine
	//		Lvl 1		2  Lists		2  Goroutines
	//		Lvl 2		4  Lists		4  Goroutines
	//		Lvl 3		8  Lists		8  Goroutines
	//		Lvl 4		16 Lists		16 Goroutines

	// On 8 core machine this will produce the value of 3.
	maxLevel := int(math.Log2(float64(runtime.NumCPU())))

	// We don't need more goroutines then we have logical processors.
	if lvl <= maxLevel {
		lvl++

		// For each split we will have 2 goroutines.
		var wg sync.WaitGroup
		wg.Add(2)

		// Sort the left side concurrently.
		go func() {
			l = NumCPU(n[:i], lvl)
			wg.Done()
		}()

		// Sort the right side concurrently.
		go func() {
			r = NumCPU(n[i:], lvl)
			wg.Done()
		}()

		// Wait for the splitting to end.
		wg.Wait()

		// Place things in order and merge ordered lists.
		return merge(l, r)
	}

	// Sort the left and right side on this goroutine.
	l = NumCPU(n[:i], lvl)
	r = NumCPU(n[i:], lvl)

	// Place things in order and merge ordered lists.
	return merge(l, r)
}

// merge performs the merging to the two lists in proper order.
func merge(l, r []int) []int {

	// Declare the sorted return list with the proper capacity.
	ret := make([]int, 0, len(l)+len(r))

	// Compare the number of items required.
	for {
		switch {
		case len(l) == 0:
			// We appended everything in the left list so now append
			// everything contained in the right and return.
			return append(ret, r...)

		case len(r) == 0:
			// We appended everything in the right list so now append
			// everything contained in the left and return.
			return append(ret, l...)

		case l[0] <= r[0]:
			// First value in the left list is smaller than the
			// first value in the right so append the left value.
			ret = append(ret, l[0])

			// Slice that first value away.
			l = l[1:]

		default:
			// First value in the right list is smaller than the
			// first value in the left so append the right value.
			ret = append(ret, r[0])

			// Slice that first value away.
			r = r[1:]
		}
	}
}
