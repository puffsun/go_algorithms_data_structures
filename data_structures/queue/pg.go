package queue

import (
	"github.com/puffsun/go_algorithms_data_structures/data_structures/heap"
)

type PQ struct {
	data *heap.Heap
}

func NewPQ() *PQ {
	return &PQ{
		data: heap.New(),
	}
}

func NewMax() *PQ {
	return &PQ{
		data: heap.NewMax(),
	}
}

func NewMin() *PQ {
	return &PQ{
		data: heap.NewMin(),
	}
}

func (pq *PQ) Length() int {
	return pq.data.Length()
}

func (pq *PQ) Insert(value, priority int) {
	pq.data.InsertWithPriority(value, priority)
}

func (pq *PQ) Extract() (int, error) {
	return pq.data.Extract()
}
