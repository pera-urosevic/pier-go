package types

type Bundle struct {
	Feeds    []Feed    `json:"feeds"`
	Articles []Article `json:"articles"`
}
