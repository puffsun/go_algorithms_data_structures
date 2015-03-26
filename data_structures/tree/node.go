package tree

type Key int

type Node struct {
	key         Key
	value       interface{}
	left        *Node
	right       *Node
	numChildren int
}

func MakeNode(key Key, value interface{}, num int) *Node {
	n := new(Node)
	n.key = key
	n.value = value
	n.numChildren = num
	return n
}
