package convert_test

import (
	"fmt"
	"time"

	. "github.com/koanhealth/gotools/convert"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conversion", func() {

	Context("String", func() {

		var (
			now      = time.Now()
			testData = map[interface{}]string{
				nil:         "",
				1:           "1",
				23.6789:     "23.6789",
				true:        "true",
				false:       "false",
				"My String": "My String",
				now:         now.String(),
			}
		)

		It("Converts correctly", func() {
			for val, expectedVal := range testData {
				convertedVal := ToString(val)
				fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
				Expect(convertedVal).To(Equal(expectedVal))
			}
		})

		It("Converts byte array correctly", func() {
			val := []byte("this is a test")
			expectedVal := "this is a test"
			convertedVal := ToString(val)
			fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
			Expect(convertedVal).To(Equal(expectedVal))
		})

	})

	Context("Boolean", func() {

		var (
			now      = time.Now()
			testData = map[interface{}]bool{
				nil:         false,
				uint64(1):   true,
				0:           false,
				23.6789:     true,
				0.0:         false,
				true:        true,
				false:       false,
				"My String": false,
				"true":      true,
				"1":         true,
				"T":         true,
				"false":     false,
				"F":         false,
				"0":         false,
				now:         false,
			}
		)

		It("Converts correctly", func() {
			for val, expectedVal := range testData {
				convertedVal := ToBool(val)
				fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
				Expect(convertedVal).To(Equal(expectedVal))
			}
		})

		It("Converts byte array correctly", func() {
			val := []byte("true")
			expectedVal := true
			convertedVal := ToBool(val)
			fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
			Expect(convertedVal).To(Equal(expectedVal))
		})

	})

	Context("Float", func() {

		var (
			now      = time.Now()
			testData = map[interface{}]float64{
				nil:         0.0,
				uint32(5):   5.0,
				0:           0.0,
				23.6789:     23.6789,
				0.0:         0.0,
				true:        1.0,
				false:       0.0,
				"My String": 0.0,
				"11.2345":   11.2345,
				"-1":        -1.0,
				now:         0.0,
			}
		)

		It("Converts correctly", func() {
			for val, expectedVal := range testData {
				convertedVal := ToFloat(val)
				fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
				Expect(convertedVal).To(Equal(expectedVal))
			}
		})

		It("Converts byte array correctly", func() {
			val := []byte("45.67")
			expectedVal := float64(45.67)
			convertedVal := ToFloat(val)
			fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
			Expect(convertedVal).To(Equal(expectedVal))
		})

	})

	Context("Integer", func() {

		var (
			now      = time.Now()
			testData = map[interface{}]int64{
				nil:         0,
				uint32(5):   5,
				0:           0,
				23.6789:     23,
				0.0:         0,
				true:        1,
				false:       0,
				"My String": 0,
				"11":        11,
				"12.2345":   12,
				"-567":      -567,
				now:         0,
			}
		)

		It("Converts correctly", func() {
			for val, expectedVal := range testData {
				convertedVal := ToInt(val)
				fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
				Expect(convertedVal).To(Equal(expectedVal))
			}
		})

		It("Converts byte array correctly", func() {
			val := []byte("45.67")
			expectedVal := int64(45)
			convertedVal := ToInt(val)
			fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
			Expect(convertedVal).To(Equal(expectedVal))
		})

	})

	Context("Time", func() {

		var (
			zeroTime = time.Time{}
			now      = time.Now()
			testData = map[interface{}]time.Time{
				now:              now,
				"2012-06-15":     time.Date(2012, time.June, 15, 0, 0, 0, 0, time.UTC),
				nil:              zeroTime,
				0:                zeroTime,
				1.5:              zeroTime,
				"Garbage String": zeroTime,
				false:            zeroTime,
				true:             zeroTime,
			}
		)

		It("Converts correctly", func() {
			for val, expectedVal := range testData {
				convertedVal := ToTime(val)
				fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
				Expect(convertedVal).To(Equal(expectedVal))
			}
		})

		It("Converts byte array correctly", func() {
			val := []byte("2012-06-15")
			expectedVal := time.Date(2012, time.June, 15, 0, 0, 0, 0, time.UTC)
			convertedVal := ToTime(val)
			fmt.Fprintf(GinkgoWriter, "\n%T Value:%v Converted: %vExpected:%v", val, val, convertedVal, expectedVal)
			Expect(convertedVal).To(Equal(expectedVal))
		})

	})

})
