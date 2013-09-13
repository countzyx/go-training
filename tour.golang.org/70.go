package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
	"os"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkSubtree(t, ch)

	close(ch)
}

func walkSubtree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkSubtree(t.Left, ch)
	}
	ch <- t.Value

	if t.Right != nil {
		walkSubtree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	bIsSame := true

	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for result1 := range ch1 {
		result2 := <-ch2
		bIsSame = result1 == result2
		if !bIsSame {
			break
		}
	}

	return bIsSame
}

func main() {
	tree1_1 := tree.New(1)
	tree1_2 := tree.New(1)
	tree2_1 := tree.New(2)
	tree2_2 := tree.New(2)

	fmt.Printf("Tree1_1 == Tree1_2? %t\n", Same(tree1_1, tree1_2))
	fmt.Printf("Tree2_1 == Tree2_2? %t\n", Same(tree2_1, tree2_2))
	fmt.Printf("Tree1_1 == Tree2_1? %t\n", Same(tree1_1, tree2_1))
	fmt.Printf("Tree1_1 == Tree2_2? %t\n", Same(tree1_1, tree2_2))
	fmt.Printf("Tree1_2 == Tree2_1? %t\n", Same(tree1_2, tree2_1))
	fmt.Printf("Tree1_2 == Tree2_2? %t\n", Same(tree1_2, tree2_2))

	os.Exit(0)
}
