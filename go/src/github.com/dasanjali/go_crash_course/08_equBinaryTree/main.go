/* 1. Implement the Walk function.
2. Test the Walk function.
3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
4. Test the Same function.*/
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// TreeVal_sender walks the tree t sending all values
// from the tree to the channel ch.
func treeVal_sender(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Similar_val determines whether the trees
// t1 and t2 contain the same values.
func Similar_val(tr1, tr2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go treeVal_sender(tr1, ch1)
	go treeVal_sender(tr2, ch2)

	for {
		i1, pass1 := <-ch1
		i2, pass2 := <-ch2

		if i1 != i2 || pass1 != pass2 {
			return false
		}

		if !pass1 {
			break
		}
	}

	return true
}

func main() {
	fmt.Println("When new 1 and 1 are same: ", Similar_val(tree.New(1), tree.New(1)))
	fmt.Println("When new 1 and 2 are same", Similar_val(tree.New(1), tree.New(2)))

}
