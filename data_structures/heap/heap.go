package heap

import (
	"errors"
	"sync"
)

type Heap struct {
	sync.Mutex
	data []*Node
	min  bool
}

func New() *Heap {
	// By default, create min heap
	return NewMin()
}

func NewMin() *Heap {
	return &Heap{
		data: make([]*Node, 0),
		min:  true,
	}
}

func NewMax() *Heap {
	return &Heap{
		data: make([]*Node, 0),
		min:  false,
	}
}

func (h *Heap) Length() int {
	return len(h.data)
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

/**
 * @value should implement method Less
 *
 */
func (h *Heap) Insert(value int) {
	h.Lock()
	defer h.Unlock()

	node := NewNode(value, nil)
	h.data = append(h.data, node)
	h.siftUp()
}

func (h *Heap) Extract() (int, error) {
	h.Lock()
	defer h.Unlock()
	if h.Length() == 0 {
		return -1, errors.New("Empty heap")
	}

	firstNode := h.data[0]
	lastNode := h.data[h.Length()-1]
	if h.Length() == 1 {
		h.data = nil
		return firstNode.value, nil
	}
	h.data = append([]*Node{lastNode}, h.data[1:h.Length()-1]...)
	h.siftDown()
	return firstNode.value, nil
}

func (h *Heap) Less(a, b *Node) bool {
	if h.min {
		return a.Less(b)
	} else {
		return b.Less(a)
	}
}

func (h *Heap) siftUp() {
	for i, parent := h.Length()-1, h.Length()-1; i > 0; i = parent {
		parent = i >> 1
		if h.Less(h.get(i), h.get(parent)) {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}
}

func (h *Heap) siftDown() {
	for i, child := 0, 1; i < h.Length() && i<<1+1 < h.Length(); i = child {
		child = i<<1 + 1

		if child+1 <= h.Length()-1 && h.Less(h.get(child+1), h.get(child)) {
			child++
		}

		if h.Less(h.get(i), h.get(child)) {
			break
		}

		h.data[i], h.data[child] = h.data[child], h.data[i]
	}
}

func (h *Heap) get(index int) *Node {
	return h.data[index]
}
