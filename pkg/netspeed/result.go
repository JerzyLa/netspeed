package netspeed

// Result stores download and upload network speeds
type Result struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
}
