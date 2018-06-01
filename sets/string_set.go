package sets

import (
	"sort"
	"strings"
)

type StringSet map[string]struct{}

func NewStringSet(items ...string) StringSet {
	set := StringSet{}
	set.Add(items...)
	return set
}

func (s StringSet) Add(items ...string) {
	for _, item := range items {
		s[item] = empty
	}
}

func (s StringSet) Items() []string {
	var items []string
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s StringSet) SortedItems() []string {
	items := s.Items()
	sort.Strings(items)
	return items
}

func (s StringSet) HasAny(items ...string) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s StringSet) HasAll(items ...string) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s StringSet) String() string {
	return strings.Join(s.SortedItems(), ", ")
}