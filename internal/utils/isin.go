package groupie

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
