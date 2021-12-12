package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var input string = "./Day09/input"

type LowPoint struct {
	pos Position
	val int
}

type Position struct {
	x, y int
}

type Basin struct {
	lowPoint LowPoint
	size     int
}

func main() {
	println("* DAY 09: ")
	lines := readInput(input)
	lowPoints := getLowPoints(lines)
	totalPart1 := solvePart1(lowPoints)
	fmt.Printf("--- PART 1 Result: %v\n", totalPart1)

	totalPart2 := solvePart2(lines, lowPoints)
	fmt.Printf("--- PART 2 Result: %v\n", totalPart2)
}

func getLowPoints(lines [][]int) []LowPoint {
	lenX := len(lines[0])
	//fmt.Printf("lenX: %v\n", lenX)
	lenY := len(lines)
	//fmt.Printf("lenY: %v\n", lenY)

	//find local mins
	lowPoints := make([]LowPoint, 0)
	for i, y := range lines {
		for j, x := range y {
			if isLocalMin(Position{y: i, x: j}, lines, lenX, lenY) {
				lowPoints = append(lowPoints, LowPoint{pos: Position{x: j, y: i}, val: x})
			}
		}
	}
	return lowPoints
}

func solvePart1(lowPoints []LowPoint) int {
	//calculate total
	total := 0
	for _, lp := range lowPoints {
		total += lp.val + 1
	}
	return total
}

func isLocalMin(p Position, lines [][]int, lenX, lenY int) bool {
	if isValidPos(Position{x: p.x, y: p.y - 1}, lines) &&
		!isSmallerThanAbove(p.y, p.x, lines, lenX, lenY) {
		// above
		return false
	} else if isValidPos(Position{x: p.x, y: p.y + 1}, lines) &&
		!isSmallerThanBelow(p.y, p.x, lines, lenX, lenY) {
		// below
		return false
	} else if isValidPos(Position{x: p.x + 1, y: p.y}, lines) &&
		!isSmallerThanRight(p.y, p.x, lines, lenX, lenY) {
		// right
		return false
	} else if isValidPos(Position{x: p.x - 1, y: p.y}, lines) &&
		!isSmallerThanLeft(p.y, p.x, lines, lenX, lenY) {
		// left
		return false
	} else {
		return true
	}
}

func isSmallerThanAbove(y, x int, lines [][]int, lenX, lenY int) bool {
	return lines[y][x] < lines[y-1][x]
}

func isSmallerThanBelow(y, x int, lines [][]int, lenX, lenY int) bool {
	return lines[y][x] < lines[y+1][x]
}

func isSmallerThanRight(y, x int, lines [][]int, lenX, lenY int) bool {
	return lines[y][x] < lines[y][x+1]
}

func isSmallerThanLeft(y, x int, lines [][]int, lenX, lenY int) bool {
	return lines[y][x] < lines[y][x-1]
}

func isValidPos(p Position, lines [][]int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(lines[0]) && p.y < len(lines)
}

func solvePart2(lines [][]int, lowPoints []LowPoint) int {
	basins := make([]Basin, 0)
	for _, lp := range lowPoints {
		b := Basin{lowPoint: lp, size: calculateBasinSize(lp, lines)}
		basins = append(basins, b)
	}

	// sorting descending
	sort.Slice(basins, func(i, j int) bool {
		return basins[i].size > basins[j].size
	})

	total := basins[0].size * basins[1].size * basins[2].size
	return total
}

func calculateBasinSize(lp LowPoint, lines [][]int) int {
	// start size in 1 to count for low point
	size := 0
	// we need to know all the positions already visited
	posAlreadyVisited := make([]Position, 0)
	visit(lp.pos, &size, &posAlreadyVisited, lines)
	return size
}

func visit(p Position, size *int, posAlreadyVisited *[]Position, lines [][]int) {
	// fmt.Printf("visit p: %+v  size: %d  pav: %v  \n", p, *size, posAlreadyVisited)
	if alreadyVisited(p, *posAlreadyVisited) {
		return
	} else {
		*size += 1
		*posAlreadyVisited = append(*posAlreadyVisited, p)
	}
	lenX := len(lines[0])
	//fmt.Printf("lenX: %v\n", lenX)
	lenY := len(lines)
	//fmt.Printf("lenY: %v\n", lenY)

	// visit each adyacent
	if isValidPos(Position{x: p.x, y: p.y - 1}, lines) &&
		isSmallerThanAbove(p.y, p.x, lines, lenX, lenY) &&
		lines[p.y-1][p.x] != 9 {
		// above
		visit(Position{x: p.x, y: p.y - 1}, size, posAlreadyVisited, lines)
	}

	if isValidPos(Position{x: p.x, y: p.y + 1}, lines) &&
		isSmallerThanBelow(p.y, p.x, lines, lenX, lenY) &&
		lines[p.y+1][p.x] != 9 {
		// below
		visit(Position{x: p.x, y: p.y + 1}, size, posAlreadyVisited, lines)
	}

	if isValidPos(Position{x: p.x + 1, y: p.y}, lines) &&
		isSmallerThanRight(p.y, p.x, lines, lenX, lenY) &&
		lines[p.y][p.x+1] != 9 {
		// right
		visit(Position{x: p.x + 1, y: p.y}, size, posAlreadyVisited, lines)
	}

	if isValidPos(Position{x: p.x - 1, y: p.y}, lines) &&
		isSmallerThanLeft(p.y, p.x, lines, lenX, lenY) &&
		lines[p.y][p.x-1] != 9 {
		// left
		visit(Position{x: p.x - 1, y: p.y}, size, posAlreadyVisited, lines)
	}
}

func alreadyVisited(pos Position, posAlreadyVisited []Position) bool {
	for _, p := range posAlreadyVisited {
		if p.x == pos.x && p.y == pos.y {
			return true
		}
	}
	return false
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
