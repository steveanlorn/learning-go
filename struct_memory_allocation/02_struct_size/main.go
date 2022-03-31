package main

import (
	"fmt"
	"unsafe"
)

type example1 struct {
	flag bool // 1 byte
	// [1]byte padding
	counter int16   // 2 bytes
	pi      float32 // 4 bytes
}

type example2 struct {
	flag bool // 1 byte
	// [7]byte padding
	counter int64   // 8 bytes
	pi      float32 // 4 bytes
	// [4]byte padding
}

type example2Optimized struct {
	counter int64   // 8 bytes
	pi      float32 // 4 bytes
	flag    bool    // 1 byte
	// [3]byte padding
}

func main() {
	var e example1
	fmt.Printf("%T: %d bytes, offset %+v\n", e.flag, unsafe.Sizeof(e.flag), unsafe.Offsetof(e.flag))
	fmt.Printf("%T %d bytes, offset %+v\n", e.counter, unsafe.Sizeof(e.counter), unsafe.Offsetof(e.counter))
	fmt.Printf("%T %d bytes, offset %+v\n", e.pi, unsafe.Sizeof(e.pi), unsafe.Offsetof(e.pi))
	fmt.Println("----------------")
	fmt.Printf("%T %d bytes, allignment %d\n", e, unsafe.Sizeof(e), unsafe.Alignof(e))
}
