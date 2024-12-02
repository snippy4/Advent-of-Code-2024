package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")
		direction := "ascending"
		one, _ := strconv.Atoi(levels[0])
		two, _ := strconv.Atoi(levels[1])
		if one > two {
			direction = "descending"
		}
		safe := true
		for index := 0; index < len(levels)-1; index++ {
			current, _ := strconv.Atoi(levels[index])
			next, _ := strconv.Atoi(levels[index+1])
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
			count++
		}

	}
	fmt.Println(count)

}
