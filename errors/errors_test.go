package errors_test

import (
	"fmt"

	. "github.com/koanhealth/gotools/errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Errors", func() {

	It("Store multiple errors", func() {
		errs := ErrorSlice{
			fmt.Errorf("First error"),
			fmt.Errorf("Second error"),
			fmt.Errorf("Third error"),
		}
		err := errs.Error()
		Expect(err).To(ContainSubstring("Errors:"))
		Expect(err).To(ContainSubstring("First error"))
		Expect(err).To(ContainSubstring("Second error"))
		Expect(err).To(ContainSubstring("Third error"))
	})

})
