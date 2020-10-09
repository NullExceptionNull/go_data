package _2_BST_recusion

import (
	"bytes"
	"fmt"
	"strconv"
)

type Node struct {
	e     int
	left  *Node
	right *Node
}

func NewNode(e int) *Node {
	return &Node{e: e}
}

type BST struct {
	root *Node
	size int
}

func (b *BST) String() string {
	var buffer = bytes.NewBufferString("")
	generateBSTString(b.root, 0, buffer)
	return buffer.String()
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Add(e int) {
	b.root = b.add(b.root, e)
}

func (b *BST) add(node *Node, e int) *Node {
	if node == nil {
		b.size++
		return NewNode(e)
	}
	if e < node.e {
		node.left = b.add(node.left, e)
	} else if e > node.e {
		node.right = b.add(node.right, e)
	}
	return node
}

func (b *BST) Contains(e int) bool {
	return b.contains(b.root, e)
}

func (b *BST) contains(node *Node, e int) bool {
	if node == nil {
		return false
	}
	if node.e == e {
		return true
	} else if e < node.e {
		return b.contains(node.left, e)
	} else if e > node.e {
		return b.contains(node.right, e)
	}
	return false
}

func (b *BST) PreOrder() {
	b.preOrder(b.root)
}

func (b *BST) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.e)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

func (b *BST) InnerOrder() {
	b.preOrder(b.root)
}

func (b *BST) innerOrder(node *Node) {
	if node == nil {
		return
	}
	b.innerOrder(node.left)
	fmt.Println(node.e)
	b.innerOrder(node.right)
}

func generateBSTString(node *Node, dept int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDeptString(dept) + "NULL\n")
		return
	}
	buffer.WriteString(generateDeptString(dept) + strconv.Itoa(node.e) + "\n")
	generateBSTString(node.left, dept+1, buffer)
	generateBSTString(node.right, dept+1, buffer)
}

func generateDeptString(dept int) string {
	var buffer = bytes.NewBufferString("")

	for i := 0; i < dept; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
