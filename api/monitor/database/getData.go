package database

import (
	"pier/api/monitor/types"
	"pier/storage"
)

func GetData() (types.MonitorData, error) {
	var monitorData = types.MonitorData{}

	db, err := storage.DB()
	if err != nil {
		return monitorData, err
	}

	res := db.Find(&monitorData.Stats)
	if res.Error != nil {
		return monitorData, res.Error
	}

	res = db.Order("timestamp asc").Find(&monitorData.Notifications)
	if res.Error != nil {
		return monitorData, res.Error
	}

	return monitorData, nil
}
