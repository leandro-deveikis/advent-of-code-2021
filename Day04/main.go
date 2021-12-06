package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	numbers [][]Field
}

type Field struct {
	value  int
	marked bool
}

// the file to read
var Input string = "./input"

func main() {
	numbers, boards := readInput(Input)
	// fmt.Printf("Numbers: %d \n", numbers)
	// fmt.Printf("Boards: %+v \n", boards)

	result := findResult(numbers, boards)
	fmt.Printf("result: %d \n", result)
}

func findResult(numbers []int, boards []*Board) int {
	for _, n := range numbers {
		for _, b := range boards {
			//check if number is present in any board
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if b.numbers[i][j].value == n {
						b.numbers[i][j].marked = true
						// checking if line is completed
						lineCompleted := true
						for k := 0; k < 5; k++ {
							lineCompleted = lineCompleted && b.numbers[i][k].marked
						}
						if lineCompleted {
							return calculateSumOfBoard(b) * n
						}
						// checking if collumn is completed
						colCompleted := true
						for k := 0; k < 5; k++ {
							colCompleted = colCompleted && b.numbers[k][j].marked
						}
						if colCompleted {
							return calculateSumOfBoard(b) * n
						}
					}
				}
			}
		}
	}
	return 0
}

func calculateSumOfBoard(b *Board) int {
	n := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.numbers[i][j].marked {
				n += b.numbers[i][j].value
			}
		}
	}
	return n
}

func readInput(filename string) ([]int, []*Board) {
	numbers := make([]int, 0)
	boards := make([]*Board, 0)

	dat, err := os.Open(Input)
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)

	// FIRST LINE
	scanner.Scan()
	firstLine := scanner.Text()

	for _, n := range strings.Split(firstLine, ",") {
		i, err := strconv.Atoi(n)
		check(err)
		numbers = append(numbers, i)
	}

	x := 0
	currentBoard := getNewEmptyBoard()
	boards = append(boards, currentBoard)

	// BOARDS
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			for _, n := range strings.Split(line, " ") {
				if len(n) > 0 && n != " " {
					intValue, err := strconv.Atoi(n)
					check(err)
					currentBoard.numbers[x] = append(currentBoard.numbers[x], Field{value: intValue, marked: false})
				}
			}

			if x == 4 {
				x = 0
				currentBoard = getNewEmptyBoard()
				boards = append(boards, currentBoard)
			} else {
				x += 1
				currentBoard.numbers = append(currentBoard.numbers, make([]Field, 0))
			}
		}
	}
	// remove last element from boards as it will be empty
	return numbers, boards[:len(boards)-1]
}

func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}

func getNewEmptyBoard() *Board {
	b := new(Board)
	b.numbers = make([][]Field, 0)
	b.numbers = append(b.numbers, make([]Field, 0))
	return b
}
