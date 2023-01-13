package fixgo

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ternary", func() {
	It("Ternary operates a into", func() {
		Expect(Ternary(true, 1, 2)).To(Equal(1))
	})
	It("Ternary operates a string", func() {
		Expect(Ternary(true, "String1", "String2")).To(Equal("String1"))
	})
	It("Ternary does ternary things", func() {
		Expect(Ternary(false, "String1", "String2")).To(Equal("String2"))
	})

})
