package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Type Assertion
// cons:
// - Losing type-safe
// - Type assertion both in caller and function code
// - Caller needs to wrap arguments in interface{}

func minAssert(s []interface{}) (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("no values given")
	}

	switch first := s[0].(type) {
	case int:
		min := first
		for _, rawV := range s[1:] {
			v := rawV.(int)
			if v < min {
				min = v
			}
		}
		return min, nil
	case float64:
		min := first
		for _, rawV := range s[1:] {
			v := rawV.(float64)
			if v < min {
				min = v
			}
		}
		return min, nil
	default:
		return nil, fmt.Errorf("unsupported element type of given slice: %T", first)
	}
}

// Reflection
// cons:
// - Abysmal readability
// - Losing type-safe
// - Lower performance that in other approaches

func minReflect(s []interface{}) (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("no values given")
	}

	first := reflect.ValueOf(s[0])

	if first.CanInt() {
		min := first.Int()
		for _, rawV := range s[1:] {
			v := reflect.ValueOf(rawV)
			if v.CanInt() {
				intV := v.Int()
				if intV < min {
					min = intV
				}
			}
		}
		return min, nil
	}

	if first.CanFloat() {
		min := first.Float()
		for _, rawV := range s[1:] {
			v := reflect.ValueOf(rawV)
			if v.CanFloat() {
				floatV := v.Float()
				if floatV < min {
					min = floatV
				}
			}
		}
		return min, nil
	}

	return nil, fmt.Errorf("unsupported element type of given slice: %T", s[0])
}

func main() {
	m1, err1 := minAssert([]interface{}{5, 4, 3, 2, 1})
	m2, err2 := minAssert([]interface{}{3.2, 1.1, 2.8})
	fmt.Println("By type assertion:")
	fmt.Println(err1, err2)
	fmt.Println(m1, m2)

	m3, err3 := minReflect([]interface{}{5, 4, 3, 2, 1})
	m4, err4 := minReflect([]interface{}{3.2, 1.1, 2.8})
	fmt.Println("By reflection:")
	fmt.Println(err3, err4)
	fmt.Println(m3, m4)
}
