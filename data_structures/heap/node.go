package heap

type Node struct {
	value    int
	priority int
	parent   *Node
}

func NewNode(val int, p *Node) *Node {
	return &Node{val, -1, p}
}

func (node *Node) Priority() int {
	return node.priority
}

func (node *Node) SetPriority(priority int) {
	node.priority = priority
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
	// priority is -1 means no priority at all
	if node.priority == -1 {
		return node.value < n.value
	} else {
		return node.priority < n.priority
	}
}
