package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	left_slice := make([]int, 0)
	right_slice := make([]int, 0)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), "   ")
		left_value, _ := strconv.Atoi(args[0])
		right_value, _ := strconv.Atoi(args[1])
		left_slice = append(left_slice, left_value)
		right_slice = append(right_slice, right_value)
	}
	sort.Ints(left_slice)
	sort.Ints(right_slice)
	diff := 0
	for index, left := range left_slice {
		if left-right_slice[index] > 0 {
			diff += left - right_slice[index]
		} else {
			diff += right_slice[index] - left
		}
	}
	fmt.Println(diff)

	// similarity score
	right_occurances := make(map[int]int)
	for _, right := range right_slice {
		if right_occurances[right] != 0 {
			right_occurances[right]++
		} else {
			right_occurances[right] = 1
		}
	}
	similarity := 0
	for _, left := range left_slice {
		count := right_occurances[left]
		similarity += left * count
	}
	fmt.Println(similarity)
}
