package binarytree

const (
	preOrder  order = iota // NLR
	inOrder                // LNR
	postOrder              // LRN
)

type (
	Node struct {
		v     string
		left  *Node
		right *Node
	}
	WalkFunc func(v string, level int) bool
	order    uint
)

func NewTree(v string, l *Node, r *Node) *Node {
	return &Node{
		v:     v,
		left:  l,
		right: r,
	}
}

func (n *Node) Value() string {
	return n.v
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) WalkPreOrder(f WalkFunc) {
	walkRecursion(preOrder, n, 0, f)
}

func (n *Node) WalkInOrder(f WalkFunc) {
	walkRecursion(inOrder, n, 0, f)
}

func (n *Node) WalkPostOrder(f WalkFunc) {
	walkRecursion(postOrder, n, 0, f)
}

func (n *Node) WalkReverseInOrder(f WalkFunc) {
	walkRecursion(postOrder, n, 0, f)
}

// https://en.wikipedia.org/wiki/Tree_traversal#Depth-first_search_of_binary_tree
func walkRecursion(o order, n *Node, level int, f WalkFunc) bool {
	if n == nil {
		return true
	}
	if o == preOrder && !f(n.v, level) {
		return false
	}
	if !walkRecursion(o, n.left, level+1, f) {
		return false
	}
	if o == inOrder && !f(n.v, level) {
		return false
	}
	if !walkRecursion(o, n.right, level+1, f) {
		return false
	}
	if o == postOrder && !f(n.v, level) {
		return false
	}
	return true
}
