package types

type Article struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	FeedName  string `json:"feed_name"`
	Discarded bool   `json:"discarded"`
}
