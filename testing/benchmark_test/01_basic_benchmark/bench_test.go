package mergsesort

import (
	"fmt"
	"testing"
)

var result string

func BenchmarkPrint(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = fmt.Sprint("Hello!!")
	}
	result = r
}

func BenchmarkPrintf(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = fmt.Sprintf("Hello!!")
	}
	result = r
}
