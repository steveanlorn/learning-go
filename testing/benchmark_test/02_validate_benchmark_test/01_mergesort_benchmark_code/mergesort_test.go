package mergesort

import (
	"testing"
)

func BenchmarkSingle(b *testing.B) {
	// b.N specifies the number of iterations; the value is not fixed, but dynamically allocated,
	// ensuring that the benchmark runs for at least one second by default.
	for i := 0; i < b.N; i++ {
		//b.StopTimer()
		// n contains the data to sort.
		var n []int

		for i := 0; i < 1_000_000; i++ {
			n = append(n, i)
		}
		//b.StartTimer()
		Single(n)
	}
}
