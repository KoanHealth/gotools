package strings

// Returns the first non-empty string
func FirstNonEmpty(choices ...string) string {
	for _, choice := range choices {
		if choice != "" {
			return choice
		}
	}
	return ""
}

// Returns the number of non empty strings
func CountNonEmpty(choices ...string) int {
	result := 0
	for _, choice := range choices {
		if choice != "" {
			result += 1
		}
	}
	return result
}
