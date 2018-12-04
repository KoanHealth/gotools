package sets

import (
	"time"
	"sort"
	"strings"
	"fmt"
)

type TimeSet map[time.Time]struct{}

func NewTimeSet(items ...time.Time) TimeSet {
	set := NewSizedTimeSet(len(items))
	set.Add(items...)
	return set
}

func NewSizedTimeSet(capacity int) TimeSet {
	return make(TimeSet, capacity)
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
	sorted := s.SortedItems()
	items := make([]string, len(sorted))

	for i := 0; i < len(sorted); i++ {
		items[i] = fmt.Sprintf("%v", sorted[i])
	}

	return strings.Join(items, ", ")
}
