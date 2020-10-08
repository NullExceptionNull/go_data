package _2_BST_recusion

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	bst := NewBST()
	bst.Add(5)
	bst.Add(3)
	bst.Add(6)
	bst.Add(8)
	bst.Add(4)
	bst.Add(2)
	//bst.PreOrder()
	fmt.Println(bst.String())

}
