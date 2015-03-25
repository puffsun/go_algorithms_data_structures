package matrix_test

import (
	. "github.com/puffsun/go_algorithms_data_structures/data_structures/matrix"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix", func() {
	var (
		matrix             *Matrix
		emptyMatrix        *Matrix
		squareMatrix       *Matrix
		doubleScaledMatrix *Matrix
	)

	BeforeEach(func() {
		matrix = NewWithParams([]float64{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2)
		doubleScaledMatrix = NewWithParams([]float64{2, 4, 6, 8, 10, 12, 14, 16}, 4, 2)
		emptyMatrix = New()
		squareMatrix = NewWithParams([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 3)
	})

	Describe("Testing a empty matrix", func() {
		Context("With a newly created matrix", func() {
			It("should be empty", func() {
				Expect(emptyMatrix.RowsCount()).To(Equal(0))
				Expect(emptyMatrix.ColsCount()).To(Equal(0))
				Expect(emptyMatrix.GetElements()).To(Equal([]float64{}))
			})
		})
	})

	Describe("Testing matrix with data", func() {
		Context("Populating empty matrix with data", func() {
			It("should contains the populated data", func() {
				Expect(matrix.ColsCount()).To(Equal(2))
				Expect(matrix.RowsCount()).To(Equal(4))

				Expect(squareMatrix.RowsCount()).To(Equal(3))
				Expect(squareMatrix.ColsCount()).To(Equal(3))

				matrix.SetElement(0, 1, 10.0)
				Expect(matrix.GetElement(0, 1)).To(Equal(10.0))
				Expect(matrix.GetElement(0, 0)).To(Equal(1.0))
			})
		})
	})

	Context("Copy matrix", func() {
		It("should contain exactly the same data", func() {
			matrixCopy := matrix.Copy()
			Expect(matrixCopy.RowsCount()).To(Equal(matrix.RowsCount()))
			Expect(matrixCopy.ColsCount()).To(Equal(matrix.ColsCount()))
			Expect(matrixCopy.GetElements()).To(Equal(matrix.GetElements()))
		})
	})

	Context("Test matrix trace", func() {
		It("calculate matrix trace", func() {
			t, err := matrix.Trace()
			Expect(t).To(Equal(-1.0))
			Expect(err).NotTo(BeNil())

			t, err = squareMatrix.Trace()
			Expect(t).To(Equal(15.0))
			Expect(err).To(BeNil())
		})
	})

	Context("Test matrix diagonal", func() {
		It("should return copy of diagonal", func() {
			d, err := squareMatrix.Diagonal()
			Expect(d).To(Equal([]float64{1.0, 5.0, 9.0}))
			Expect(err).To(BeNil())

			d, err = matrix.Diagonal()
			Expect(d).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})

	Context("Test matrix scale", func() {
		It("should scale the matrix", func() {
			matrix.Scale(2.0)
			Expect(matrix.RowsCount()).To(Equal(doubleScaledMatrix.RowsCount()))
			Expect(matrix.ColsCount()).To(Equal(doubleScaledMatrix.ColsCount()))
			Expect(matrix.GetElements()).To(Equal(doubleScaledMatrix.GetElements()))
		})
	})

	Context("Test adding two matrix", func() {
		It("should report error on two matrix with different row or column size", func() {
			err := matrix.Add(squareMatrix)
			Expect(err).NotTo(BeNil())
		})

		It("should add the matrix with another matrix", func() {
			err := matrix.Add(matrix)
			Expect(err).To(BeNil())
			Expect(matrix.GetElements()).To(Equal(doubleScaledMatrix.GetElements()))
		})
	})
})
