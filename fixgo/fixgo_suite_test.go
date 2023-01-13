package fixgo

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fixgo Suite")
}
