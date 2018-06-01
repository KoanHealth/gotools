package codes

// Common code lookup interface that can be implemented by multiple structs
type CodeFinder interface {
	HasAny(...string) bool
	HasAll(...string) bool
}
