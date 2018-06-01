package sets

import (
	"time"
	"sort"
	"strings"
	"fmt"
)

type TimeSet map[time.Time]struct{}

func NewTimeSet(items ...time.Time) TimeSet {
	set := TimeSet{}
	set.Add(items...)
	return set
}

func (s TimeSet) Add(items ...time.Time) {
	for _, item := range items {
		s[item] = empty
	}
}

func (s TimeSet) Items() []time.Time {
	var items []time.Time
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s TimeSet) SortedItems() []time.Time {
	items := s.Items()

	sort.Slice(items, func(i, j int) bool {
		return items[i].Before(items[j])
	})

	return items
}

func (s TimeSet) HasAny(items ...time.Time) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s TimeSet) HasAll(items ...time.Time) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s TimeSet) String() string {
	items := make([]string, 0, len(s))
	for k := range s {
		items = append(items, fmt.Sprintf("%v", k))
	}

	return strings.Join(items, ", ")
}
