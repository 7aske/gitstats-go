package print

import "fmt"

func PrintMap(m map[int]int, keys []int) {
	row := 7
	for i := row -1; i > -1; i-- {
		for j := len(keys) - 1; j+i > -1; j -= 7 {
			field := m[j+i]
			if field > 9 {
				//fmt.Printf("[%d]", field)
				printCell(field)
			} else {
				printCell(field)
				//fmt.Printf("[ %d]", field)
			}
		}
		fmt.Println()
	}
}
func printCell(val int) {
	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 4:
		escape = "\033[1;30;47m"
	case val >= 4 && val < 9:
		escape = "\033[1;30;43m"
	case val >= 9:
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