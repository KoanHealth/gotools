package strings

import "strings"

// IS this building
type StringPredicate func(s string) bool

var IsBlank = StringPredicate(func(s string) bool { return len(strings.TrimSpace(s)) == 0 })

func (predicate StringPredicate) And(other StringPredicate) StringPredicate {
	return func(s string) bool {
		return predicate(s) && other(s)
	}
}

func (predicate StringPredicate) Or(other StringPredicate) StringPredicate {
	return func(s string) bool {
		return predicate(s) || other(s)
	}
}

func Not(predicate StringPredicate) StringPredicate {
	return func(s string) bool {
		return !predicate(s)
	}
}
