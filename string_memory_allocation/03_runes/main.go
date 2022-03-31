package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "Hello"
	//s := "HelᏝo"

	for position, value := range s {
		fmt.Printf("position:%d, type:%v, binary:%08b, glyph:%c, value:%d\n", position, reflect.TypeOf(value), value, value, value)
	}
}
