package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkTree(t, ch)
	close(ch)
}

func WalkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkTree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	if <-ch1 != <-ch2 {
		return false
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for range [10]int{} {
		fmt.Println(<-ch)
	}
	fmt.Printf("Same(tree.New(1), tree.New(1)) returns %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("Same(tree.New(1), tree.New(2)) returns %v\n", Same(tree.New(1), tree.New(2)))
}
