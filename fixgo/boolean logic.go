package fixgo

// All Simple logic function that doesn't short circuit, e.g., all inputs for All() are evaluated even if the first is false
func All(input ...bool) bool {
	if len(input) > 0 {
		result := true
		for _, i := range input {
			result = result && i
		}
		return result
	} else {
		return false
	}
}

// Any Simple logic function that doesn't short circuit, e.g., all inputs for Any() are evaluated even if the first is true
func Any(input ...bool) bool {
	if len(input) > 0 {
		result := false
		for _, i := range input {
			result = result || i
		}
		return result
	} else {
		return false
	}
}
