package main

import (
	"fmt"
	"math/bits"
)

func main() {
	println("* DAY 03: ")
	input := Challenge_input
	length := getLengthRequired(input)

	/****** PART 1 ******/
	gamma := calculateGamma(input, length)
	epsilon := calculateEpsilon(gamma, length)

	resultPart1 := gamma * epsilon
	fmt.Printf("--- PART 1 - Gamma result:   binary=%b  decimal=%d \n", gamma, gamma)
	fmt.Printf("--- PART 1 - Epsilon result: binary=%b  decimal=%d \n", epsilon, epsilon)
	fmt.Printf("--- PART 1 - Final result:   binary=%b  decimal=%d \n", resultPart1, resultPart1)

	/****** PART 2 ******/
	oxigen := calculateOxigen(input, length)
	co2 := calculateCO2(input, length)
	resultPart2 := oxigen * co2
	fmt.Printf("--- PART 2 - oxigen binary=%b  decimal=%d \n", oxigen, oxigen)
	fmt.Printf("--- PART 2 - oxigen binary=%b  decimal=%d \n", co2, co2)
	fmt.Printf("--- PART 2 - Final result: binary=%b  decimal=%d \n", resultPart2, resultPart2)
}

func getLengthRequired(input []uint) int {
	length := 0
	// find length required
	for _, p := range input {
		l := bits.Len(p)
		if length < l {
			length = l
		}
	}
	return length
}

func getMostCommonBit(pos int, input []uint) uint {
	// fmt.Printf("Pos: %d oneCount: %d \n", i+1, oneCount)
	// add to gama result
	if float32(getOneCount(pos, input)) >= (float32(len(input)) / 2) {
		return 1
	} else {
		return 0
	}
}

func getLessCommonBit(pos int, input []uint) uint {
	if getMostCommonBit(pos, input) == 1 {
		return 0
	} else {
		return 1
	}
}

func getOneCount(pos int, input []uint) int {
	var p uint = 0b1 << pos
	oneCount := 0
	for _, j := range input {
		if (j & p) != 0 {
			oneCount += 1
		}
	}
	return oneCount
}

func filterInput(input []uint, mostCommon uint, pos int) []uint {
	result := make([]uint, 0)
	var p uint = 0b1 << pos
	for _, i := range input {
		if (i & p) == (mostCommon << pos) {
			result = append(result, i)
		}
	}
	return result
}

func calculateEpsilon(gamma uint, length int) uint {
	var mask uint = 0b0
	for i := 0; i < length; i++ {
		mask += 1 << i
	}
	return ^gamma & mask
}

func calculateGamma(input []uint, length int) uint {
	// invert gamma and remove digits that are not used
	var gamma uint
	for i := 0; i < length; i++ {
		// add to gama result
		if getMostCommonBit(i, input) == 1 {
			gamma += 0b1 << i
		}
	}
	return gamma
}

func calculateOxigen(input []uint, length int) uint {
	filteredInput := input
	// fmt.Printf("input=%b \n", input)
	for i := length - 1; i >= 0; i-- {
		mostCommon := getMostCommonBit(i, filteredInput)
		filteredInput = filterInput(filteredInput, mostCommon, i)
		// fmt.Printf("most common=%b filtered=%b \n", mostCommon, filteredInput)
		if len(filteredInput) == 1 {
			break
		}
	}
	return filteredInput[0]
}

func calculateCO2(input []uint, length int) uint {
	filteredInput := input
	// fmt.Printf("input=%b \n", input)
	for i := length - 1; i >= 0; i-- {
		lessCommon := getLessCommonBit(i, filteredInput)
		filteredInput = filterInput(filteredInput, lessCommon, i)
		// fmt.Printf("less common=%b filtered=%b \n", lessCommon, filteredInput)
		if len(filteredInput) == 1 {
			break
		}
	}
	return filteredInput[0]
}
