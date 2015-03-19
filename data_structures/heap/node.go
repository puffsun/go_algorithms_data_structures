package heap

type Node struct {
	value  int
	parent *Node
}

func NewNode(val int, p *Node) *Node {
	return &Node{val, p}
}

func (node *Node) Parent() *Node {
	return node.parent
}

func (node *Node) SetParent(p *Node) {
	node.parent = p
}

func (node *Node) Value() int {
	return node.value
}

func (node *Node) SetValue(val int) {
	node.value = val
}

func (node *Node) Less(n *Node) bool {
	return node.value < n.value
}
