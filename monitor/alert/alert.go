package alert

import (
	"fmt"
	"strconv"

	"github.com/pera-urosevic/pier/database"
	"github.com/pera-urosevic/pier/notify"
)

func Signal(subject string, threshold int, increase bool, text string) {
	db := database.Connect()

	if !increase {
		db.HSet(database.Ctx, "monitor:alert", subject, 0)
		return
	}

	res, err := db.HGet(database.Ctx, "monitor:alert", subject).Result()
	if err != nil {
		res = "0"
	}
	score, err := strconv.Atoi(res)
	if err != nil {
		notify.ErrorAlert("monitor", "parse score", err)
		return
	}

	score = score + 1
	notify.Alert("monitor", fmt.Sprintf("%s - %d", subject, score))

	if score >= threshold {
		notify.Alert(subject, text)
		db.HSet(database.Ctx, "monitor:alert", subject, 0)
		return
	}
	db.HSet(database.Ctx, "monitor:alert", subject, score)
}
