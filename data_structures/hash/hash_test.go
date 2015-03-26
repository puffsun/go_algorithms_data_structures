package hash_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/hash"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hash", func() {
	var (
		hash         *Hash
		hashCapacity int
	)

	BeforeEach(func() {
		hashCapacity = 10
		hash = New(hashCapacity)
	})

	Describe("Testing new hash", func() {
		Context("Create an empty hash", func() {
			It("should be empty", func() {
				Expect(hash.Size()).To(Equal(0))
				Expect(hash.IsEmpty()).To(BeTrue())
				Expect(hash.Capacity()).To(Equal(hashCapacity))
			})
		})
	})

	Describe("Test hash with data", func() {
		Context("Populating hash with data", func() {
			It("should contains the populated data", func() {
				hash.Put("One", 1)
				Expect(hash.IsEmpty()).To(BeFalse())
				Expect(hash.Size()).To(Equal(1))

				hash.Put("Two", 2)
				Expect(hash.Size()).To(Equal(2))

				val, err := hash.Get("One")
				Expect(val).To(Equal(1))
				Expect(err).To(BeNil())

				err = hash.Remove("One")
				Expect(hash.Size()).To(Equal(1))
				Expect(err).To(BeNil())

				err = hash.Remove("Ten")
				Expect(hash.Size()).To(Equal(1))
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
