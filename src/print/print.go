package print

import (
	"fmt"
	"strings"
	"time"
)

func OutputHistory(m map[int]int, keys []int, t time.Time) {
	daysInWeek := 7
	printHeader(len(keys), t)
	e1, e2 := "\033[1;30;36m", "\033[0m"
	for i := daysInWeek; i > 0; i-- {
		for j := len(keys) - 1 + dayOffset(t); j+i > -1; j -= daysInWeek {
			field := m[j+i]
			if j == len(keys)-1+dayOffset(t) {
				fmt.Printf("%s%s %s", e1, getDay(i), e2)
			} else {
				printCell(field, j+i == 0)
			}
		}
		fmt.Println()
	}
}

func printHeader(l int, t time.Time) {
	out := ""
	weekOffset := (31-t.Day())/7 - 1
	offset := 0
	e1, e2 := "\033[1;30;35m", "\033[0m"
	stringOffset := "     "
	for i := 0; i < l; i += 7 * 4 {
		mon := strings.ToUpper(t.AddDate(0, offset, 0).Month().String()[:3]) + stringOffset
		out = mon + out
		offset--
	}
	for j := 0; j < weekOffset; j++ {
		out = "    " + out
	}
	out = e1 + out + e2
	fmt.Println(out[:len(out)-12])
}

func printCell(val int, today bool) {
	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 4:
		escape = "\033[1;30;47m"
	case val >= 4 && val < 8:
		escape = "\033[1;30;43m"
	case val >= 8:
		escape = "\033[1;30;42m"
	}

	if today {
		escape = "\033[1;37;46m"
	}

	if val == 0 {
		fmt.Printf("\033[0;30;47;1m" + "██" + "\033[0m")
		return
	}

	str := " %d"
	switch {
	case val >= 10:
		str = "%d"
	case val >= 100:
		str = "%d"
	}

	fmt.Printf(escape+str+"\033[0m", val)
}
func getDay(i int) string {
	out := ""
	switch i {
	case 7:
		out = "SUN"
	case 6:
		out = "MON"
	case 5:
		out = "TUE"
	case 4:
		out = "WED"
	case 3:
		out = "THU"
	case 2:
		out = "FRI"
	case 1:
		out = "SAT"
	}
	return out
}
func dayOffset(t time.Time) int {
	switch t.Weekday().String() {
	case "Monday":
		return -6
	case "Tuesday":
		return -5
	case "Wednesday":
		return -4
	case "Thursday":
		return -3
	case "Friday":
		return -2
	case "Saturday":
		return -1
	case "Sunday":
		return 0
	default:
		return 0
	}
}
