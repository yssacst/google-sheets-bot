package core

import (
	"time"
)

type Row struct {
	Name string
	Date time.Time
}

func IsOnDutyTomorrow(rows []Row, userName string, today time.Time) bool {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	today.In(loc)

	tomorrow := today.AddDate(0, 0, 1)

	for _, r := range rows {

		if r.Name != userName {
			continue
		}

		if sameDay(r.Date, tomorrow) {
			return true
		}
	}

	return false
}

func sameDay(a, b time.Time) bool {
	return a.Year() == b.Year() &&
		a.Month() == b.Month() &&
		a.Day() == b.Day()
}