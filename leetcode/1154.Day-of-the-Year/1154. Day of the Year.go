package leetcode

import "time"

func dayOfYear(date string) int {
	first := date[:4] + "-01-01"
	firstDay, _ := time.Parse("2006-01-02", first)
	dateDay, _ := time.Parse("2006-01-02", date)
	duration := dateDay.Sub(firstDay)
	return int(duration.Hours())/24 + 1
}
