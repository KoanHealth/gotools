package codes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCodes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Codes Suite")
}
