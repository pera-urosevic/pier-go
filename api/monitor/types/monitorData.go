package types

type MonitorData struct {
	Stats         Stats          `json:"stats"`
	Notifications []Notification `json:"notifications"`
}
