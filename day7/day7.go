package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	textInput, err := os.ReadFile("./day7.input")
	if err != nil {
		panic(err)
	}
	funMap = map[int]func(a, b int) int{
		0: func(a, b int) int {
			return a + b
		},
		1: func(a, b int) int {
			return a * b
		},
	}
	runRecursion(textInput, 1)
	//Add new concat operation to map of possible operations
	// funMap[2] = func(a, b int) int {
	// 	aDigits := make([]int, 0)
	// 	for a > 0 {
	// 		aDigits = append(aDigits, a%10)
	// 		a /= 10
	// 	}
	// 	bDigits := make([]int, 0)
	// 	for b > 0 {
	// 		bDigits = append(bDigits, b%10)
	// 		b /= 10
	// 	}
	// 	slices.Reverse(aDigits)
	// 	slices.Reverse(bDigits)
	// 	concatSlice := slices.Concat(aDigits, bDigits)
	// 	result := 0
	// 	for _, num := range concatSlice {
	// 		result *= 10
	// 		result += num
	// 	}
	// 	return result

	// }
	/** More efficient implementation of the concat function than above */
	funMap[2] = func(a, b int) int {
		var numBDigits float64 = 0.0
		tempB := b
		for tempB > 0 {
			numBDigits++
			tempB /= 10
		}
		aOffsetMultiplier := int(math.Pow(10, numBDigits))
		return a*aOffsetMultiplier + b
	}
	runRecursion(textInput, 2)
}

func runRecursion(textInput []byte, part int) {
	splitEquations := bytes.Split(textInput, []byte{'\n'})
	result, resultTotal := 0, 0
	for _, equation := range splitEquations {
		equationString := string(equation)
		splitTargetNum := strings.Split(equationString, ":")
		targetNumStr, equationNums := splitTargetNum[0], splitTargetNum[1]
		splitEquationNumStrs := strings.Split(equationNums, " ")[1:]
		splitEquationNums := make([]int, len(splitEquationNumStrs), len(splitEquationNumStrs))
		for i, numStr := range splitEquationNumStrs {
			parsedNum, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			splitEquationNums[i] = parsedNum
		}
		targetNum, err := strconv.Atoi(targetNumStr)
		if err != nil {
			panic(err)
		}
		if recursivePossibilites(targetNum, splitEquationNums, 1, 0, false) {
			// fmt.Printf("Satisfied target num %v with splitNums %v\n", targetNum, splitEquationNums)
			result++
			resultTotal += targetNum
		}
	}
	fmt.Printf("Result for part %v had %v possible equations that could be satisfied with a target nums totaling %v\n", part, result, resultTotal)
}

var funMap map[int]func(a, b int) int

func recursivePossibilites(targetNum int, equationNums []int, curIndex int, curTotal int, debug bool) bool {
	if debug {
		fmt.Printf("Entering recursivePossibilites: targetNum=%d, curIndex=%d, curTotal=%d, equationNums=%v\n", targetNum, curIndex, curTotal, equationNums)
	}
	if curIndex == len(equationNums) {
		if debug {
			fmt.Printf("At leaf node: curTotal=%d, targetNum=%d returning %v\n", curTotal, targetNum, curTotal == targetNum)
		}
		return curTotal == targetNum
	}
	for opIndex, operationFun := range funMap {
		var newTotal int
		if curIndex == 1 {
			newTotal = operationFun(equationNums[curIndex-1], equationNums[curIndex])
		} else {
			newTotal = operationFun(curTotal, equationNums[curIndex])
		}
		// newTotal := curTotal + operationFun(curTotal, equationNums[curIndex])
		if debug {
			fmt.Printf("Applying operation %d: newTotal=%d\n", opIndex, newTotal)
		}
		if recursivePossibilites(targetNum, equationNums, curIndex+1, newTotal, debug) {
			return true
		}
	}
	if debug {
		fmt.Printf("Exiting recursivePossibilites: targetNum=%d, curIndex=%d, curTotal=%d\n", targetNum, curIndex, curTotal)
	}
	return false
}

/** Before combination with part1 as "runRecursion" function */
// func part2(textInput []byte) {
// 	splitEquations := bytes.Split(textInput, []byte{'\n'})
// 	result, resultTotal := 0, 0
// 	for _, equation := range splitEquations {
// 		equationString := string(equation)
// 		splitTargetNum := strings.Split(equationString, ":")
// 		targetNumStr, equationNums := splitTargetNum[0], splitTargetNum[1]
// 		splitEquationNumStrs := strings.Split(equationNums, " ")[1:]
// 		splitEquationNums := make([]int, len(splitEquationNumStrs), len(splitEquationNumStrs))
// 		for i, numStr := range splitEquationNumStrs {
// 			parsedNum, err := strconv.Atoi(numStr)
// 			if err != nil {
// 				panic(err)
// 			}
// 			splitEquationNums[i] = parsedNum
// 		}
// 		targetNum, err := strconv.Atoi(targetNumStr)
// 		if err != nil {
// 			panic(err)
// 		}
// 		funMap = map[int]func(a, b int) int{
// 			0: func(a, b int) int {
// 				return a + b
// 			},
// 			1: func(a, b int) int {
// 				return a * b
// 			},
// 			2: func(a, b int) int {
// 				aDigits := make([]int, 0)
// 				for a > 0 {
// 					aDigits = append(aDigits, a%10)
// 					a /= 10
// 				}
// 				bDigits := make([]int, 0)
// 				for b > 0 {
// 					bDigits = append(bDigits, b%10)
// 					b /= 10
// 				}
// 				slices.Reverse(aDigits)
// 				slices.Reverse(bDigits)
// 				concatSlice := slices.Concat(aDigits, bDigits)
// 				result := 0
// 				for _, num := range concatSlice {
// 					result *= 10
// 					result += num
// 				}
// 				return result
// 			},
// 		}
// 		if recursivePossibilites(targetNum, splitEquationNums, 1, 0, false) {
// 			// fmt.Printf("Satisfied target num %v with splitNums %v\n", targetNum, splitEquationNums)
// 			result++
// 			resultTotal += targetNum
// 		}
// 	}
// 	fmt.Printf("Result for part 2 had %v possible equations that could be satisfied with a target nums totaling %v\n", result, resultTotal)

// }
