package groupie

// Date is a struct that contains an Index field.
// Index is a slice of structs, each containing an Id int and a Dates []string.
type Date struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}
