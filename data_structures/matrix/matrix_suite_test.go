package matrix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMatrix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Matrix Suite")
}
