package utils

import (
	"time"
)

func GetMonthsRange(month int, year int, month_range int) (int64, int64) {
	startOfMonth := time.Date(year, time.Month(month - (month_range - 1)), 1, 0, 0, 0, 0, time.UTC)

	var endOfMonth time.Time
	endOfMonth = time.Date(year, time.Month(month + 1), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth = endOfMonth.Add(-time.Second)

	startTimestamp := startOfMonth.Unix()
	endTimestamp := endOfMonth.Unix()

	return startTimestamp, endTimestamp
}