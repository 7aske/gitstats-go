package print

import (
	"fmt"
	"strings"
	"time"
)

func PrintMap(m map[int]int, keys []int) {
	daysInWeek := 7
	printHeader(len(keys))
	e1, e2 := "\033[1;30;36m", "\033[0m"
	for i := daysInWeek; i > 0; i-- {
		for j := len(keys) - 1  + dayOffset(); j+i > -1; j -= daysInWeek {
			field := m[j+i]
			if j == len(keys) - 1 + dayOffset(){
				fmt.Printf("%s%s %s", e1, getDay(i), e2)
			} else {
				printCell(field)

			}
		}
		fmt.Println()
	}
	dayOffset()
}

func printHeader(l int){
	out := ""
	weekOffset := int((31 - time.Now().Day()) / 7 - 1)
	offset := 0
	e1, e2 := "\033[1;30;35m", "\033[0m"
	stringOffset := "             "
	for i := 0; i < l; i += 7 * 4 {
		mon := strings.ToUpper(time.Now().AddDate(0,offset,0).Month().String()[:3]) + stringOffset
		out = mon + out
		offset--
	}
	fmt.Println(weekOffset)
	for j := 0; j < weekOffset; j++ {
		out = "    " + out
	}
	fmt.Println(e1 + out + e2)
}
func getDay(i int) string {
	out := ""
	switch i {
	case 7:
		out = "SUN"
	case 6:
		out = "MON"
	case 5:
		out =  "TUE"
	case 4:
		out =  "WED"
	case 3:
		out =  "THU"
	case 2:
		out =  "FRI"
	case 1:
		out =  "SAT"
	}
	return out
}
func dayOffset() int {
	switch time.Now().Weekday().String() {
	case "Monday":
		return 3
	case "Tuesday":
		return 2
	case "Wednesday":
		return 1
	case "Thursday":
		return 0
	case "Friday":
		return -1
	case "Saturday":
		return -2
	case "Sunday":
		return -3
	default:
		return 0
	}
}
func printCell(val int) {
	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 4:
		escape = "\033[1;30;47m"
	case val >= 4 && val < 8:
		escape = "\033[1;30;43m"
	case val >= 8:
		escape = "\033[1;30;42m"
	}

	//if today {
	//	escape = "\033[1;37;45m"
	//}

	if val == 0 {
		fmt.Printf(escape + "  - " + "\033[0m")
		return
	}

	str := "  %d "
	switch {
	case val >= 10:
		str = " %d "
	case val >= 100:
		str = "%d "
	}

	fmt.Printf(escape+str+"\033[0m", val)
}
