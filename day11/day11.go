package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	textInput, err := os.ReadFile("day11.input")
	if err != nil {
		panic(err)
	}
	splitNumberStrings := strings.Split(string(textInput), " ")
	part1(splitNumberStrings, 25)
	part2(splitNumberStrings, 75)
}

func part1(splitNumberStrings []string, blinks int) {
	curRow := make([]string, len(splitNumberStrings), len(splitNumberStrings))
	copy(curRow, splitNumberStrings)
	for i := 0; i < blinks; i++ {
		nextRow := make([]string, 0, len(curRow)) //next rows can only get larger, start with min capacity of curRow
		for _, numString := range curRow {
			numlength := len(numString)
			num, err := strconv.Atoi(numString)
			if err != nil {
				panic(err)
			}
			switch true {
			case num == 0:
				nextRow = append(nextRow, "1")
			case numlength%2 == 0:
				halfIndex := numlength / 2
				parsedFirstHalf, err := strconv.Atoi(numString[0:halfIndex])
				if err != nil {
					panic(err)
				}

				parsedSecondHalf, err := strconv.Atoi(numString[halfIndex:])
				if err != nil {
					panic(err)
				}
				nextRow = append(nextRow, strconv.Itoa(parsedFirstHalf), strconv.Itoa(parsedSecondHalf))
			default:
				num *= 2024
				nextRow = append(nextRow, strconv.Itoa(num))
			}
		}
		curRow = nextRow
		nextRow = make([]string, 0, len(curRow))
	}
	fmt.Printf("Part 1 total stones after %v blinks: %v\n", blinks, len(curRow))
}

func part2(stoneStrings []string, blinks int) {
	result := 0
	for _, stoneString := range stoneStrings {
		stoneVal, err := strconv.Atoi(stoneString)
		if err != nil {
			panic(err)
		}
		result += part2RecursiveHelper(stoneVal, blinks)
	}
	fmt.Printf("Part 2 total stones after %v blinks: %v\n", blinks, result)
}

type StoneBlinkPair struct {
	Stone  int
	Blinks int
}

var part2DpMap map[StoneBlinkPair]int = map[StoneBlinkPair]int{}

func part2RecursiveHelper(stone, remainingBlinks int) int {
	if remainingBlinks == 0 {
		return 1
	}
	curDpMapEntry := StoneBlinkPair{
		Stone:  stone,
		Blinks: remainingBlinks,
	}
	if dpVal, ok := part2DpMap[curDpMapEntry]; ok {
		return dpVal
	}
	if stone == 0 {
		result := part2RecursiveHelper(1, remainingBlinks-1)
		part2DpMap[curDpMapEntry] = result
		return result
	}
	numString := strconv.Itoa(stone)
	length, halfIndex := len(numString), len(numString)/2
	if length%2 == 0 {
		parsedFirstHalf, err := strconv.Atoi(numString[0:halfIndex])
		if err != nil {
			panic(err)
		}

		parsedSecondHalf, err := strconv.Atoi(numString[halfIndex:])
		if err != nil {
			panic(err)
		}
		result := part2RecursiveHelper(parsedFirstHalf, remainingBlinks-1) + part2RecursiveHelper(parsedSecondHalf, remainingBlinks-1)
		part2DpMap[curDpMapEntry] = result
		return result
	}
	result := part2RecursiveHelper(stone*2024, remainingBlinks-1)
	part2DpMap[curDpMapEntry] = result
	return result
}
