// http://en.wikipedia.org/wiki/Hash_table#Separate_chaining_with_linked_lists
package hash

import (
	"errors"
	"github.com/puffsun/go_algorithms_data_structures/data_structures/list"
	"math"
)

type Hash struct {
	table    map[int]*list.LinkedList
	size     int
	capacity int
}

type entry struct {
	key   string
	value interface{}
}

func New(capacity int) *Hash {
	h := new(Hash)
	h.table = make(map[int]*list.LinkedList, capacity)
	h.size = 0
	h.capacity = capacity
	return h
}

func (h *Hash) Size() int {
	return h.size
}

func (h *Hash) IsEmpty() bool {
	return h.size == 0
}

func (h *Hash) Capacity() int {
	return h.capacity
}

func (h *Hash) Get(key string) (interface{}, error) {
	index := h.position(key)
	result, err := h.find(key, index)
	return result.value, err
}

func (h *Hash) Put(key string, value interface{}) {
	index := h.position(key)
	if h.table[index] == nil {
		h.table[index] = list.New()
	}

	e := &entry{key: key, value: value}
	val, err := h.find(key, index)
	if err != nil {
		h.table[index].Append(e)
		h.size += 1
	} else {
		// Update exist entry
		val.value = value
	}
}

func (h *Hash) Remove(key string) error {
	index := h.position(key)
	if index < 0 || index > h.size {
		return errors.New("key not found")
	}

	lst := h.table[index]
	var val *entry

	lst.Each(func(node *list.Node) {
		if node.Value().(*entry).key == key {
			val = node.Value().(*entry)
		}
	})

	if val == nil {
		return nil
	} else {
		h.size -= 1
		return lst.Remove(val)
	}
}

func (h *Hash) find(key string, index int) (*entry, error) {
	lst := h.table[index]
	var val *entry

	lst.Each(func(node *list.Node) {
		if node.Value().(*entry).key == key {
			val = node.Value().((*entry))
		}
	})

	if val == nil {
		return nil, errors.New("Element not found")
	}
	return val, nil
}

// Calculate element index in the hash
func (h *Hash) position(str string) int {
	return hashCode(str) % h.Capacity()
}

func hashCode(str string) int {
	hash := int32(0)
	for i := 0; i < len(str); i++ {
		hash = int32(hash<<5-hash) + int32(str[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}
