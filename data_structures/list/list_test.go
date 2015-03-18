package list_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/list"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List", func() {
	var (
		list *LinkedList
	)

	BeforeEach(func() {
		list = New()
	})

	Describe("Testing new LinkedList", func() {
		Context("Create a brand new LinkedList", func() {
			It("should be empty", func() {
				l := New()
				Expect(l.Length()).To(Equal(0))
				Expect(l.IsEmpty()).To(BeTrue())
			})
		})
	})

	Describe("Testing linked list with data", func() {
		Context("Appending data", func() {
			It("should contains the same data", func() {
				list.Prepend("One")
				Expect(list.Length()).To(Equal(1))
				Expect(list.RemoveFirst()).To(Equal("One"))
				Expect(list.Length()).To(Equal(0))
			})
		})
	})
})
