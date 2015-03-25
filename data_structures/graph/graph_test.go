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

			It("should have different directed attribute", func() {
				Expect(dg.IsDirected()).To(BeTrue())
				Expect(ug.IsDirected()).To(BeFalse())
			})
		})
	})

	Describe("Testing un-directed graph", func() {
		Context("Adding edges and Vertexes", func() {
			It("should contains the added edges and vertexes", func() {
				for i := 0; i < 10; i++ {
					v := VertexId(i)
					ug.AddVertex(v)
				}
				Expect(ug.VerticesCount()).To(Equal(10))
				for i := 0; i < 10; i++ {
					Expect(ug.IsVertex(VertexId(i))).To(BeTrue())
				}
				Expect(ug.IsVertex(VertexId(10))).To(BeFalse())

				for i := 0; i < 10; i++ {
					ug.AddEdge(VertexId(i), VertexId(i%2), 1)
				}
				// Cannot add edge the the same from and to vertex
				// in this case, it is 0 and 1
				Expect(ug.EdgesCount()).To(Equal(8))
				Expect(ug.IsEdge(0, 8)).To(BeTrue())
				Expect(ug.IsEdge(8, 0)).To(BeTrue())

				Expect(ug.IsEdge(9, 0)).To(BeFalse())
				Expect(ug.IsEdge(0, 9)).To(BeFalse())

				Expect(ug.CheckVertex(2)).To(BeTrue())

				// Test removing edges and vertexes
				err := ug.RemoveEdge(0, 2)
				Expect(err).To(BeNil())
				Expect(ug.EdgesCount()).To(Equal(7))

				err = ug.RemoveEdge(0, 9)
				Expect(err).NotTo(BeNil())
				Expect(ug.EdgesCount()).To(Equal(7))

				err = ug.RemoveVertex(10)
				Expect(err).NotTo(BeNil())

				err = ug.RemoveVertex(9)
				Expect(err).To(BeNil())
				Expect(ug.VerticesCount()).To(Equal(9))

				for i := 0; i < 9; i++ {
					ug.RemoveVertex(VertexId(i))
				}
				Expect(ug.VerticesCount()).To(Equal(0))
			})

			// Continue with above context
			It("should return error to add existing edges", func() {
				err := ug.AddEdge(0, 2, 1)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
