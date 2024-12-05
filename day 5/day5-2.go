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
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	rules := make([][]string, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		rule := strings.Split(scanner.Text(), "|")
		rules = append(rules, rule)
	}
	total := 0
	for scanner.Scan() {
		pages := strings.Split(scanner.Text(), ",")
		valid := true
		for i := 0; i < len(pages); i++ {
			for runs := 0; runs < 2; runs++ {
				for _, rule := range rules {
					if rule[0] == pages[i] || rule[1] == pages[i] {
						var rule1, rule2 int
						found1, found2 := false, false
						for k, _ := range pages {
							if pages[k] == rule[0] {
								rule1 = k
								found1 = true
							}
							if pages[k] == rule[1] {
								rule2 = k
								found2 = true
							}
						}
						if !(rule1 < rule2) && found1 && found2 {
							valid = false
							pages[rule1], pages[rule2] = rule[1], rule[0]
						}
					}
				}
			}

		}
		if !valid {
			fmt.Println(pages)
			value, err := strconv.Atoi(pages[(len(pages))/2])
			if err != nil {
				fmt.Println(err)
				return
			}
			total += value
		}
	}
	fmt.Println(total)
}
