package strings_test

import (
	"strings"

	. "github.com/koanhealth/gotools/strings"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String Predicates", func() {
	var ()
	It("IsBlank should work", func() {
		Expect(IsBlank("")).To(BeTrue())
		Expect(IsBlank(" ")).To(BeTrue())
		Expect(IsBlank("Foo")).To(BeFalse())
		Expect(IsBlank("Foo")).To(BeFalse())
		Expect(IsBlank(" Foo")).To(BeFalse())
		Expect(len(strings.TrimSpace(" Foo"))).To(Equal(3))
	})
	var (
		True  = StringPredicate(func(s string) bool { return true })
		False = StringPredicate(func(s string) bool { return false })
	)

	It("And", func() {
		Expect(False.And(False)("")).To(BeFalse())
		Expect(True.And(False)("")).To(BeFalse())
		Expect(True.And(True)("")).To(BeTrue())
		Expect(False.And(True)("")).To(BeFalse())
	})

	It("Or", func() {
		Expect(False.Or(False)("")).To(BeFalse())
		Expect(True.Or(False)("")).To(BeTrue())
		Expect(True.Or(True)("")).To(BeTrue())
		Expect(False.Or(True)("")).To(BeTrue())
	})

	It("Not", func() {
		Expect(False.Not()("")).To(BeTrue())
		Expect(True.Not()("")).To(BeFalse())
	})
})
