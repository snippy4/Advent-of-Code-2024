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
		for _, page := range pages {
			for _, rule := range rules {
				if rule[0] == page {
					if !(strings.LastIndex(scanner.Text(), page) < strings.Index(scanner.Text(), rule[1])) && strings.LastIndex(scanner.Text(), rule[1]) >= 0 {
						valid = false
					}
				}
			}

		}
		if valid {
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
