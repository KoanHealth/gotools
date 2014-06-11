package slices_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSlices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slices Suite")
}
