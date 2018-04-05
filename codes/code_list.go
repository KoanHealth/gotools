package codes

import (
	"errors"
	"fmt"
	"strings"

	"github.com/koanhealth/gotools/slices"
	"regexp"
	"sort"
)

var (
	ErrBlankCode         = errors.New("Code cannot be blank")
	ErrInvalidCodeType   = errors.New("Invalid code type")
	ErrInvalidCodeRange  = errors.New("Beginning and End of Code Range must have the same length")
	ErrMalformedCodeList = errors.New("Malformed code list")
)

type CodeList struct {
	codes       map[string]bool
	codeRanges  []codeRange
	strictMatch bool
}

func (cc *CodeList) String() string {
	keys := make([]string, 0, len(cc.codes)+len(cc.codeRanges))
	for key := range cc.codes {
		keys = append(keys, key)
	}

	for _, cr := range cc.codeRanges {
		keys = append(keys, cr.begin+".."+cr.end)
	}

	sort.Strings(keys)

	return strings.Join(keys, ",")
}

func CompactCodes(minimumRangeLength int, codeStrings ...string) (result string, err error) {
	input := slices.StringSlice(codeStrings)
	hasCodeRanges := input.Any(func(element string) bool {
		return strings.Contains(element, "..")
	})
	for index, in := range input {
		input[index] = strings.TrimSpace(strings.ToUpper(in))
	}

	if hasCodeRanges {
		err = fmt.Errorf("Input contains code range element ('..')")
		return
	}
	input.Sort()

	output := make([]string, 0)
	for index := 0; index < len(input); {

		rangeEndIndex := index + 1
		rangeExpression := ""
		for ; rangeEndIndex < len(input); rangeEndIndex += 1 {
			if exp, detected := detectRange(input[index : rangeEndIndex+1]); detected {
				rangeExpression = exp
			} else {
				break
			}
		}

		if rangeExpression != "" && (rangeEndIndex-index) >= minimumRangeLength {
			output = append(output, rangeExpression)
			index = rangeEndIndex
		} else {
			output = append(output, input[index])
			index += 1
		}

	}

	return strings.Join(output, ","), nil
}

func detectRange(codes []string) (rangeExpression string, detected bool) {
	rangeExpression = ""
	detected = false

	switch len(codes) {
	case 0:
		return
	case 1:
		rangeExpression = codes[0]
		detected = true
		return
	}

	current := codes[0]
	length := len(current)

	for index, code := range codes {
		if index == 0 {
			continue
		}

		if len(code) != length {
			return
		}

		if code == IncrementString(current) {
			current = code
		} else {
			return
		}
	}
	rangeExpression = fmt.Sprintf("%s..%s", codes[0], codes[len(codes)-1])
	detected = true

	return
}

func IncrementString(input string) string {
	runes := slices.StringSlice(strings.Split(input, "")).Reverse()
	increment := true
	failed := false
	for index, rune := range runes {
		if failed || !increment {
			continue
		}

		switch {
		case rune == "9":
			runes[index] = "0"
		case rune == "Z":
			runes[index] = "A"
		case rune >= "0" && rune < "9":
			fallthrough
		case rune >= "A" && rune <= "Y":

			increment = false
			v := rune[0]
			v += 1
			runes[index] = string(v)
		default:
			failed = true
		}
	}

	if increment || failed {
		return ""
	} else {
		return strings.Join(runes.Reverse(), "")
	}
}

func ParseCodeList(codeList string) *CodeList {
	cl, err := TryParseCodeList(codeList)
	if err != nil {
		panic(err)
	}
	return cl
}

func TryParseCodeList(codeList string) (*CodeList, error) {
	codeList = strings.TrimSpace(strings.ToUpper(codeList))
	if codeList == "" {
		return nil, ErrBlankCode
	}

	individualCodes := make(map[string]bool)
	codeRanges := make([]codeRange, 0, 5)

	splitter := regexp.MustCompile("[A-Za-z0-9\\.]+")

	for _, code := range splitter.FindAllString(codeList, -1) {
		rangeBounds := strings.Split(code, "..")
		if len(rangeBounds) == 1 {
			strippedCode := strings.TrimSpace(code)
			if len(strippedCode) > 0 {
				individualCodes[strippedCode] = true
			}
		} else if len(rangeBounds) == 2 {
			strippedCode1 := strings.TrimSpace(rangeBounds[0])
			strippedCode2 := strings.TrimSpace(rangeBounds[1])

			if len(strippedCode1) > 0 && len(strippedCode2) > 0 {
				newRange, err := newCodeRange(strippedCode1, strippedCode2)
				if err != nil {
					return nil, err
				}
				codeRanges = append(codeRanges, newRange)
			} else {
				return nil, ErrMalformedCodeList
			}
		} else {
			return nil, ErrMalformedCodeList
		}
	}
	return &CodeList{codes: individualCodes, codeRanges: codeRanges}, nil
}

func (c *CodeList) WithStrictMatching() *CodeList {
	c.strictMatch = true
	return c
}

func (c *CodeList) Includes(code string) bool {
	code = strings.ToUpper(code)
	_, present := c.codes[code]
	if present {
		return true
	} else {
		for _, codeRange := range c.codeRanges {
			if codeRange.contains(code, c.strictMatch) {
				return true
			}
		}
	}

	return false
}

func (c *CodeList) IncludesAny(codeList ...string) bool {
	for _, code := range codeList {
		if c.Includes(code) {
			return true
		}
	}
	return false
}

type codeRange struct {
	begin string
	end   string
}

func newCodeRange(begin string, end string) (codeRange, error) {
	if len(begin) != len(end) {
		return codeRange{}, ErrInvalidCodeRange
	} else {
		return codeRange{begin: begin, end: end}, nil
	}

}

func (cr *codeRange) contains(code string, strictMatch bool) bool {
	return (!strictMatch || len(code) == len(cr.begin)) && (code >= cr.begin && code <= cr.end)
}
