package main

import (
	"fmt"
	"unsafe"
)

// The size of a variable can be determined by using unsafe.Sizeof(a).
// The result will remain the same for a given type (i.e. int, int64, string, struct etc),
// irrespective of the value it holds.
// However, for type string, you may be interested in the size of the string that the variable references,
// and this is determined by using len(a) function on a given string.
// The following snippet illustrates that the size of a variable of type string is
// 16 but the length of a string that variable references can vary:

func main() {
	s1 := "foo"
	s2 := "foobar"

	fmt.Printf("s1 size: %T, %d\n", s1, unsafe.Sizeof(s1))
	fmt.Printf("s2 size: %T, %d\n", s2, unsafe.Sizeof(s2))
	fmt.Printf("s1 len: %T, %d\n", s1, len(s1))
	fmt.Printf("s2 len: %T, %d\n", s2, len(s2))
}
