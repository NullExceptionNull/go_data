package _1_BST

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
	if b.root == nil {
		b.root = NewNode(e)
		b.size++
	} else {
		b.add(b.root, e)
	}
}

func (b *BST) add(n *Node, e int) {

	if n.e == e {
		return
	}
	if e < n.e {
		if n.left == nil {
			n.left = NewNode(e)
			b.size++
			return
		} else {
			b.add(n.left, e)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(e)
			b.size++
			return
		} else {
			b.add(n.right, e)
		}
	}

}
