package core

import (
	"time"
)

type Row struct {
	Name string
	Date time.Time
}

// if the cron job runs before noon, notify the person on duty today.
// if the cron job runs after noon, notify the person on duty tomorrow.
func WhoIsOnDuty(rows []Row) (string, string) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	today := time.Now().In(loc)
	
	day := today
	when := "hoje"

	if time.Now().Hour() > 12 {
		day = today.AddDate(0, 0, 1)
		when = "amanhã"
	}

	for _, r := range rows {
		if sameDay(r.Date, day) {
			return r.Name, when
		}
	}

	return "",""
}

func sameDay(current, expect time.Time) bool {
	return current.Year() == expect.Year() &&
		current.Month() == expect.Month() &&
		current.Day() == expect.Day()
}
