package core

import (
	"strings"
	"time"
)

type Row struct {
	Name string
	Date time.Time
}

func WhoIsOnDutyTomorrow(rows []Row) string {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	today := time.Now().In(loc)
	tomorrow := today.AddDate(0, 0, 1)

	for _, r := range rows {
		if sameDay(r.Date, tomorrow) {
			return getFirtName(r.Name)
		}
	}

	return ""
}

func getFirtName(fullName string) string {
	fullName = strings.TrimSpace(fullName)

	if fullName == "" {
		return ""
	}

	if i := strings.IndexByte(fullName, ' '); i != -1 {
		return fullName[:i]
	}

	return fullName
}

func sameDay(current, expect time.Time) bool {
	return current.Year() == expect.Year() &&
		current.Month() == expect.Month() &&
		current.Day() == expect.Day()
}
