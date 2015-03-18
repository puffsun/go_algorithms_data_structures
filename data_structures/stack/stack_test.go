package stack_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/stack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	var (
		st *Stack
	)

	BeforeEach(func() {
		st = New()
	})

	Describe("Testing new stack", func() {
		Context("Pop from empty stack", func() {
			It("should be nil", func() {
				result := st.Pop()
				Expect(result).To(BeNil())
			})
		})

		Context("Inspect a empty stack", func() {
			It("should be nil", func() {
				result := st.Peek()
				Expect(result).To(BeNil())
			})
		})

		Context("The length of a empty stack should be 0", func() {
			It("should be 0", func() {
				Expect(st.Len()).To(Equal(0))
				Expect(st.IsEmpty()).To(BeTrue())
			})
		})

		Context("Pushing a new element", func() {
			It("Length should be 1", func() {
				st.Push("One")
				Expect(st.Len()).To(Equal(1))
				result := st.Pop()
				Expect(result).To(Equal("One"))
			})
		})
	})
})
