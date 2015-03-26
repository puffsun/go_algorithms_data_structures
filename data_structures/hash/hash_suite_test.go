package hash_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHash(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hash Suite")
}
