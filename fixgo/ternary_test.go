package fixgo

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type RandomType struct {
	Value int
}

var _ = Describe("Ternary", func() {
	It("Ternary operates an int", func() {
		Expect(Ternary(true, 1, 2)).To(Equal(1))
	})
	It("Ternary operates a string", func() {
		Expect(Ternary(true, "String1", "String2")).To(Equal("String1"))
	})
	It("Ternary does ternary things", func() {
		Expect(Ternary(false, "String1", "String2")).To(Equal("String2"))
	})

	It("Ternary does ternary things to random types", func() {
		var result RandomType
		result = Ternary(true, RandomType{Value: 1}, RandomType{Value: 2})
		Expect(result).To(Equal(RandomType{Value: 1}))
	})

})
