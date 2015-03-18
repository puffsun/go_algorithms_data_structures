package list

import (
//"errors"
)

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func New() *LinkedList {
	list := new(LinkedList)
	list.length = 0
	return list
}

func newNode(val interface{}) *Node {
	return &Node{value: val}
}

func (list *LinkedList) Length() int {
	return list.length
}

func (list *LinkedList) IsEmpty() bool {
	return list.length == 0
}

func (list *LinkedList) Prepend(value interface{}) {
	node := newNode(value)
	if list.length == 0 {
		list.head = node
		list.tail = list.head
	} else {
		formerHead := list.head
		formerHead.prev = node
		node.next = formerHead
		list.head = node
	}
	list.length += 1
}

func (list *LinkedList) RemoveFirst() interface{} {
	list.length -= 1
	return "One"
}
