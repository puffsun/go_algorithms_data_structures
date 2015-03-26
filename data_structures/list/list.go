package list

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%s", n.value)
}

func (n *Node) Value() interface{} {
	return n.value
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

func (list *LinkedList) Append(value interface{}) {
	node := newNode(value)
	if list.length == 0 {
		list.tail = node
		list.head = list.tail
	} else {
		formerTail := list.tail
		formerTail.next = node
		node.prev = formerTail
		list.tail = node
	}
	list.length += 1
}

func (list *LinkedList) RemoveFirst() interface{} {
	if list.length == 0 {
		return nil
	}

	firstNode := list.head
	list.head = list.head.next

	if list.head != nil {
		list.head.prev = nil
	}
	list.length -= 1
	return firstNode.value
}

func (list *LinkedList) RemoveLast() interface{} {
	if list.length == 0 {
		return nil
	}

	lastNode := list.tail
	list.tail = list.tail.prev
	if list.tail != nil {
		list.tail.next = nil
	}
	list.length -= 1
	return lastNode.value
}

func (list *LinkedList) Clear() {
	list.length = 0
	list.head = nil
	list.tail = nil
}

func (list *LinkedList) Get(index int) (interface{}, error) {
	result, err := list.getNode(index)
	if err != nil {
		return nil, err
	} else {
		return result.value, nil
	}
}

func (list *LinkedList) Insert(index int, value interface{}) error {
	if index > list.length {
		return errors.New("Index out of range")
	}

	node := newNode(value)
	if list.length == 0 || index == 0 {
		list.Prepend(value)
		return nil
	}

	if list.length-1 == index {
		list.Append(value)
		return nil
	}

	nextNode, _ := list.getNode(index)
	prevNode := nextNode.prev
	prevNode.next = node
	node.prev = prevNode
	nextNode.prev = node
	node.next = nextNode
	list.length += 1
	return nil
}

func (list *LinkedList) Remove(value interface{}) error {
	if list.length == 0 {
		return errors.New("Empty list")
	}

	if list.head.value == value {
		list.RemoveFirst()
		return nil
	}

	found := false
	for node := list.head; node != nil; node = node.next {
		if node.value == value {
			if node.next != nil {
				node.next.prev = node.prev
			}

			if node.prev != nil {
				node.prev.next = node.next
			}
			list.length -= 1
			found = true
		}
	}

	if !found {
		return errors.New("Node not found")
	}
	return nil
}

func (list *LinkedList) getNode(index int) (*Node, error) {
	if list.length < index || list.length == 0 {
		return nil, errors.New("Index out of range")
	}

	node := list.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node, nil
}

func (list *LinkedList) Each(f func(node *Node)) {
	for node := list.head; node != nil; node = node.next {
		f(node)
	}
}
