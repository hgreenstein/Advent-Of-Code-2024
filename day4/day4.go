package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	textData, err := os.ReadFile("./day4.input")
	if err != nil {
		panic(err)
	}
	splitLines := bytes.Split(textData, []byte("\n"))
	searchGrid := make([][]byte, 0, len(splitLines))
	for _, line := range splitLines {
		searchGrid = append(searchGrid, line)
	}
	part1(searchGrid)
	part2(searchGrid)
}

var searchWordBytes []byte

func part1(searchGrid [][]byte) {
	ROWS, COLS := len(searchGrid), len(searchGrid[0])
	searchWordBytes = []byte{'X', 'M', 'A', 'S'}
	result := 0
	gridCopy := make([][]byte, ROWS)
	for i, row := range searchGrid {
		gridCopy[i] = make([]byte, COLS)
		copy(gridCopy[i], row)
	}
	for rowIndex, row := range searchGrid {
		for colIndex, cell := range row {
			if cell == 'X' {
				foundWords := bfs(searchGrid, rowIndex, colIndex, ROWS, COLS, 1, [2]int{0, 0})
				if foundWords > 0 {
					// searchGrid[rowIndex][colIndex] = '*'
					result += foundWords
					// for _, row := range searchGrid {
					// 	fmt.Println(string(row))
					// }
					// fmt.Println()
					// removeCurrentStars(searchGrid, gridCopy)
				}
			}
		}
	}
	fmt.Printf("Found %v words\n", result)
}

func part2(searchGrid [][]byte) {
	result := 0
	ROWS, COLS := len(searchGrid), len(searchGrid[0])
	for row := 1; row < ROWS-1; row++ {
		for col := 1; col < COLS-1; col++ {
			if searchGrid[row][col] == 'A' {
				firstDiagUp, firstDiagDown := searchGrid[row-1][col-1], searchGrid[row+1][col+1]
				secondDiagUp, secondDiagDown := searchGrid[row-1][col+1], searchGrid[row+1][col-1]
				if firstDiagUp == 'M' && firstDiagDown == 'S' || firstDiagUp == 'S' && firstDiagDown == 'M' {
					if secondDiagUp == 'M' && secondDiagDown == 'S' || secondDiagUp == 'S' && secondDiagDown == 'M' {
						result++
					}
				}
			}
		}
	}
	fmt.Printf("Found %v X-MASes\n", result)
}

func bfs(searchGrid [][]byte, rowIndex, colIndex, ROWS, COLS, i int, setDirection [2]int) int {
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	if i == 1 {
		result := 0
		for _, direction := range directions {
			x, y := rowIndex+direction[0], colIndex+direction[1]
			if inBounds(x, y, ROWS, COLS) && searchGrid[x][y] == searchWordBytes[i] {
				if i == len(searchWordBytes)-1 {
					return 1
				}
				result += bfs(searchGrid, x, y, ROWS, COLS, i+1, direction)
			}
		}
		return result
	} else {
		x, y := rowIndex+setDirection[0], colIndex+setDirection[1]
		if inBounds(x, y, ROWS, COLS) && searchGrid[x][y] == searchWordBytes[i] {
			// if i == len(searchWordBytes)-1 || bfs(searchGrid, x, y, ROWS, COLS, i+1, setDirection) {
			// 	searchGrid[x][y] = '*'
			// 	return true
			// }
			if i == len(searchWordBytes)-1 {
				return 1
			}
			result := bfs(searchGrid, x, y, ROWS, COLS, i+1, setDirection)
			return result
		}

	}
	return 0
}

func inBounds(x, y, ROWS, COLS int) bool {
	return x >= 0 && x < ROWS && y >= 0 && y < COLS
}
