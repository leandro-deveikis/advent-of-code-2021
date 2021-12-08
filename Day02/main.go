package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/2

type Position struct {
	horizontal int
	depth      int
}

type Position2 struct {
	horizontal int
	depth      int
	aim        int
}

func main() {
	println("* DAY 02: ")
	/**** PART 1 ****/
	resultPart1 := calculatePosition_part1(Challenge_input)
	fmt.Printf("--- PART 1 - Result position:  %+v \n", resultPart1)
	fmt.Printf("--- PART 1 - Multiplication result: %d \n", resultPart1.depth*resultPart1.horizontal)

	/**** PART 2 ****/
	resultPart2 := calculatePosition_part2(Challenge_input)
	fmt.Printf("--- PART 2 - Result position:  %+v \n", resultPart2)
	fmt.Printf("--- PART 2 - Multiplication result: %d \n", resultPart2.depth*resultPart2.horizontal)
}

func calculatePosition_part1(input []string) Position {
	result := Position{0, 0}
	for _, step := range input {
		switch {
		case strings.HasPrefix(step, "forward"):
			solveForwardPart1(&result, step)
		case strings.HasPrefix(step, "down"):
			solveDownPart1(&result, step)
		case strings.HasPrefix(step, "up"):
			solveUpPart1(&result, step)
		default:
			panic("step not recognized: " + step)
		}
	}
	return result
}

func solveForwardPart1(result *Position, step string) {
	result.horizontal += getAmount(step, "forward ")
}

func solveDownPart1(result *Position, step string) {
	result.depth += getAmount(step, "down ")
}

func solveUpPart1(result *Position, step string) {
	result.depth -= getAmount(step, "up ")
}

func calculatePosition_part2(input []string) Position2 {
	result := Position2{0, 0, 0}
	for _, step := range input {
		switch {
		case strings.HasPrefix(step, "forward"):
			solveForwardPart2(&result, step)
		case strings.HasPrefix(step, "down"):
			solveDownPart2(&result, step)
		case strings.HasPrefix(step, "up"):
			solveUpPart2(&result, step)
		default:
			panic("step not recognized: " + step)
		}
	}
	return result
}

func solveForwardPart2(result *Position2, step string) {
	a := getAmount(step, "forward ")
	result.depth += (result.aim * a)
	result.horizontal += a
}

func solveDownPart2(result *Position2, step string) {
	result.aim += getAmount(step, "down ")
}

func solveUpPart2(result *Position2, step string) {
	result.aim -= getAmount(step, "up ")
}

func getAmount(step string, toTrim string) int {
	i, err := strconv.Atoi(strings.TrimPrefix(step, toTrim))
	if err != nil {
		panic("ERROR in step: " + step + " error: " + err.Error())
	}
	return i
}
