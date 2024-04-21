package groupie

// FeatureCollection is a struct that contains a slice of Features.
// Each Feature has a Center field that is a slice of float64 values.
type FeatureCollection struct {
	Features []struct {
		Center []float64 `json:"center"`
	} `json:"features"`
}
