package matrix

import (
	"errors"
)

type Matrix struct {
	rows     int
	cols     int
	elements []float64
	// Offset between rows
	step int
}

func New() *Matrix {
	return NewWithParams([]float64{}, 0, 0)
}

func NewWithParams(elements []float64, rows, cols int) *Matrix {
	m := new(Matrix)
	m.rows = rows
	m.cols = cols
	m.step = cols
	m.elements = elements

	return m
}

func (m *Matrix) RowsCount() int {
	return m.rows
}

func (m *Matrix) ColsCount() int {
	return m.cols
}

func (m *Matrix) GetElements() []float64 {
	return m.elements
}

func (m *Matrix) GetElement(i, j int) float64 {
	return m.elements[i*m.step+j]
}

func (m *Matrix) SetElement(i, j int, v float64) {
	m.elements[i*m.step+j] = v
}

func (m *Matrix) Copy() *Matrix {
	copyMatrix := new(Matrix)
	copyMatrix.rows = m.rows
	copyMatrix.cols = m.cols
	copyMatrix.step = m.step
	copyMatrix.elements = make([]float64, m.rows*m.cols)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			copyMatrix.SetElement(i, j, m.GetElement(i, j))
		}
	}
	return copyMatrix
}

func (m *Matrix) Trace() (float64, error) {
	if m.rows != m.cols {
		return -1, errors.New("Not a square matrix.")
	}
	var result float64 = 0.0
	for i := 0; i < m.cols; i++ {
		result += m.GetElement(i, i)
	}
	return result, nil
}

func (m *Matrix) Diagonal() ([]float64, error) {
	if m.rows != m.cols {
		return nil, errors.New("Not a square matrix")
	}

	d := make([]float64, m.cols)
	for i := 0; i < m.cols; i++ {
		d[i] = m.GetElement(i, i)
	}
	return d, nil
}

// Add, Subtract, Scale, Multiply
