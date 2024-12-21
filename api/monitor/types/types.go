package types

type MonitorData struct {
	Stats         Stats          `json:"stats"`
	Notifications []Notification `json:"notifications"`
}

type Stat struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Stats = []Stat

type Notification struct {
	ID        int64  `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Channel   string `json:"channel"`
	Topic     string `json:"topic"`
	Message   string `json:"message"`
}

type Notifications = map[string][]Notification
