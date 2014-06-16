package strings_test

import (
	. "github.com/koanhealth/gotools/strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringBuilder", func() {

	Describe("Behaves like fmt", func() {

		It("Print", func() {
			sb := NewStringBuilder()
			sb.Print("This is a ", "test")
			Expect(sb.String()).To(Equal("This is a test"))
		})

		It("Printf", func() {
			sb := NewStringBuilder()
			sb.Printf("This %d is a %s", 21, "number")
			Expect(sb.String()).To(Equal("This 21 is a number"))
		})

		It("Println", func() {
			sb := NewStringBuilder()
			sb.Println("This", "is", "a", "test")
			Expect(sb.String()).To(Equal("This is a test\n"))
		})

	})

})
