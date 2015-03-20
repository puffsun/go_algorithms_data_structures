package sorting_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/algorithms/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SelectionSort", func() {
	var (
		slice       []int
		sortedSlice []int
		emptySlice  []int
	)

	BeforeEach(func() {
		slice = []int{9, 6, 3, 10, 21, 2}
		sortedSlice = []int{2, 3, 6, 9, 10, 21}
		emptySlice = []int{}
	})

	Describe("Sorting on a empty slice", func() {
		Context("Performing selection sort", func() {
			It("should return silently", func() {
				SelectionSort(emptySlice)
				Expect(emptySlice).To(Equal([]int{}))
			})
		})
	})

	Describe("Sorting a unordered slice", func() {
		Context("Performing selection sort", func() {
			It("should sort the slice", func() {
				SelectionSort(slice)
				Expect(slice).To(Equal(sortedSlice))
			})
		})
	})
})
