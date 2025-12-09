package main

import (
	"aoc_25/src"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "../../data/day_2/input_2.txt"
	data, err := os.ReadFile(path)
	src.Check(err)
	// fmt.Println(string(data))
	elements := strings.Split(string(data), ",")
	limit := elements
	invalid_ids := 0
	for _, elem := range limit {

		// fmt.Println(elem)
		num_range := strings.Split(elem, "-")
		start, err := strconv.Atoi(num_range[0])
		src.Check(err)
		end, err := strconv.Atoi(num_range[1])
		src.Check(err)

		for j := start; j < end+1; j++ {
			digits_as_strings := strconv.Itoa(j)
			number_of_digits := len(digits_as_strings)
			if number_of_digits%2 == 0 {
				midpoint := number_of_digits / 2
				first_half := digits_as_strings[0:midpoint]
				second_half := digits_as_strings[midpoint:number_of_digits]
				if first_half == second_half {
					fmt.Println("MATCHING PAIR FOUND:", j)
					invalid_ids += j
				}
			} else {
				continue
			}

		}
	}
	fmt.Println("Sum of invalid IDs=", invalid_ids)

}
