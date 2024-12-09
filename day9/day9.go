package main

import (
	"fmt"
	"os"
)

func main() {
	textInput, err := os.ReadFile("./day9.input")
	if err != nil {
		panic(err)
	}
	decodedData, curId, i := make([]int, 0, len(textInput)), 0, 0
	for ; i < len(textInput)-1; i += 2 {
		curFill, curFree := textInput[i]-'0', textInput[i+1]-'0'
		for curFill > 0 {
			decodedData = append(decodedData, curId)
			curFill--
		}
		for curFree > 0 {
			decodedData = append(decodedData, -1)
			curFree--
		}
		curId++
	}
	if i < len(textInput) {
		curFill := textInput[i] - '0'
		for curFill > 0 {
			decodedData = append(decodedData, curId)
			curFill--
		}
	}
	part1DecodedData := make([]int, len(decodedData), len(decodedData))
	copy(part1DecodedData, decodedData)
	part1(part1DecodedData)
	part2(decodedData)
}

func part1(decodedData []int) {
	left, right, length := 0, len(decodedData)-1, len(decodedData)
	for left < right {
		for left < length && decodedData[left] != -1 {
			left++
		}
		for left < length && decodedData[left] == -1 && right >= 0 && left < right {
			if decodedData[right] == -1 {
				right--
			} else {
				decodedData[left], decodedData[right] = decodedData[right], -1
				left++
				right--
			}

		}

	}
	result := 0
	for index, id := range decodedData {
		if id == -1 {
			break
		}
		result += index * id
	}
	fmt.Printf("Result of part 1: %v\n", result)
}

func part2(decodedData []int) {
	right := len(decodedData) - 1
	for right >= 0 {
		for right >= 0 && decodedData[right] == -1 {
			right--
		}
		if right == -1 {
			break
		}
		curId, blockSize := decodedData[right], 1
		right--
		for right >= 0 && decodedData[right] == curId {
			right--
			blockSize++
		}
		blockIndex := checkForEmptySpace(decodedData, blockSize, right)
		if blockIndex == -1 {
			continue
		}
		if curId == -1 {
			panic("-1 Id")
		}
		for i := right + blockSize; i >= right+1; i-- {
			if decodedData[blockIndex] != -1 {
				panic("Overwrote valid data")
			}
			decodedData[blockIndex] = decodedData[i]
			decodedData[i] = -1
			blockIndex++
		}
	}
	result := 0
	for index, id := range decodedData {
		if id == -1 {
			continue //Continue don't break, first empty space doesn't mean no more ids remaining like part 1
		}
		result += index * id
	}
	fmt.Printf("Result of part 2: %v\n", result)
}

func checkForEmptySpace(decodedData []int, length, maxIndex int) int {
	for i := 0; i <= maxIndex-length+1; i++ {
		found := true
		for j := 0; j < length; j++ {
			if decodedData[i+j] != -1 {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
