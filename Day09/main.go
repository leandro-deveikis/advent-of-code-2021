package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input string = "./input"

func main() {
	println("* DAY 09: ")

	totalPart1 := solvePart1(input)
	fmt.Printf("--- PART 1 Result: %v\n", totalPart1)

	totalPart2 := solvePart2(input)
	fmt.Printf("--- PART 2 Result: %v\n", totalPart2)
}

func solvePart1(input string) int {
	lines := readInput(input)

	lenX := len(lines[0])
	//fmt.Printf("lenX: %v\n", lenX)
	lenY := len(lines)
	//fmt.Printf("lenY: %v\n", lenY)

	//find local mins
	localMins := make([]int, 0)
	for i, y := range lines {
		for j, x := range y {
			if isLocalMin(i, j, lines, lenX, lenY) {
				localMins = append(localMins, x)
			}
		}
	}

	//calculate total
	total := 0
	for _, i := range localMins {
		total += i + 1
	}
	return total
}

func isLocalMin(i, j int, lines [][]int, lenX, lenY int) bool {
	val := lines[i][j]

	if i+1 < lenY && val >= lines[i+1][j] {
		// above
		return false
	} else if i != 0 && val >= lines[i-1][j] {
		// below
		return false
	} else if j+1 < lenX && val >= lines[i][j+1] {
		// right
		return false
	} else if j != 0 && val >= lines[i][j-1] {
		// left
		return false
	} else {
		return true
	}
}

func solvePart2(input string) int {
	return 0
}

func readInput(input string) [][]int {
	dat, err := os.Open(input)
	check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)

	lines := make([][]int, 0)
	for scanner.Scan() {
		currentLine := make([]int, 0)
		for _, n := range strings.Split(scanner.Text(), "") {
			i, err := strconv.Atoi(n)
			check(err)
			currentLine = append(currentLine, i)
		}
		lines = append(lines, currentLine)
	}

	// fmt.Printf("lines: %+v\n", lines)
	return lines
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
