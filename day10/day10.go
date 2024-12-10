package main

import (
	"bytes"
	"fmt"
	"os"
)

var topographicMap [][]byte
var uniqueReachedNineCoordinates map[Coordinate]bool
var directions [4][2]int
var dp [][]int
var ROWS int
var COLS int

func main() {
	textInput, err := os.ReadFile("./day10.input")
	if err != nil {
		panic(err)
	}
	topographicMap = bytes.Split(textInput, []byte{'\n'})
	ROWS, COLS = len(topographicMap), len(topographicMap[0])
	directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dp = make([][]int, ROWS)
	for i := range dp {
		dp[i] = make([]int, COLS)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	result := 0
	part2Result := 0
	for rowIndex, row := range topographicMap {
		for colIndex, cell := range row {
			if cell == '0' {
				uniqueReachedNineCoordinates = make(map[Coordinate]bool)
				dfs(rowIndex, colIndex, '0', []string{})
				result += len(uniqueReachedNineCoordinates)
				part2Result += dfsPart2(rowIndex, colIndex, '0')
			}
		}
	}
	fmt.Printf("Total paths from trailheads part 1 %v\n", result)
	fmt.Printf("Total paths from trailheads part 2 %v\n", part2Result)
}

type Coordinate struct {
	X int
	Y int
}

func dfs(row, col int, targetVal byte, path []string) {
	if !inBounds(row, col) {
		return
	}
	curVal := topographicMap[row][col]
	if curVal != targetVal {
		return
	}
	path = append(path, fmt.Sprintf("(%d, %d)", row, col))
	if targetVal == '9' {
		uniqueReachedNineCoordinates[Coordinate{
			X: row,
			Y: col,
		}] = true
		return
	}
	for _, direction := range directions {
		nextX, nextY := row+direction[0], col+direction[1]
		// fmt.Printf("Moving from cell %v, %v to cell %v, %v with target value %c\n", row, col, nextX, nextY, targetVal+1)
		dfs(nextX, nextY, targetVal+1, path)
	}
	return
}

type CoordinateWithTarget struct {
	Coordinate
	TargetVal byte
}

var part2DP map[CoordinateWithTarget]int = map[CoordinateWithTarget]int{}

func dfsPart2(row, col int, targetVal byte) int {
	if !inBounds(row, col) {
		return 0
	}
	curVal := topographicMap[row][col]
	if curVal != targetVal {
		return 0
	}
	if targetVal == '9' {
		return 1
	}
	curCoordinate := CoordinateWithTarget{
		Coordinate: Coordinate{
			X: row,
			Y: col,
		},
		TargetVal: targetVal,
	}
	if dpVal, ok := part2DP[curCoordinate]; ok {
		return dpVal
	}
	result := 0
	for _, direction := range directions {
		nextX, nextY := row+direction[0], col+direction[1]
		// fmt.Printf("Moving from cell %v, %v to cell %v, %v with target value %c\n", row, col, nextX, nextY, targetVal+1)
		result += dfsPart2(nextX, nextY, targetVal+1)
	}
	part2DP[curCoordinate] = result
	return result
}

func inBounds(row, col int) bool {
	return row >= 0 && row < ROWS && col >= 0 && col < COLS
}
