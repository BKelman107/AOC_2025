package main

import (
	"aoc_25/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "../../data/day_3/input.txt"
	data, err := os.ReadFile(path)
	src.Check(err)
	combined_num := 0
	for i, bank := range strings.Split(string(data), "\n") {
		bank = strings.TrimSpace(bank)
		// fmt.Println(i, bank)
		// fmt.Printf("Type of bank: %T\n", bank)
		max_l := int(bank[0] - '0')
		max_r := int(bank[1] - '0')
		// fmt.Println("Length of bank string:", len(bank[2:]))
		for j := range bank[2:] {
			battery := int(bank[j+2] - '0')
			// fmt.Println(j, battery)
			if j+1 == len(bank)-2 {
				if battery > max_r {
					max_r = battery
					break
				}
			}
			if battery > max_l {
				max_l = battery
				max_r = int(bank[j+1] - '0')
			} else if battery > max_r {
				max_r = battery
			}
		}
		larget_num := max_l*10 + max_r
		combined_num += larget_num
		fmt.Println(i, larget_num, combined_num)
	}
	// combined_num += combined_num
	fmt.Println("Combined number", combined_num)
}
