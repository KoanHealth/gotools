package sets

import (
	"sort"
	"strconv"
	"strings"
)

type IntegerSet map[int]struct{}

func NewIntegerSet(items ...int) IntegerSet {
	set := IntegerSet{}
	set.Add(items...)
	return set
}

func (s IntegerSet) Add(items ...int) {
	for _, item := range items {
		s[item] = empty
	}
}

func (s IntegerSet) Items() []int {
	var items []int
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s IntegerSet) SortedItems() []int {
	items := s.Items()
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	return items
}

func (s IntegerSet) HasAny(items ...int) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s IntegerSet) HasAll(items ...int) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s IntegerSet) String() string {
	sorted := s.SortedItems()
	items := make([]string, len(sorted))

	for i := 0; i < len(sorted); i++ {
		items[i] = strconv.Itoa(sorted[i])
	}

	return strings.Join(items, ", ")
}
