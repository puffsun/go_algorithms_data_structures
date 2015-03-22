package sorting_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/algorithms/sorting"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShellSort", func() {

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
		Context("Performing shell sort", func() {
			It("should return immediately", func() {
				ShellSort(emptySlice)
				Expect(emptySlice).To(Equal([]int{}))
			})
		})
	})

	Describe("Sorting a unsorted slice", func() {
		Context("Performing shell sort", func() {
			It("should sort the slice", func() {
				ShellSort(slice)
				Expect(slice).To(Equal(sortedSlice))
			})
		})
	})
})
