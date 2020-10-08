package _2_BST_recusion

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	bst := NewBST()
	bst.Add(10)
	bst.Add(8)
	bst.Add(1)
	bst.Add(13)
	fmt.Println(bst)

	fmt.Println(bst.Contains(3))
	fmt.Println(bst.Contains(13))
}
