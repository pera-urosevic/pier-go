package notify

import (
	"fmt"
	"time"

	"pier/database"
)

// internal

func notify(channel string, topic string, message string, expire time.Duration) {
	db := database.Connect()
	// timestamp := time.Now().Format("20060102-150405-000000000")
	timestamp := time.Now().UnixMicro()
	key := fmt.Sprintf("notify:%s:%d:%s", channel, timestamp, topic)
	db.Set(database.Ctx, key, message, expire)
}

// notifications by type

func Alert(topic string, message string) {
	notify("alert", topic, message, time.Hour*24*7)
}

func Warn(topic string, message string) {
	notify("warn", topic, message, time.Hour*24)
}

func Info(topic string, message string) {
	notify("info", topic, message, time.Hour)
}

// utils

func ErrorAlert(topic string, subtopic string, err error) {
	message := fmt.Sprintf("%s\n%v", subtopic, err)
	Alert(topic, message)
}

func ErrorWarn(topic string, subtopic string, err error) {
	message := fmt.Sprintf("%s\n%v", subtopic, err)
	Warn(topic, message)
}

func ErrorInfo(topic string, subtopic string, err error) {
	message := fmt.Sprintf("%s\n%v", subtopic, err)
	Info(topic, message)
}
