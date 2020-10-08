package _2_BST_recusion

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
