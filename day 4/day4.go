package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	grid := make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, strings.TrimSpace(scanner.Text()))
	}
	count := 0
	for y := 0; y < len(grid); y++ {
		for x, char := range grid[y] {
			if char == 'X' {
				for searchx := -1; searchx <= 1; searchx++ {
					for searchy := -1; searchy <= 1; searchy++ {
						if x+searchx*3 >= 0 && x+searchx*3 < len(grid[y]) && y+searchy*3 >= 0 && y+searchy*3 < len(grid) {
							if grid[y+searchy][x+searchx] == 'M' {
								if grid[y+searchy*2][x+searchx*2] == 'A' {
									if grid[y+searchy*3][x+searchx*3] == 'S' {
										count++
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}
