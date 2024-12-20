package types

type Recipe struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Coffee float64 `json:"coffee"`
	Water  float64 `json:"water"`
	Grind  string  `json:"grind"`
	Time   int64   `json:"time"`
	Notes  string  `json:"notes"`
}
