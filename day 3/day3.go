package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	total := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {

		tokens := strings.SplitAfter(scanner.Text(), "mul(")
		for i, token := range tokens {
			if i > 0 {
				if strings.Contains(token, ")") {
					if strings.HasSuffix(tokens[i-1], "mul(") {
						multoken := strings.Split(token, ")")[0]
						multokenpieces := strings.Split(multoken, ",")
						firstint, err := strconv.Atoi(multokenpieces[0])
						if err != nil {
							continue
						}
						secondint, err := strconv.Atoi(multokenpieces[1])
						if err != nil {
							continue
						}
						if firstint < 1000 && secondint < 1000 && len(multokenpieces) == 2 {
							total += firstint * secondint
						}
					}
				}
			}

		}
	}
	fmt.Println(total)
}
