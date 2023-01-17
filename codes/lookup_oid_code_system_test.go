package codes

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CodeList", func() {

	It("unknown", func() {
		Expect(LookupOidCodeSystem("foo")).To(Equal(CODE_SYSTEM_UNKNOWN))
	})
	It("CPT", func() {
		Expect(LookupOidCodeSystem("2.16.840.1.113883.6.12")).To(Equal(CODE_SYSTEM_CPT))
	})
	It("LOINC", func() {
		Expect(LookupOidCodeSystem("2.16.840.1.113883.6.1")).To(Equal(CODE_SYSTEM_LOINC))
	})
})
