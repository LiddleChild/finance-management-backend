package utils

import (
	"time"
)

func GetMonthsRange(month int, year int, month_range int) (int64, int64) {
	startOfMonth := time.Date(year, time.Month(month-(month_range-1)), 1, 0, 0, 0, 0, time.UTC)

	var endOfMonth time.Time
	endOfMonth = time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth = endOfMonth.Add(-time.Second)

	return startOfMonth.Unix(), endOfMonth.Unix()
}

func GetTodayRange() (int64, int64) {
	currentTime := time.Unix(1690787558, 0)
	startOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)

	var endOfDay time.Time
	endOfDay = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 0, 0, 0, 0, time.UTC)
	endOfDay = endOfDay.Add(-time.Second)

	return startOfDay.Unix(), endOfDay.Unix()
}
