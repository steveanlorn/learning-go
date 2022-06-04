package mergesort

import (
	"os"
	"testing"
)

// n contains the data to sort.
var n []int

// https://pkg.go.dev/testing#hdr-Main
func TestMain(m *testing.M) {
	for i := 0; i < 1_000_000; i++ {
		n = append(n, i)
	}

	code := m.Run()
	os.Exit(code)
}

func BenchmarkSort(b *testing.B) {
	b.Run("single", benchSingle)
	b.Run("unlimited", benchUnlimited)
	b.Run("numCPU", benchNumCPU)
}

func benchSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Single(n)
	}
}

func benchUnlimited(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unlimited(n)
	}
}

func benchNumCPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumCPU(n, 0)
	}
}
