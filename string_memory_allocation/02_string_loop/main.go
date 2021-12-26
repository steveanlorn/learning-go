package main

import (
	"fmt"
)

func main() {
	s := "Hello"
	//s := "HelᏝo"

	// same as fmt.Println([]byte(s))
	for i:=0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// verbose print
	//for i:=0; i < len(s); i++ {
	//	fmt.Printf("position:%d, type:%v, binary:%08b, glyph:%c, value:%d, code-point:%U\n", i, reflect.TypeOf(s[i]), s[i], s[i], s[i], s[i])
	//}
}
