package heap_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/heap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Heap", func() {
	var (
		heap    *Heap
		minHeap *Heap
		maxHeap *Heap
	)

	BeforeEach(func() {
		heap = New()
		minHeap = NewMin()
		maxHeap = NewMax()
	})

	Describe("Newly created Heap", func() {
		Context("With new Heap", func() {
			It("should have no data", func() {
				Expect(heap.Length()).To(Equal(0))
				Expect(heap.IsEmpty()).To(BeTrue())

				Expect(minHeap.Length()).To(Equal(0))
				Expect(minHeap.IsEmpty()).To(BeTrue())

				Expect(maxHeap.Length()).To(Equal(0))
				Expect(maxHeap.IsEmpty()).To(BeTrue())
			})
		})
	})

	Describe("Inserting data", func() {
		Context("Insert data into min heap", func() {
			It("should order the data been inserted", func() {
				// this is a mean heap
				heap.Insert(10)
				heap.Insert(3)
				heap.Insert(6)
				heap.Insert(9)
				Expect(heap.Length()).To(Equal(4))
				value, err := heap.Extract()
				Expect(value).To(Equal(3))
				Expect(err).To(BeNil())
				Expect(heap.Length()).To(Equal(3))
				value, _ = heap.Extract()
				Expect(value).To(Equal(6))
				value, _ = heap.Extract()
				Expect(value).To(Equal(9))
				value, _ = heap.Extract()
				Expect(value).To(Equal(10))
				Expect(heap.Length()).To(Equal(0))
			})
		})

		Context("Insert data into max heap, then extract", func() {
			It("should order the elements", func() {
				maxHeap.Insert(3)
				maxHeap.Insert(6)
				maxHeap.Insert(9)
				maxHeap.Insert(1)
				Expect(maxHeap.Length()).To(Equal(4))

				value, err := maxHeap.Extract()
				Expect(value).To(Equal(9))
				Expect(err).To(BeNil())
				maxHeap.Extract()
				maxHeap.Extract()
				maxHeap.Extract()
				Expect(maxHeap.Length()).To(Equal(0))
			})
		})

		Context("Extract from empty heap", func() {
			It("should return error", func() {
				value, err := minHeap.Extract()
				Expect(value).To(Equal(-1))
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
