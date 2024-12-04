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
			if char == 'A' {
				for searchx := -1; searchx <= 1; searchx += 2 {
					for searchy := -1; searchy <= 1; searchy += 2 {
						if x > 0 && x+1 < len(grid[y]) && y > 0 && y+1 < len(grid) {
							if grid[y+searchy][x+searchx] == 'M' && grid[y-searchy][x-searchx] == 'S' {
								if grid[y+searchy][x-searchx] == 'M' && grid[y-searchy][x+searchx] == 'S' {
									count++
								}
								if grid[y+searchy][x-searchx] == 'S' && grid[y-searchy][x+searchx] == 'M' {
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count / 2)
}
