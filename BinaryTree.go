package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

var root *tree.Tree

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if root == nil {
		root = t
	}

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}

	if t == root {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	cap := 10

	ch1 := make(chan int, cap)
	ch2 := make(chan int, cap)

	go Walk(tree.New(1), ch1)
	go Walk(tree.New(1), ch2)

	for v := range ch1 {
		
	}

	return true
}

func main() {

	root = nil
	ch1 := make(chan int)
	t1 := tree.New(1)

	go Walk(t1, ch1)

	for v := range ch1 {
		fmt.Println("Received ", v)
	}
}
