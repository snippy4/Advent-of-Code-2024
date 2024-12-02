package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RemoveIndex(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		totalsafe := false
		for i := 0; i < len(levels); i++ {
			newlevels := make([]string, len(levels))
			copy(newlevels, levels)
			RemoveIndex(newlevels, i)
			newlevels = newlevels[:len(newlevels)-1]
			direction := "ascending"
			one, _ := strconv.Atoi(newlevels[0])
			two, _ := strconv.Atoi(newlevels[1])
			if one > two {
				direction = "descending"
			}
			safe := true
			for index := 0; index < len(newlevels)-1; index++ {
				current, _ := strconv.Atoi(newlevels[index])
				next, _ := strconv.Atoi(newlevels[index+1])
				if direction == "ascending" && !(current < next) {
					safe = false
				}
				if direction == "descending" && !(current > next) {
					safe = false
				}
				if (current-next > 3) || (current-next < -3) {
					safe = false
				}
			}
			if safe {
				totalsafe = true
			}
		}
		if totalsafe {
			count++
		}

	}
	fmt.Println(count)

}
