package database

import (
	"pier/api/monitor/types"
	"pier/database"

	_ "modernc.org/sqlite"
)

/*
try {
	const rows = await pier.query('SELECT * FROM `monitor`')
	for (const row of rows) {
		const [sensor, metric] = row.key.split(':')
		if (!records[sensor]) records[sensor] = {}
		records[sensor][metric] = row.value
	}
} catch (err) {
	console.log(err)
}
const stats = records as Stats

const notifications: Notifications = {}
try {
	const rows = await pier.query('SELECT * FROM `notify` ORDER BY `timestamp` DESC')
	for (const row of rows) {
		const { id, timestamp, channel, topic, message } = row
		if (!notifications[channel]) notifications[channel] = []
		notifications[channel].push({
			id: id,
			timestamp: timestamp,
			channel: channel,
			topic: topic,
			message: message,
		})
	}
} catch (err) {
	console.log(err)
}
*/

func GetData() (types.MonitorData, error) {
	var monitorData = types.MonitorData{}
	db := database.Connect()
	rows, err := db.Query("SELECT * FROM `monitor`")
	if err != nil {
		return monitorData, err
	}
	var stats types.Stats
	for rows.Next() {
		var stat types.Stat
		err := rows.Scan(&stat.Key, &stat.Value)
		if err != nil {
			return monitorData, err
		}
		stats = append(stats, stat)
	}
	monitorData.Stats = stats

	rows, err = db.Query("SELECT `id`, `timestamp`, `channel`, `topic`, `message` FROM `notify` ORDER BY `timestamp` DESC")
	if err != nil {
		return monitorData, err
	}
	var notifications []types.Notification
	for rows.Next() {
		var notification types.Notification
		err := rows.Scan(&notification.ID, &notification.Timestamp, &notification.Channel, &notification.Topic, &notification.Message)
		if err != nil {
			return monitorData, err
		}
		notifications = append(notifications, notification)
	}
	monitorData.Notifications = notifications

	return monitorData, nil
}

func RemoveNotification(id int64) error {
	db := database.Connect()
	_, err := db.Exec("DELETE FROM `notify` WHERE `id` = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func RemoveNotifications(channel string) error {
	db := database.Connect()
	_, err := db.Exec("DELETE FROM `notify` WHERE `channel` = ?", channel)
	if err != nil {
		return err
	}
	return nil
}
