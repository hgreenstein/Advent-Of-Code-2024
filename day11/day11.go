package main

import (
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
	day1(splitNumberStrings, 25)
}

func day1(splitNumberStrings []string, blinks int) {
	curRow := splitNumberStrings
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
	println(len(curRow))
}
