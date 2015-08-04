package json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	ksl "github.com/koanhealth/gotools/slices"
)

type JsonFormatOptions struct {
	EliminateNulls          bool
	SingleLineItemThreshold int
	ForceMultiLineOnFields  ksl.StringSlice
}

func DefaultJsonFormatOptions() *JsonFormatOptions {
	return &JsonFormatOptions{SingleLineItemThreshold: 8}
}

func (options *JsonFormatOptions) ForceMultiLineArray(fields ...string) *JsonFormatOptions {
	for _, field := range fields {
		options.ForceMultiLineOnFields = append(options.ForceMultiLineOnFields, field)
	}
	return options
}
func (options *JsonFormatOptions) EliminateNullFields() *JsonFormatOptions {
	options.EliminateNulls = true
	return options
}

func (options *JsonFormatOptions) ForceMultiLine(fieldName string) bool {
	return options.ForceMultiLineOnFields.Contains(fieldName)
}

func FormatJsonFile(path string, options *JsonFormatOptions) (result ksl.StringSlice, err error) {

	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	result, err = FormatJsonReader(bufio.NewReader(f), options)
	return result, err
}

func FormatJsonReader(source io.Reader, options *JsonFormatOptions) (result ksl.StringSlice, err error) {
	dec := json.NewDecoder(source)
	var generic interface{}
	err = dec.Decode(&generic)
	if err != nil {
		if err != io.EOF {
			return
		}
	}

	result, err = FormatJson(generic, options), nil
	return
}

func FormatJson(source interface{}, options *JsonFormatOptions) (result ksl.StringSlice) {
	switch val := source.(type) {
	case map[string]interface{}:
		return formatJsonObject(val, "", options)
	case []interface{}:
		return formatJsonArray(val, "", options)
	default:
		return ksl.StringSlice{"Unrecognized"}
	}
}

func formatJsonValue(source interface{}, fieldName string, options *JsonFormatOptions) (result ksl.StringSlice, use bool) {
	switch val := source.(type) {
	case map[string]interface{}:
		return formatJsonObject(val, fieldName, options), true
	case []interface{}:
		return formatJsonArray(val, fieldName, options), true
	case string:
		return ksl.StringSlice{fmt.Sprintf("\"%s\"", val)}, true
	case nil:
		return ksl.StringSlice{fmt.Sprintf("null")}, !options.EliminateNulls
	default:
		return ksl.StringSlice{fmt.Sprintf("%v", val)}, true
	}
}

func formatJsonObject(source map[string]interface{}, fieldName string, options *JsonFormatOptions) (result ksl.StringSlice) {
	rawResult := ksl.StringSlice{}
	var joinElements string

	cleanSource := make(map[string]ksl.StringSlice)

	for key, rawValue := range source {
		if element, use := formatJsonValue(rawValue, key, options); use {
			cleanSource[key] = element
		}
	}

	elementsToCome := len(cleanSource)
	for key, element := range cleanSource {
		if elementsToCome--; elementsToCome > 0 {
			joinElements = ","
		} else {
			joinElements = ""
		}

		innerResult := prepareInnerResult(
			element,
			fmt.Sprintf("\"%s\": ", key),
			joinElements)
		rawResult = append(rawResult, innerResult...)
	}
	result = bakeResult(rawResult, "{}", len(rawResult) > options.SingleLineItemThreshold)
	return
}

func formatJsonArray(source []interface{}, fieldName string, options *JsonFormatOptions) (result ksl.StringSlice) {
	rawResult := ksl.StringSlice{}
	var joinElements string
	elementsToCome := len(source)

	for _, val := range source {
		if elementsToCome--; elementsToCome > 0 {
			joinElements = ","
		} else {
			joinElements = ""
		}

		innerResult, use := formatJsonValue(val, fieldName, options)
		if use {
			for index, line := range innerResult {
				if index == (len(innerResult) - 1) {
					rawResult = append(rawResult, line+joinElements)
				} else {
					rawResult = append(rawResult, line)
				}
			}
		}
	}

	result = bakeResult(rawResult, "[]", len(rawResult) > options.SingleLineItemThreshold || options.ForceMultiLine(fieldName))
	return
}

func prepareInnerResult(innerResult ksl.StringSlice, firstLinePrefix string, lastLineSuffix string) (result ksl.StringSlice) {
	for index, line := range innerResult {
		firstLine := index == 0
		lastLine := index == (len(innerResult) - 1)
		middleLine := !firstLine && !lastLine

		if firstLine {
			line = firstLinePrefix + line
		}
		if lastLine {
			line = line + lastLineSuffix
		}
		if middleLine {
			line = "\t" + line
		}
		result = append(result, line)
	}
	return
}

func bakeResult(rawResult ksl.StringSlice, delimiters string, multiline bool) (result ksl.StringSlice) {
	if multiline {
		indentString := "\t"
		result = ksl.StringSlice{delimiters[0:1]}
		for _, line := range rawResult {
			indentedLine := indentString + line
			result = append(result, indentedLine)
		}
		result = append(result, delimiters[1:2])
	} else {
		result = ksl.StringSlice{delimiters[0:1] + strings.Join(rawResult, "") + delimiters[1:2]}
	}
	return
}
