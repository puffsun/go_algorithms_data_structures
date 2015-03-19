package searching_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/algorithms/searching"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinarySearch", func() {
	var (
		sortedSlice []int
		emptySlice  []int
	)

	BeforeEach(func() {
		sortedSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		emptySlice = []int{}
	})

	Describe("Searching a sorted slice", func() {
		Context("Performing binary search", func() {
			It("should return index of the element", func() {
				value := BinarySearch(sortedSlice, 3)
				Expect(value).To(Equal(2))
			})
		})
	})

	Describe("Searching a empty slice", func() {
		Context("Performing binary search", func() {
			It("should return error instead", func() {
				value := BinarySearch(emptySlice, 1)
				Expect(value).To(Equal(-1))
			})
		})
	})

	Describe("Searching a element not exist", func() {
		Context("Performing binary search", func() {
			It("should return error instead", func() {
				value := BinarySearch(sortedSlice, 20)
				Expect(value).To(Equal(-1))
			})
		})
	})
})
