package fixgo

import (
	"sort"
	"strconv"
	"strings"
)

var empty struct{}

type SetKey interface {
	~int | ~int64 | ~string
}

type Set[T SetKey] map[T]struct{}

func NewSet[V SetKey](items ...V) Set[V] {
	set := Set[V]{}
	set.Add(items...)
	return set
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = empty
	}
}

func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s Set[T]) SortedItems() []T {
	items := s.Items()
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	return items
}

func (s Set[T]) HasAny(items ...T) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s Set[T]) HasAll(items ...T) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) String() string {
	sorted := s.SortedItems()
	items := make([]string, len(sorted))

	for i := 0; i < len(sorted); i++ {
		var s interface{} = sorted[i]
		switch v := s.(type) {
		case int:
			items[i] = strconv.Itoa(v)
		case int64:
			items[i] = strconv.FormatInt(v, 10)
		case string:
			items[i] = v
		}
	}

	return strings.Join(items, ", ")
}
