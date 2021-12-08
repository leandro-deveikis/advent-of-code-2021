package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

// the file to read
var Input string = "./Day08/input"

func main() {
	println("--- Day 08: ")

	totalPart1 := calculatePart1(Input)
	fmt.Printf("Total part 1: %v\n", totalPart1)

	// Rules:
	// * 1 -> only with 2 digits
	// * 2 -> reminder after all
	// * 3 -> 5 digits, has the same as 1 but is not 9
	// * 4 -> only with 4 digits
	// * 5 -> 5 digits, same as 6 but with 1 char less
	// * 6 -> 6 digits, that does not have the same as 1
	// * 7 -> only with 3 digits
	// * 8 -> only with 7 digits
	// * 9 -> 6 digits, same as 4 but with 2 more digit
	// * 0 -> 6 digits and has the same as 1

	// warning: wires and digit are unsorted, should be sorted by characters to be compared
	lines := readInput(Input)
	totalPart2 := getTotal(lines)
	fmt.Printf("Total part 2: %v\n", totalPart2)
}

//--------------------------------//
//------------ PART 1 ------------//
//--------------------------------//
func calculatePart1(Input string) int {
	total := 0

	dat, err := os.Open(Input)
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	// BOARDS
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " | ")

		// we only need the second part
		for _, d := range strings.Split(l[1], " ") {
			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				total += 1
			}
		}
	}
	return total
}

//--------------------------------//
//------------ PART 2 ------------//
//--------------------------------//
func readInput(Input string) []string {
	dat, err := os.Open(Input)
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	// BOARDS
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getTotal(lines []string) int {
	total := 0
	for _, line := range lines {
		l := strings.Split(line, " | ")

		// wires will store the code for each number (sorted string) using the number as index
		wires := make([]string, 10)

		for areDigitUndiscovered(wires) {
			for _, d := range strings.Split(l[0], " ") {
				switch len(d) {
				case 2:
					wires[1] = sortString(d)
				case 3:
					wires[7] = sortString(d)
				case 4:
					wires[4] = sortString(d)
				case 5:
					if wires[6] != "" && containsAllButOne(d, wires[6]) {
						wires[5] = sortString(d)
					} else if wires[1] != "" && wires[9] != "" && wires[9] != d && containsAllChars(d, wires[1]) {
						wires[3] = sortString(d)
						break
					} else if allSetBut2(d, wires) {
						wires[2] = sortString(d)
					}
				case 6:
					if wires[4] != "" && containsAllChars(d, wires[4]) {
						wires[9] = sortString(d)
					} else if wires[1] != "" {
						if containsAllChars(d, wires[1]) {
							wires[0] = sortString(d)
						} else {
							wires[6] = sortString(d)
						}
					}
				case 7:
					wires[8] = sortString(d)
				}
				//fmt.Printf("wires: %v\n", wires)
			}
		}

		// decode second part
		total += calculateDigit(l[1], wires)
	}
	return total
}

func calculateDigit(line string, wires []string) int {
	digit := 0
	for i, d := range strings.Split(line, " ") {
		sortedD := sortString(d)
		for j, w := range wires {
			if sortedD == w {
				digit += j * (int(math.Pow(10, float64(3-i))))
			}
		}
	}
	return digit
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// with this will loop untill all are discovered
func areDigitUndiscovered(wires []string) bool {
	for _, w := range wires {
		if w == "" {
			return true
		}
	}
	return false
}

// check if 2 is the only uncovered
func allSetBut2(d string, wires []string) bool {
	for i, w := range wires {
		if i != 2 && w == "" {
			return false
		}
	}
	return true
}

func containsAllButOne(d, s string) bool {
	oneAlreadyNotFound := false
	for _, b := range []byte(s) {
		if strings.IndexByte(d, b) == -1 {
			if oneAlreadyNotFound {
				return false
			} else {
				oneAlreadyNotFound = true
			}
		}
	}
	return oneAlreadyNotFound
}

func containsAllChars(d, s string) bool {
	for _, b := range []byte(s) {
		if strings.IndexByte(d, b) == -1 {
			return false
		}
	}
	return true
}

// ----------- HELPER ----------//
func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}
