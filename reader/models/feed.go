package models

type Feed struct {
	Name     string
	Url      string
	Web      string
	Tokens   string
	Updated  int64
	Disabled bool
}
