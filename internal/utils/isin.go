package groupie

// Isin checks if the given element is present in the provided slice of strings.
// If the slice is nil, it returns false.
// Otherwise, it iterates through the slice and returns true if the element is found, false otherwise.
func Isin(element string, tab []string) bool {
	if tab == nil {
		return false
	}

	for _, elem := range tab {
		if elem == element {
			return true
		}
	}

	return false
}
