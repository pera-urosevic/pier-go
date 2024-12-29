package types

type Feed struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Web      string `json:"web"`
	Icon     string `json:"icon"`
	Tokens   string `json:"tokens"`
	Disabled bool   `json:"disabled"`
	Updated  int64  `json:"updated"`
	Style    string `json:"style"`
}
