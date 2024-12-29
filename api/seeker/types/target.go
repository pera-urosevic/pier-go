package types

type Target struct {
	Title   string `json:"title"`
	Sources string `json:"sources"`
	Release string `json:"release"`
	Checked string `json:"checked"`
	Note    string `json:"note"`
}
