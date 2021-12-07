package main

import (
	"fmt"
	"strconv"
	"strings"
)

//var initialState string = "3,4,3,1,2"

var initialState string = "2,5,5,3,2,2,5,1,4,5,2,1,5,5,1,2,3,3,4,1,4,1,4,4,2,1,5,5,3,5,4,3,4,1,5,4,1,5,5,5,4,3,1,2,1,5,1,4,4,1,4,1,3,1,1,1,3,1,1,2,1,3,1,1,1,2,3,5,5,3,2,3,3,2,2,1,3,1,3,1,5,5,1,2,3,2,1,1,2,1,2,1,2,2,1,3,5,4,3,3,2,2,3,1,4,2,2,1,3,4,5,4,2,5,4,1,2,1,3,5,3,3,5,4,1,1,5,2,4,4,1,2,2,5,5,3,1,2,4,3,3,1,4,2,5,1,5,1,2,1,1,1,1,3,5,5,1,5,5,1,2,2,1,2,1,2,1,2,1,4,5,1,2,4,3,3,3,1,5,3,2,2,1,4,2,4,2,3,2,5,1,5,1,1,1,3,1,1,3,5,4,2,5,3,2,2,1,4,5,1,3,2,5,1,2,1,4,1,5,5,1,2,2,1,2,4,5,3,3,1,4,4,3,1,4,2,4,4,3,4,1,4,5,3,1,4,2,2,3,4,4,4,1,4,3,1,3,4,5,1,5,4,4,4,5,5,5,2,1,3,4,3,2,5,3,1,3,2,2,3,1,4,5,3,5,5,3,2,3,1,2,5,2,1,3,1,1,1,5,1"

func main() {
	/** OLD brute force implementation
	fishesStr := strings.Split(initialState, ",")
	fishes := make([]int, 0)
	for _, f := range fishesStr {
		i, err := strconv.Atoi(f)
		check(err)
		fishes = append(fishes, i)
	}

	fishesPart1 := simulate(fishes, 80)
	fmt.Printf("Part 1 result : %d \n", len(fishesPart1))

	fishesPart2 := simulate(fishes, 256)
	fmt.Printf("Part 2 result : %d \n", len(fishesPart2))
	---- **/

	fishesStr := strings.Split(initialState, ",")
	// each possition in the array is the amount of fishes in that day
	// so we need 9 possitions -> 0 to 8
	fishes := make([]int64, 9)
	for _, f := range fishesStr {
		i, err := strconv.Atoi(f)
		check(err)
		fishes[i] += 1
	}
	fishesPart1 := simulate(fishes, 80)
	fmt.Printf("Part 1 result : %d \n", getTotal(fishesPart1))

	fishesPart2 := simulate(fishes, 256)
	fmt.Printf("Part 2 result : %d \n", getTotal(fishesPart2))
}

func getTotal(fishes []int64) int64 {
	total := int64(0)
	for _, i := range fishes {
		total += i
	}
	return total
}

func simulate(fishes []int64, days int) []int64 {
	for d := 1; d <= days; d++ {
		// fmt.Printf("Day %d Fishes %+v \n", d, fishes)
		newFishes := make([]int64, 9)
		for i, f := range fishes {
			if i == 0 {
				//new fish
				newFishes[8] += f
				//new iteration
				newFishes[6] += f
			} else {
				newFishes[i-1] += f
			}
			fishes = newFishes
		}
	}
	return fishes
}

func check(err error) {
	if err != nil {
		panic("ERROR: " + err.Error())
	}
}

/** OLD implementation
func simulateOLD(fishes []int, days int) []int {
	for d := 1; d <= days; d++ {
		newFishes := make([]int, 0)
		for _, f := range fishes {
			if f == 0 {
				//new fish
				newFishes = append(newFishes, 8)
				//new iteration
				newFishes = append(newFishes, 6)
			} else {
				newFishes = append(newFishes, f-1)
			}
			fishes = newFishes
		}
	}
	return fishes
}
**/
