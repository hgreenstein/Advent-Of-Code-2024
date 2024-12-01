package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	textData, err := os.ReadFile("./day1.txt")
	if err != nil {
		panic(err)
	}
	part1(textData)
	part2(textData)
}

func parseLines(textData []byte) (firstNums, secondNums []int) {
	splitLines, curLine := make([][]byte, 0), make([]byte, 0, 5)
	for _, c := range textData {
		if c == '\n' {
			splitLines = append(splitLines, curLine)
			curLine = make([]byte, 0, 5)
		} else {
			if c == ' ' {
				if len(curLine) > 0 && curLine[len(curLine)-1] != ',' {
					curLine = append(curLine, ',')
				}
			} else {
				curLine = append(curLine, c)
			}
		}
	}
	splitLines = append(splitLines, curLine)
	firstNums, secondNums = make([]int, 0, len(splitLines)), make([]int, 0, len(splitLines))
	for _, line := range splitLines {
		firstNum, secondNum := 0, 0
		var err error
		var sb strings.Builder
		for _, c := range line {
			if c == ',' {
				firstNum, err = strconv.Atoi(sb.String())
				if err != nil {
					panic(err)
				}
				sb.Reset()
			} else {
				sb.WriteByte(c)
			}
		}
		secondNum, err = strconv.Atoi(sb.String())
		if err != nil {
			panic(err)
		}
		firstNums = append(firstNums, firstNum)
		secondNums = append(secondNums, secondNum)
	}
	return firstNums, secondNums
}

func part1(textData []byte) {
	firstNums, secondNums := parseLines(textData)
	sort.Ints(firstNums)
	sort.Ints(secondNums)
	fmt.Println(firstNums, secondNums)
	total := 0
	for i, num := range firstNums {
		diff := int(math.Abs(float64(num - secondNums[i])))
		// fmt.Printf("The diff of %v and %v is %v\n", num, secondNums[i], diff)
		total += diff
	}
	fmt.Println(total)
}

func part2(textData []byte) {
	firstNums, secondNums := parseLines(textData)
	secondNumFreqMap := make(map[int]int, len(secondNums))
	for _, num := range secondNums {
		secondNumFreqMap[num]++
	}
	totalSimilarity := 0
	for _, num := range firstNums {
		freq, exists := secondNumFreqMap[num]
		if !exists { //multiplied by 0, doesn't affect total similarity
			continue
		}
		totalSimilarity += num * freq
	}
	fmt.Println(totalSimilarity)
}
