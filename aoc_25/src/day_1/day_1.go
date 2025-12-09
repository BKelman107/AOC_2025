package main

import (
	"aoc_25/src"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func Check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	path := "../../data/day_1/test_input.txt"
	data, err := os.ReadFile(path)
	src.Check(err)

	elements := strings.Split(string(data), "\n")
	zeroCounter := 0
	dial := 50
	limit := elements //[300:350]

	for _, elem := range limit {
		elem = strings.TrimSpace(elem)
		if len(elem) == 0 {
			continue
		}
		numberStr := elem[1:]
		number, err := strconv.Atoi(numberStr)
		src.Check(err)
		direction := elem[0]
		step := 1
		if direction == 'L' {
			step = -1
		}
		for i := 0; i < number; i++ {
			dial += step
			if dial < 0 {
				dial = 99
			}
			if dial > 99 {
				dial = 0
			}
			if dial == 0 {
				zeroCounter++
			}
		}
	}
	fmt.Println("zeroCounter=", zeroCounter)

}
