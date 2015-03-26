package tree_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/tree"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bst", func() {
	var (
		bst *BST
	)

	BeforeEach(func() {
		bst = New()
	})

	Describe("Testing empty BST", func() {
		Context("With newly created BST", func() {
			It("should be empty", func() {
				Expect(bst.Size()).To(Equal(0))
				Expect(bst.IsEmpty()).To(BeTrue())
				Expect(bst.IsBST()).To(BeTrue())
			})

			It("should not contains any element", func() {
				Expect(bst.Contains(12)).To(BeFalse())
				Expect(bst.Min()).To(BeNil())
				err := bst.DeleteMin()
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("Testing BST with data", func() {
		Context("Populate BST with data", func() {
			It("should contains the populated data", func() {
				bst.Put(10, "Ten")
				Expect(bst.Get(10)).To(Equal("Ten"))
				Expect(bst.Get(100)).To(BeNil())
				Expect(bst.Size()).To(Equal(1))

				bst.Put(11, "Eleven")
				bst.Put(1, "One")
				bst.Put(6, "Six")
				Expect(bst.Get(1)).To(Equal("One"))
				Expect(bst.Get(11)).To(Equal("Eleven"))
				Expect(bst.Get(6)).To(Equal("Six"))
				Expect(bst.Size()).To(Equal(4))

				bst.Delete(10)
				Expect(bst.Get(10)).To(BeNil())
				Expect(bst.Size()).To(Equal(3))

				// populate nil data is same as delete the node
				bst.Put(11, nil)
				Expect(bst.Size()).To(Equal(2))
				Expect(bst.Get(11)).To(BeNil())
			})
		})

		Context("Testing min/deleteMin/max/deleteMax", func() {
			It("should return min node", func() {
				bst.Put(1, "one")
				bst.Put(2, "two")
				bst.Put(3, "three")
				bst.Put(4, "four")
				bst.Put(5, "five")
				Expect(bst.IsBST()).To(BeTrue())
				Expect(bst.Size()).To(Equal(5))
				// remember, type Key int in node.go
				Expect(bst.Min()).To(Equal(Key(1)))
				Expect(bst.Max()).To(Equal(Key(5)))

				bst.DeleteMin()
				Expect(bst.Size()).To(Equal(4))
				Expect(bst.Get(1)).To(BeNil())

				bst.DeleteMax()
				Expect(bst.Size()).To(Equal(3))
				Expect(bst.Get(5)).To(BeNil())
			})
		})
	})
})
