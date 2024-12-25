package types

import "time"

type TVShow struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Status    string     `json:"status"`
	Premiered *time.Time `json:"premiered"`
	TVMaze    int64      `json:"tvmaze"`
	IMDB      *string    `json:"imdb"`
	Website   *string    `json:"website"`
	Updated   time.Time  `json:"updated"`
	Episodes  string     `json:"episodes"`
	Watching  int64      `json:"watching"`
	Image     *string    `json:"image"`
	Stream    *string    `json:"stream"`
	Runtime   *int64     `json:"runtime"`
}

type TVMazeSearchResult struct {
	Title   string  `json:"title"`
	Label   string  `json:"label"`
	ID      int64   `json:"id"`
	Text    string  `json:"text"`
	URL     string  `json:"url"`
	Image   *string `json:"image"`
	Year    string  `json:"year"`
	Status  string  `json:"status"`
	Runtime *int64  `json:"runtime"`
}
