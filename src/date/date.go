package date

import (
	"sort"
	"time"
)

func DaysSince(date time.Time) int {
	days := 0
	yr, mn, dy := time.Now().Date()
	start := time.Date(yr, mn, dy, 0, 0, 0, 0, time.Now().Location())
	for date.Before(start) {
		date = date.Add(24 * time.Hour)
		days++
	}
	return days
}

func GenerateHistory(limit time.Time) map[int]int {
	today := time.Now()
	var history = make(map[int]int, DaysSince(limit))
	count := 0
	for limit.Before(today) {
		limit = limit.Add(time.Hour * 24)
		history[count] = 0
		count++
	}
	return history
}

func GetSortedHistoryKeys(m *map[int]int) []int {
	var keys []int
	for _, k := range *m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
