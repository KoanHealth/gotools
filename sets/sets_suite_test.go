package sets_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sets Suite")
}
