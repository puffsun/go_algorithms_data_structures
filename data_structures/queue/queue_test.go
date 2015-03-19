package queue_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/queue"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queue", func() {
	var (
		queue *Queue
	)

	BeforeEach(func() {
		queue = New()
	})

	Describe("New queue", func() {
		Context("Testing a empty queue", func() {
			It("should be empty", func() {
				Expect(queue.Length()).To(Equal(0))
				Expect(queue.IsEmpty()).To(BeTrue())
			})
		})
	})

	Describe("Enqueue and Dequeue", func() {
		Context("Adding new elements into Queue", func() {
			It("should contains the newly added elements", func() {
				queue.Enqueue("One")
				Expect(queue.Length()).To(Equal(1))

				value, err := queue.Dequeue()
				Expect(value).To(Equal("One"))
				Expect(err).To(BeNil())
				Expect(queue.Length()).To(Equal(0))

				value, err = queue.Dequeue()
				Expect(value).To(BeNil())
				Expect(err).NotTo(BeNil())

				queue.Enqueue("Two")
				queue.Enqueue("Three")
				queue.Enqueue("Four")
				Expect(queue.Length()).To(Equal(3))
				value, err = queue.Dequeue()
				Expect(value).To(Equal("Two"))
				Expect(queue.Length()).To(Equal(2))
			})
		})
	})
})
