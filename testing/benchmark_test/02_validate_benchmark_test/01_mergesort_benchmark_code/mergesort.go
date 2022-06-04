package mergesort

// Codes from
// https://github.com/ardanlabs/gotraining/blob/master/topics/go/testing/benchmarks/validate/validate_test.go

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
