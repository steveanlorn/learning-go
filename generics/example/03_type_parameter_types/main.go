package main

type Tree[T interface{}] struct {
	left, right *Tree[T]
	data        T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	return nil
}

var stringTree Tree[string]
var intTree Tree[int]

func main() {

}
