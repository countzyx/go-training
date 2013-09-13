package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
	"os"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkSubtree(t, ch)

	close(ch)
}

func WalkSubtree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkSubtree(t.Left, ch)
	}
	ch <- t.Value

	if t.Right != nil {
		WalkSubtree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for result := range ch {
		fmt.Println(result)
	}

	os.Exit(0)
}
