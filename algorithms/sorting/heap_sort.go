package sorting

import (
	"github.com/puffsun/go_algorithms_data_structures/data_structures/heap"
)

func HeapSort(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}
	h := heap.NewMin()

	for i := 0; i < len(slice); i += 1 {
		h.Insert(slice[i])
	}

	for i := 0; i < len(slice); i += 1 {
		value, _ := h.Extract()
		slice[i] = value
	}
}
