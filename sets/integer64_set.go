package sets

import (
	"sort"
	"strings"
	"strconv"
)

type Integer64Set map[int64]struct{}

func NewInteger64Set(items ...int64) Integer64Set {
	set := Integer64Set{}
	set.Add(items...)
	return set
}

func (s Integer64Set) Add(items ...int64) {
	for _, item := range items {
		s[item] = empty
	}
}

func (s Integer64Set) Items() []int64 {
	var items []int64
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s Integer64Set) SortedItems() []int64 {
	items := s.Items()
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	return items
}

func (s Integer64Set) HasAny(items ...int64) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s Integer64Set) HasAll(items ...int64) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s Integer64Set) String() string {
	items := make([]string, 0, len(s))
	for k := range s {
		items = append(items, strconv.FormatInt(k, 10))
	}

	return strings.Join(items, ", ")
}