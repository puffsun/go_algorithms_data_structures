package graph_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/graph"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Graph", func() {
	var (
		ug *Graph
		dg *Graph
	)

	BeforeEach(func() {
		ug = NewUndirected()
		dg = NewDirected()
	})

	Describe("Testing newly created graphs", func() {
		Context("Empty graphs has been created", func() {
			It("should be nil", func() {
				Expect(ug).NotTo(BeNil())
				Expect(dg).NotTo(BeNil())
			})

			It("should be empty", func() {
				Expect(ug.EdgesCount()).To(Equal(0))
				Expect(dg.EdgesCount()).To(Equal(0))
			})
		})
	})
})
