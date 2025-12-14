package main

import (
	"aoc_25/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "../../data/day_4/input_4.txt"
	data, err := os.ReadFile(path)
	src.Check(err)
	// fmt.Println(string(data))

	var grid [][]string
	for _, line := range strings.Split(string(data), "\n") {
		var row []string
		for _, ch := range line {
			row = append(row, string(ch))
		}
		grid = append(grid, row)

	}
	// Example: print the grid
	// fmt.Println(grid)
	fmt.Println("Grid length:", len(grid), "Row length:", len(grid[0]))
	totalCounter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// fmt.Println("i=", i, "j=", j, "value=", grid[i][j])
			if grid[i][j] == "@" {
				adj := countAdjacent(grid, i, j)
				if adj < 4 {
					totalCounter++
				}
			}
		}
	}
	fmt.Println("Total counter:", totalCounter)
	// Current thought process is to create a 2d array and iterate through it
	// Create a serperate function that iterates through ranges (i-1,i+1) and (j-1)(j+1)
	// i.e. each adjacent square and adds to a list if there is something there
	// it returns the number of times it found an object and if this is less than four
	// then i can add one to the total counter. The final result is the total counter after all loops
}

// countAdjacent counts the number of non-empty adjacent cells around (i, j)
func countAdjacent(grid [][]string, i, j int) int {
	counter := 0
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			// Skip the center cell
			if k == i && l == j {
				continue
			}
			// Check bounds
			// if k < 0 || k >= len(grid) || l < 0 || l >= len(grid[0]) {
			// 	continue
			// }
			if k < 0 || k >= len(grid) || l < 0 || l >= len(grid[k]) {
				continue
			}
			if grid[k][l] == "@" {
				// Found a non-empty adjacent cell
				counter++
			}
		}
	}
	return counter

}
