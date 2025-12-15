package main

import (
	"aoc_25/src"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "../../data/day_5/test_input.txt"
	data, err := os.ReadFile(path)
	src.Check(err)

	test := strings.Split(string(data), "\n")

	fmt.Println(test)

	ids_ranges, ids := splitInputs(test)
	fmt.Println(ids_ranges)
	fmt.Println(ids)

}

func splitInputs(test []string) ([]string, []int) {
	id_ranges := []string{}
	ids := []int{}

	for _, lines := range test {
		lines = strings.TrimSpace(lines)
		if lines == "" {
			continue
		}
		if strings.Contains(lines, "-") {
			id_ranges = append(id_ranges, lines)
		} else {
			id, err := strconv.Atoi(lines)
			src.Check(err)
			fmt.Println("Parsed ID:", id)
			ids = append(ids, id)
		}

	}
	return id_ranges, ids
}
