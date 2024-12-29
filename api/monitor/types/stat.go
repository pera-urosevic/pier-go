package types

type Stat struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Stats = []Stat
