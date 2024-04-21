package groupie

// GetLocation is a struct that contains an Index field, which is
// an array of structs containing Id and Locations fields. The Id field is an
// integer and the Locations field is an array of strings.
type GetLocation struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}
