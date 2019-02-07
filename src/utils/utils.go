package utils

import (
	"bufio"
	"os"
	"strings"
)

func Contains(q string, l []string) bool {
	for _, s := range l {
		if s == q {
			return true
		}
	}
	return false
}

func HasUser(q string, fi *os.File) bool {
	reader := bufio.NewReader(fi)
	for {
		if line, err := reader.ReadString('\n'); err == nil {
			if strings.HasPrefix(line, "\turl") {
				if strings.Contains(line, q) {
					//fmt.Print(line)
					return true
				}
			}
		} else {
			break
		}
	}
	return false

}

