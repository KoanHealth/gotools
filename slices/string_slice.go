package slices

import (
	"fmt"
	"sort"

	. "github.com/koanhealth/gotools/strings"
)

type StringSlice []string
type StringFunc func(s string)

func ToStringSlice(src interface{}) StringSlice {
	if src != nil {
		switch typedSource := src.(type) {
		case string:
			return []string{typedSource}
		case []string:
			return typedSource
		case []interface{}:
			result := make([]string, 0, len(typedSource))
			for _, val := range typedSource {
				result = append(result, ToStringSlice(val)...)
			}
			return result
		default:
			return []string{fmt.Sprintf("%v", src)}
		}
	}
	return nil
}

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p StringSlice) Sort() { sort.Sort(p) }

func (p StringSlice) Reverse() StringSlice {
	output := make(StringSlice, len(p))
	targetIndex := len(output) - 1
	for _, str := range p {
		output[targetIndex] = str
		targetIndex -= 1
	}
	return output
}

// Returns subset of slice where filter is true
func (slice StringSlice) Select(filter StringPredicate) (result StringSlice) {
	for _, s := range slice {
		if filter(s) {
			result = append(result, s)
		}
	}
	return
}

// Removes items matching predicate
func (slice StringSlice) DeleteIf(filter StringPredicate) (result StringSlice) {
	return slice.Select(filter.Not())
}

// Removes items matching predicate (alias for DeleteIf)
func (slice StringSlice) Reject(filter StringPredicate) (result StringSlice) {
	return slice.DeleteIf(filter)
}

// Return true if any filter passes
func (slice StringSlice) Any(filter StringPredicate) bool {
	for _, s := range slice {
		if filter(s) {
			return true
		}
	}
	return false
}

// Returns true if filters pass for every item
func (slice StringSlice) All(filter StringPredicate) bool {
	for _, s := range slice {
		if !filter(s) {
			return false
		}
	}
	return true
}

// Returns true if string found in slice
func (slice StringSlice) Contains(match string) bool {
	filter := func(s string) bool { return s == match }
	return slice.Any(filter)
}

// Returns index of matching string
func (slice StringSlice) Index(match string) int {
	for i, s := range slice {
		if s == match {
			return i
		}
	}
	return -1
}

// Range over each item and call function
func (slice StringSlice) Each(f StringFunc) {
	for _, s := range slice {
		f(s)
	}
}

// Removes blank values
func (slice StringSlice) Compact() StringSlice {
	return slice.Reject(IsBlank)
}

// Returns first value
func (slice StringSlice) First() (string, bool) {
	if len(slice) > 0 {
		return slice[0], true
	}
	return "", false
}

// Returns last value
func (slice StringSlice) Last() (string, bool) {
	if len(slice) > 0 {
		return slice[len(slice)-1], true
	}
	return "", false
}
