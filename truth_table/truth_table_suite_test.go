package truth_table

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTruthTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TruthTable Suite")
}
