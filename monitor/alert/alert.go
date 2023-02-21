package alert

import (
	"fmt"
	"localhost/pier/database"
	"localhost/pier/monitor/email"
	"strconv"
)

func Signal(subject string, threshold int, increase bool, text string) {
	db := database.Connect()

	if !increase {
		db.HSet(database.Ctx, "monitor:alert", subject, 0)
		return
	}

	res, err := db.HGet(database.Ctx, "monitor:alert", subject).Result()
	if err != nil {
		fmt.Println(err)
		res = "0"
	}
	score, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	score = score + 1
	fmt.Println("alert:", subject, score)

	if score >= threshold {
		email.Send(subject, text)
		db.HSet(database.Ctx, "monitor:alert", subject, 0)
		return
	}
	db.HSet(database.Ctx, "monitor:alert", subject, score)
}
