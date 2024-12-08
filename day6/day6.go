package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	textInput, err := os.ReadFile("./day6.input")
	if err != nil {
		panic(err)
	}
	byteMap := bytes.Split(textInput, []byte{'\n'})
	// part1(byteMap)
	part2(byteMap)
}

func part1(grid [][]byte) {
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	// directionMap := map[int]string{
	// 	0: "Up",
	// 	1: "Right",
	// 	2: "Down",
	// 	3: "Left",
	// }
	curX, curY, curDirection := 0, 0, 0
	ROWS, COLS := len(grid), len(grid[0])
	//Get start position
startLoop:
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			if cell == '^' {
				curX, curY = rowIndex, colIndex
				break startLoop
			}
		}
	}
	result, numRedundant := 1, 0
	visited := make([][]bool, ROWS, ROWS)
	for i := range visited {
		visited[i] = make([]bool, COLS, COLS)
	}
	visited[curX][curY] = true
	for {
		nextX, nextY := curX+directions[curDirection][0], curY+directions[curDirection][1]
		if !inBounds(nextX, nextY, ROWS, COLS) {
			break
		}
		if grid[nextX][nextY] == '#' {
			// curDirectionStr, _ := directionMap[curDirection]
			curDirection = (curDirection + 1) % 4
			// nextDirectionStr, _ := directionMap[curDirection]
			// fmt.Printf("Rotating at pos (%v, %v) was facing %v now facing %v\n", curX, curY, curDirectionStr, nextDirectionStr)
		} else {
			// fmt.Printf("Moving from (%v, %v) to (%v, %v)\n", curX, curY, nextX, nextY)
			if !visited[nextX][nextY] {
				result++
				visited[nextX][nextY] = true
			} else {
				numRedundant++
			}
			curX, curY = nextX, nextY
		}
	}
	// for _, row := range grid {
	// 	fmt.Println(string(row))
	// }
	fmt.Printf("Part 1 total squares walked: %v\n", result)
	fmt.Printf("Num redundant %v\n", numRedundant)
}

func part2(grid [][]byte) {
	// directionMap := map[int]string{
	// 	0: "Up",
	// 	1: "Right",
	// 	2: "Down",
	// 	3: "Left",
	// }
	curX, curY, result := 0, 0, 0
	ROWS, COLS := len(grid), len(grid[0])
	//Get start position
startLoop:
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			if cell == '^' {
				curX, curY = rowIndex, colIndex
				break startLoop
			}
		}
	}
	initialX, initialY := curX, curY
	for xObstacle := range ROWS {
		for yObstacle := range COLS {
			if (xObstacle == initialX && yObstacle == initialY) || grid[xObstacle][yObstacle] != '.' {
				continue
			}
			grid[xObstacle][yObstacle] = '#'
			if checkForLoops(grid, curX, curY, ROWS, COLS) {
				result++
			}
			grid[xObstacle][yObstacle] = '.'
		}
	}
	fmt.Printf("Part 2 result %v\n", result)
}

type CoordWithDir struct {
	X         int
	Y         int
	Direction int
}

func checkForLoops(grid [][]byte, curX, curY, ROWS, COLS int) bool {
	directions, curDirection := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}, 0
	visited := make(map[CoordWithDir]bool)
	for {
		nextX, nextY := curX+directions[curDirection][0], curY+directions[curDirection][1]
		if !inBounds(nextX, nextY, ROWS, COLS) {
			return false
		}
		visited[CoordWithDir{
			X:         curX,
			Y:         curY,
			Direction: curDirection,
		}] = true
		if grid[nextX][nextY] == '#' {
			curDirection = (curDirection + 1) % 4
		} else {
			if _, ok := visited[CoordWithDir{
				X:         nextX,
				Y:         nextY,
				Direction: curDirection,
			}]; ok {
				return true
			}
			curX, curY = nextX, nextY
		}
	}
}

func inBounds(x, y, ROWS, COLS int) bool {
	return x >= 0 && y >= 0 && x < ROWS && y < COLS
}
