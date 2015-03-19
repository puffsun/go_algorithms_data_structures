package queue_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/queue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Priority queue", func() {
	var (
		pq    *PQ
		maxPQ *PQ
		minPQ *PQ
	)

	BeforeEach(func() {
		pq = NewPQ()
		maxPQ = NewMax()
		minPQ = NewMin()
	})

	Describe("Create new PQ", func() {
		Context("With new PQ", func() {
			It("should have 0 length", func() {
				Expect(pq.Length()).To(Equal(0))
				Expect(minPQ.Length()).To(Equal(0))
				Expect(maxPQ.Length()).To(Equal(0))
			})
		})

		Context("Inserting data into PQ", func() {
			It("should order the elements by priority", func() {
				pq.Insert(1, 10)
				Expect(pq.Length()).To(Equal(1))
				value, _ := pq.Extract()
				Expect(pq.Length()).To(Equal(0))
				Expect(value).To(Equal(1))

				pq.Insert(9, 10)
				pq.Insert(6, 20)
				pq.Insert(3, 30)
				Expect(pq.Length()).To(Equal(3))
				value, _ = pq.Extract()
				Expect(value).To(Equal(9))
				pq.Extract()
				value, _ = pq.Extract()
				Expect(value).To(Equal(3))
			})
		})
	})
})
