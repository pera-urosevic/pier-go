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

type Article struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	FeedName  string `json:"feed_name"`
	Discarded bool   `json:"discarded"`
}

type Bundle struct {
	Feeds    []Feed    `json:"feeds"`
	Articles []Article `json:"articles"`
}
