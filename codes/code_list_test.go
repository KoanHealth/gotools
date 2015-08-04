package codes_test

import (
	. "github.com/koanhealth/gotools/codes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CodeList", func() {

	Context("Bad Parse", func() {
		It("parse with empty code returns error", func() {
			_, err := TryParseCodeList("")
			Expect(err).To(Equal(ErrBlankCode))
		})
		It("malformed range produces error", func() {
			_, err := TryParseCodeList("V90..V98..V99")
			Expect(err).To(Equal(ErrMalformedCodeList))
		})
		It("mismatched range bounds produce error", func() {
			_, err := TryParseCodeList("V90..V98.1")
			Expect(err).To(Equal(ErrInvalidCodeRange))
		})

		It("parse with empty code returns error", func() {
			Expect(func() {
				ParseCodeList("")
			}).To(Panic())
		})
		It("malformed range produces error", func() {
			Expect(func() {
				ParseCodeList("V90..V98..V99")
			}).To(Panic())
		})
		It("mismatched range bounds produce error", func() {
			Expect(func() {
				ParseCodeList("V90..V98.1")
			}).To(Panic())
		})
	})

	Context("Good Parse", func() {
		It("Parse single code", func() {
			c := ParseCodeList("code1")
			Expect(c.Includes("code1")).To(BeTrue())
		})

		It("Parses a single set of codes", func() {
			c := ParseCodeList("A001, A002,A003")
			Expect(c.Includes("A001")).To(BeTrue())
			Expect(c.Includes("A002")).To(BeTrue())
			Expect(c.Includes("A003")).To(BeTrue())
			Expect(c.Includes("A004")).To(BeFalse())
		})

		It("Parses a single code range", func() {
			c := ParseCodeList("100..110")
			Expect(c.Includes("100")).To(BeTrue())
			Expect(c.Includes("105")).To(BeTrue())
			Expect(c.Includes("110")).To(BeTrue())
		})

		It("parse multiple codes", func() {
			c := ParseCodeList("code1, code2, code3")
			Expect(c.Includes("code1")).To(BeTrue())
			Expect(c.Includes("code2")).To(BeTrue())
			Expect(c.Includes("code3")).To(BeTrue())
			Expect(c.Includes("code44")).To(BeFalse())
		})

		It("parse multiple codes with space separators", func() {
			c := ParseCodeList("code1 code2 code3 A101..A103")
			Expect(c.Includes("code1")).To(BeTrue())
			Expect(c.Includes("code2")).To(BeTrue())
			Expect(c.Includes("code3")).To(BeTrue())
			Expect(c.Includes("A101")).To(BeTrue())
			Expect(c.Includes("A102")).To(BeTrue())
			Expect(c.Includes("A103")).To(BeTrue())
			Expect(c.Includes("code44")).To(BeFalse())
		})

		It("parse complex codeset", func() {
			c := ParseCodeList("code12..code20 , ,  code3,code99,V90..V99")
			Expect(c.Includes("code12")).To(BeTrue())
			Expect(c.Includes("code125")).To(BeTrue())
			Expect(c.Includes("code17")).To(BeTrue())
			Expect(c.Includes("code20")).To(BeTrue())
			Expect(c.Includes("code3")).To(BeTrue())
			Expect(c.Includes("code99")).To(BeTrue())
			Expect(c.Includes("V90")).To(BeTrue())
			Expect(c.Includes("V95")).To(BeTrue())
			Expect(c.Includes("V99")).To(BeTrue())
		})

		It("parses stupidly long codeset", func() {
			c := ParseCodeList(`
D1..D5 D7..D9 D54 D55 D60 D61 D63 D76..D78 D82..D85 D87 D89..D93 D99 D100 D102 D104 D107 D109 D112
D116 D118 D120 D122..D131 D135 D137 D139 D140 D142 D145 D146 D148 D153 D154 D157 D159 D165 D168
D172 D197 D198 D225..D230 D232..D235 D237..D247 D249..D253 D259 D650..D653 D656 D658 D660..D663 D670`).WithStrictMatching()
			Expect(c.Includes("D8")).To(BeTrue())
		})
	})

	Context("Stict Matching", func() {
		It("Does not match interstitial codes when matching is strict", func() {
			c := ParseCodeList("code12..code20").WithStrictMatching()
			Expect(c.Includes("code12")).To(BeTrue())
			Expect(c.Includes("code17")).To(BeTrue())

			// Because of Strict Matching, the code list will not interpolate to eaxtra characters in the range
			Expect(c.Includes("code125")).To(BeFalse())

		})
	})

	Context("Other Matching", func() {
		It("IncludesAny returns true if any code exists", func() {
			c := ParseCodeList("code1, code2, code3")
			Expect(c.IncludesAny("code4", "code5", "code2")).To(BeTrue())
		})

		It("IncludesAny returns false if none of the codes exist", func() {
			c := ParseCodeList("code1, code2, code3")
			Expect(c.IncludesAny("code4", "code5", "code6")).To(BeFalse())
		})

		It("Matches codes in ranges that have extra stuff", func() {
			c := ParseCodeList("V90..V99")
			Expect(c.Includes("V90.00")).To(BeTrue())
			Expect(c.Includes("V90.1")).To(BeTrue())
		})
	})
})
