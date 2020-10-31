package utils

import (
	"strconv"
	"strings"
)

func FilterWeeks(weeksStr string) string {
	weeks := ""
	w := []rune(strings.Replace(weeksStr, "周", "", -1))
	w = w[1 : len(w)-1]
	weeksArr := strings.Split(string(w), ",")
	for _, week := range weeksArr {
		if !strings.Contains(week, "-") {
			weeks += week + ","
			continue
		}

		weekArr := strings.Split(week, "-")
		startWeek, _ := strconv.Atoi(weekArr[0])
		endWeek, _ := strconv.Atoi(weekArr[1])

		for j := startWeek; j <= endWeek; j++ {
			if j == endWeek {
				weeks += strconv.Itoa(j)
				break
			}
			weeks += strconv.Itoa(j) + ","
		}
	}
	return weeks
}