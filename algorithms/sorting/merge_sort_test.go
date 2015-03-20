package sorting_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/algorithms/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MergeSort", func() {
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

	Describe("Sorting on an empty slice", func() {
		Context("Performing merge sort", func() {
			It("should return silently", func() {
				MergeSort(emptySlice)
				Expect(emptySlice).To(Equal([]int{}))
			})
		})
	})

	Describe("Sorting a unordered slice", func() {
		Context("Performing merge sort", func() {
			It("should sort the slice", func() {
				MergeSort(slice)
				Expect(slice).To(Equal(sortedSlice))
			})
		})
	})
})
