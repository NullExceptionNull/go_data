package _1_linked_list_basic

import "fmt"

type Node struct {
	Next *Node
	E    interface{}
}

func (n *Node) String() string {
	return fmt.Sprint(n.E)
}
