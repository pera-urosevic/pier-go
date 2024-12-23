package types

type Entry struct {
	Name string `json:"name"`
	Dir  bool   `json:"dir"`
}

type SubtitleTrack struct {
	UUID      string
	TrackID   int
	TrackLang string
	TrackName string
}
