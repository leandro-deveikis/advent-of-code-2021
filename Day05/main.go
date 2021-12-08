package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	x1, y1, x2, y2 int
}

// the file to read
var Input string = "./Day05/input"

func main() {
	println("* DAY 05: ")
	lines := readInput(Input)
	// fmt.Printf("Lines %+v \n:", lines)
	filteredLines := filterOnlyEqualLines(lines)
	maxX, maxY := gerMaxXandY(filteredLines)
	diagramPart1 := buildDiagram(filteredLines, maxX, maxY)
	// fmt.Printf("Result  %+v \n:", diagram)
	resultPart1 := getCountOfDangerousSteps(diagramPart1)
	fmt.Printf("--- PART 1 Result: %d \n", resultPart1)

	//PART 2
	maxX2, maxY2 := gerMaxXandY(lines)
	diagramPart2 := buildDiagram(lines, maxX2, maxY2)
	//fmt.Printf("Result  %+v \n:", diagramPart2)
	resultPart2 := getCountOfDangerousSteps(diagramPart2)
	fmt.Printf("--- PART 2 Result: %d \n", resultPart2)
}

func getCountOfDangerousSteps(diagram [][]int) int {
	count := 0
	for _, d := range diagram {
		for _, x := range d {
			if x >= 2 {
				count += 1
			}
		}
	}
	return count
}

func buildDiagram(lines []Line, maxX, maxY int) [][]int {
	result := make([][]int, maxX+1)
	for i := 0; i <= maxY+1; i++ {
		result[i] = make([]int, maxX+1)
	}
	for _, l := range lines {
		diff, dirX, dirY := 0, 0, 0
		diffX := int(math.Abs(float64(l.x1 - l.x2)))
		if diffX != 0 {
			dirX = getDirection(l.x1, l.x2)
			diff = diffX
		}

		diffY := int(math.Abs(float64(l.y1 - l.y2)))
		if diffY != 0 {
			dirY = getDirection(l.y1, l.y2)
			diff = diffY
		}

		for i := 0; i <= diff; i++ {
			result[l.x1+(i*dirX)][l.y1+(i*dirY)] += 1
		}

		//for _, x := range result {
		//	fmt.Printf("%+v\n", x)
		//}
	}
	return result
}

func getDirection(i1, i2 int) int {
	d := i2 - i1
	if d == 0 {
		return 0
	} else if d > 1 {
		return 1
	} else {
		return -1
	}
}

func gerMaxXandY(lines []Line) (int, int) {
	maxX, maxY := 0, 0
	for _, l := range lines {
		maxX = max(l.x1, maxX)
		maxX = max(l.x2, maxX)

		maxY = max(l.y1, maxY)
		maxY = max(l.y2, maxY)
	}
	return maxX, maxY
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func filterOnlyEqualLines(lines []Line) []Line {
	result := make([]Line, 0)
	for _, l := range lines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			result = append(result, l)
		}
	}
	return result
}

func readInput(Input string) []Line {
	lines := make([]Line, 0)
	dat, err := os.Open(Input)
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	// BOARDS
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " -> ")
		val1 := strings.Split(c[0], ",")
		x1, err := strconv.Atoi(val1[0])
		check(err)
		y1, err := strconv.Atoi(val1[1])
		check(err)

		val2 := strings.Split(c[1], ",")
		x2, err := strconv.Atoi(val2[0])
		check(err)
		y2, err := strconv.Atoi(val2[1])
		check(err)

		lines = append(lines, Line{x1: x1, y1: y1, x2: x2, y2: y2})
	}
	return lines
}

func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
