package groupie

// isInList checks if a string element is contained in a string slice.
// It returns true if the element is found in the slice, false otherwise.
// The function returns false if the slice is nil.
func Isin(element string, tab []string) bool {
	if tab == nil {
		return false
	} else {
		for _, elements := range tab {
			if elements == element {
				return true
			}
		}
	}
	return false
}
