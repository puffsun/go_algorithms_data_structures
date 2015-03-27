package sorting_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/algorithms/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BucketSort", func() {
	var (
		slice       []int
		sortedSlice []int
		emptySlice  []int
	)

	BeforeEach(func() {
		slice = []int{3, 3, 2, 1, 1, 2, 3}
		sortedSlice = []int{1, 1, 2, 2, 3, 3, 3}
		emptySlice = []int{}
	})

	Describe("Testing bucket sorting algorithms", func() {
		Context("Sort an empty slice", func() {
			It("should return immediately", func() {
				BucketSort(emptySlice, 0)
				Expect(emptySlice).To(Equal([]int{}))
			})
		})

		Context("Sort a list with int data", func() {
			It("should sort the slice", func() {
				BucketSort(slice, 3)
				Expect(slice).To(Equal(sortedSlice))
			})
		})
	})
})
