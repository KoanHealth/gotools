package codes

import (
	"errors"
	"strings"
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
	for _, code1 := range strings.Split(codeList, ",") {
		for _, code := range strings.Split(code1, " ") {
			rangeBounds := strings.Split(code, "..")
			if len(rangeBounds) == 1 {
				individualCodes[strings.TrimSpace(code)] = true
			} else if len(rangeBounds) == 2 {
				newRange, err := newCodeRange(strings.TrimSpace(rangeBounds[0]), strings.TrimSpace(rangeBounds[1]))
				if err != nil {
					return nil, err
				}
				codeRanges = append(codeRanges, newRange)
			} else {
				return nil, ErrMalformedCodeList
			}
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
