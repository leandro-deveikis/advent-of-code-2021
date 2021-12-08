package main

// https://adventofcode.com/2021/day/1

import (
	"fmt"
)

func main() {
	println("* DAY 01: ")
	// first part - increase count
	fmt.Printf("--- PART 1 Result: %d \n", countIncreases(Challenge_input))
	// second part - three measurement sliding window
	fmt.Printf("--- PART 2 Result: %d \n", countIncreases(buildWindows(Challenge_input)))
}

func countIncreases(input []int) int {
	// this is not necesary but just in case
	if len(input) == 0 {
		return 0
	}
	// as we count the increases, the first iteration will
	// be compared with itself, so it will not count as an
	// increase
	var previous int = input[0]
	increase_count := 0
	for _, i := range input {
		if i > previous {
			increase_count++
		}
		previous = i
	}
	return increase_count
}

func buildWindows(challenge_input []int) []int {
	// this will build the three-measurement sliding window for the input passed. Start from the begging until we reach the last complete window possible
	output := make([]int, 0)
	for index := range challenge_input {
		if index < len(challenge_input)-2 {
			output = append(output, challenge_input[index]+challenge_input[index+1]+challenge_input[index+2])
		}
	}
	return output
}
