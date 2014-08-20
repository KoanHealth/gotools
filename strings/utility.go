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
