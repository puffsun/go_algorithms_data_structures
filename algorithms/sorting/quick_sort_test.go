package sorting_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/algorithms/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuickSort", func() {
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

	Describe("Sorting a empty slice", func() {
		Context("Performing quick sort", func() {
			It("should return silently", func() {
				QuickSortWithAnonymousFunc(emptySlice)
				Expect(emptySlice).To(Equal([]int{}))
			})
		})
	})

	Describe("Sorting a slice", func() {
		Context("Performing quick sort", func() {
			It("should sort the slice", func() {
				QuickSortWithAnonymousFunc(slice)
				Expect(slice).To(Equal(sortedSlice))
			})
		})
	})

	Describe("Sorting a empty slice", func() {
		Context("Performing quick sort", func() {
			It("should return silently", func() {
				QuickSort(emptySlice)
				Expect(emptySlice).To(Equal([]int{}))
			})
		})
	})

	Describe("Sorting a slice", func() {
		Context("Performing quick sort", func() {
			It("should sort the slice", func() {
				QuickSort(slice)
				Expect(slice).To(Equal(sortedSlice))
			})
		})
	})
})
