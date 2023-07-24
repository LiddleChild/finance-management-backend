package utils

import (
	"time"
)

func GetStartEndOfMonth(month int, year int) (int64, int64) {
	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	var endOfMonth time.Time
	endOfMonth = time.Date(year, time.Month(month + 1), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth = endOfMonth.Add(-time.Second)

	startTimestamp := startOfMonth.Unix()
	endTimestamp := endOfMonth.Unix()

	return startTimestamp, endTimestamp
}

func GetThreeMonthsRange(month int, year int) (int64, int64) {
	startOfMonth := time.Date(year, time.Month(month - 2), 1, 0, 0, 0, 0, time.UTC)

	var endOfMonth time.Time
	endOfMonth = time.Date(year, time.Month(month + 1), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth = endOfMonth.Add(-time.Second)

	startTimestamp := startOfMonth.Unix()
	endTimestamp := endOfMonth.Unix()

	return startTimestamp, endTimestamp
}