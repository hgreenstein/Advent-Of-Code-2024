package main

import (
	"bytes"
	"fmt"
	"os"
)

type Coord struct {
	X int
	Y int
}

func main() {
	textInput, err := os.ReadFile("./day8.input")
	if err != nil {
		panic(err)
	}
	grid := bytes.Split(textInput, []byte{'\n'})
	charCoordMap := make(map[byte][]Coord)
	for rowIndex, row := range grid {
		for colIndex, cell := range row {
			if cell == '.' {
				continue
			}
			newCoord := Coord{
				X: rowIndex,
				Y: colIndex,
			}
			if curSlice, ok := charCoordMap[cell]; ok {
				charCoordMap[cell] = append(curSlice, newCoord)
			} else {
				charCoordMap[cell] = []Coord{newCoord}
			}
		}
	}
	ROWS, COLS := len(grid), len(grid[0])
	antiNodeMap := make(map[Coord]bool)

	part1(charCoordMap, ROWS, COLS, antiNodeMap)
	antiNodeMap = make(map[Coord]bool)
	part2(charCoordMap, ROWS, COLS, antiNodeMap)
}

func part2(charCoordMap map[byte][]Coord, ROWS, COLS int, antiNodeMap map[Coord]bool) {
	for _, coordSlice := range charCoordMap {

		for i, curCoord := range coordSlice {
			antiNodeMap[curCoord] = true
			for j := i + 1; j < len(coordSlice); j++ {

				nextCoord := coordSlice[j]
				xDiff, yDiff := curCoord.X-nextCoord.X, curCoord.Y-nextCoord.Y

				x1, y1 := curCoord.X+xDiff, curCoord.Y+yDiff
				x2, y2 := nextCoord.X-xDiff, nextCoord.Y-yDiff

				for inBounds(x1, y1, ROWS, COLS) {
					additionAntiNode := Coord{
						X: x1,
						Y: y1,
					}
					if _, ok := antiNodeMap[additionAntiNode]; !ok {
						antiNodeMap[additionAntiNode] = true
					}
					x1, y1 = x1+xDiff, y1+yDiff
				}
				for inBounds(x2, y2, ROWS, COLS) {
					subtractionAntiNode := Coord{
						X: x2,
						Y: y2,
					}
					if _, ok := antiNodeMap[subtractionAntiNode]; !ok {
						antiNodeMap[subtractionAntiNode] = true
					}
					x2, y2 = x2-xDiff, y2-yDiff
				}
			}
		}
	}
	fmt.Printf("Part 2 result %v\n", len(antiNodeMap))
}

func part1(charCoordMap map[byte][]Coord, ROWS int, COLS int, antiNodeMap map[Coord]bool) {
	for _, coordSlice := range charCoordMap {

		for i, curCoord := range coordSlice {

			for j := i + 1; j < len(coordSlice); j++ {

				nextCoord := coordSlice[j]
				xDiff, yDiff := curCoord.X-nextCoord.X, curCoord.Y-nextCoord.Y

				x1, y1 := curCoord.X+xDiff, curCoord.Y+yDiff
				x2, y2 := nextCoord.X-xDiff, nextCoord.Y-yDiff

				if inBounds(x1, y1, ROWS, COLS) {
					antiNode1 := Coord{
						X: x1,
						Y: y1,
					}
					if _, ok := antiNodeMap[antiNode1]; !ok {
						antiNodeMap[antiNode1] = true
					}
				}
				if inBounds(x2, y2, ROWS, COLS) {
					antiNode2 := Coord{
						X: x2,
						Y: y2,
					}
					if _, ok := antiNodeMap[antiNode2]; !ok {
						antiNodeMap[antiNode2] = true
					}
				}
			}
		}
	}
	fmt.Printf("Part 1 result %v\n", len(antiNodeMap))
}

func inBounds(x, y, ROWS, COLS int) bool {
	return x >= 0 && y >= 0 && x < ROWS && y < COLS
}
