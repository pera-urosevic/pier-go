package notify

import (
	"fmt"
	"time"

	"pier/database"
)

// internal

func notify(channel string, topic string, message string) {
	fmt.Println(channel, topic, message)
	db := database.Connect()
	now := time.Now()
	_, err := db.Exec("INSERT INTO `notify` (`timestamp`, `channel`, `topic`, `message`) VALUES (?, ?, ?, ?)", now.Unix(), channel, topic, message)
	if err != nil {
		fmt.Println(err)
	}
	db.Exec("DELETE FROM `notify` WHERE `channel`='info' AND `timestamp` < ?", now.Add(-1*time.Hour).Unix())
	db.Exec("DELETE FROM `notify` WHERE `channel`='warn' AND `timestamp` < ?", now.Add(-1*time.Hour*24).Unix())
	db.Exec("DELETE FROM `notify` WHERE `channel`='alert' AND `timestamp` < ?", now.Add(-1*time.Hour*24*7).Unix())
}

// notifications by type

func Info(topic string, message string) {
	notify("info", topic, message)
}

func Warn(topic string, message string) {
	notify("warn", topic, message)
}

func Alert(topic string, message string) {
	notify("alert", topic, message)
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
