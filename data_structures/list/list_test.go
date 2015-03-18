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
				Expect(l.RemoveFirst()).To(BeNil())
				Expect(l.RemoveLast()).To(BeNil())
			})
		})
	})

	Describe("Testing linked list with data", func() {
		Context("Prepending data", func() {
			It("should contains the same data", func() {
				list.Prepend("One")
				Expect(list.Length()).To(Equal(1))
				Expect(list.RemoveFirst()).To(Equal("One"))
				Expect(list.Length()).To(Equal(0))

				list.Prepend("One")
				list.Prepend("Two")
				Expect(list.Length()).To(Equal(2))
				Expect(list.RemoveFirst()).To(Equal("Two"))
				Expect(list.RemoveFirst()).To(Equal("One"))
			})
		})

		Context("Appending data", func() {
			It("should contains the same data", func() {
				list.Append("One")
				Expect(list.Length()).To(Equal(1))
				Expect(list.RemoveLast()).To(Equal("One"))
				Expect(list.Length()).To(Equal(0))

				list.Append("One")
				list.Append("Two")
				Expect(list.Length()).To(Equal(2))
				Expect(list.RemoveLast()).To(Equal("Two"))
				Expect(list.RemoveLast()).To(Equal("One"))
				Expect(list.Length()).To(Equal(0))

				list.Append("One")
				list.Prepend("Two")
				list.Append("Three")
				list.Prepend("Four")
				Expect(list.Length()).To(Equal(4))
				Expect(list.RemoveFirst()).To(Equal("Four"))
				Expect(list.RemoveLast()).To(Equal("Three"))
			})
		})

		Context("Clear list data", func() {
			It("should remove all data in the list", func() {
				list.Append("One")
				list.Append("Two")
				Expect(list.Length()).To(Equal(2))
				list.Clear()
				Expect(list.Length()).To(Equal(0))
			})
		})

		Context("Get list data with specified index", func() {
			It("should return data in the slot", func() {
				Expect(list.Length()).To(Equal(0))
				val, _ := list.Get(0)
				Expect(val).To(BeNil())
				list.Append("One")
				list.Append("Two")
				Expect(list.Get(0)).To(Equal("One"))
				Expect(list.Get(1)).To(Equal("Two"))
			})
		})

		Context("Adding data with invalid index", func() {
			It("should return error", func() {
				_ = list.Insert(1, "One")
				Expect(list.Length()).To(Equal(0))
			})
		})

		Context("Adding data with valie index", func() {
			It("should be added into list", func() {
				list.Insert(0, "Two")
				Expect(list.Length()).To(Equal(1))
			})
		})
	})

	// TODO Remove TOD
})
