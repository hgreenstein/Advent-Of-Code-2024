package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	textData, err := os.ReadFile("./day3.input")
	if err != nil {
		panic(err)
	}
  part1(textData)
  part2(textData)
}

func part1(textData []byte) {
	inputStr, inputLength := string(textData), len(textData)
  result := 0
outer:
	for i := 0; i < inputLength-3; i++ {
		if inputStr[i:i+4] == "mul(" {
			i += 4
      if i >= inputLength {
        break outer
      }
      firstNumBytes := make([]byte, 0, 3)
      for i < inputLength && isDigit(inputStr[i]) {
        firstNumBytes = append(firstNumBytes, inputStr[i])
        if len(firstNumBytes) > 3 {
          continue outer
        }
        i++
      }
      if i >= inputLength {
        break outer
      }
      if inputStr[i] != ',' {
        continue outer
      }
      i++
      firstNum, err := strconv.Atoi(string(firstNumBytes))
      if err != nil {
        panic(err)
      }
      // fmt.Printf("First num %v\n", firstNum)
      if i >= inputLength {
        break outer
      }
      secondNumBytes := make([]byte, 0, 3)
      for i < inputLength && isDigit(inputStr[i]) {
        secondNumBytes = append(secondNumBytes, inputStr[i])
        if len(secondNumBytes) > 3 {
          continue outer
        }
        i++
      }
      if i >= inputLength {
        break outer
      }
      if inputStr[i] != ')' {
        continue outer
      }
      secondNum, err := strconv.Atoi(string(secondNumBytes))
      // fmt.Printf("Second num %v\n", secondNum)
      if err != nil {
        panic(err)
      }
      result += firstNum * secondNum
      // fmt.Println(secondNum)
		}
	}
  fmt.Printf("Result %v\n", result)
}

func part2(textData []byte) {
	inputStr, inputLength := string(textData), len(textData)
  result := 0
outer:
	for i := 0; i < inputLength-3; i++ {
    if i < inputLength - 6 {
      fmt.Println(inputStr[i:i+7])
    }
    if i < inputLength - 6 && inputStr[i:i+7] == "don't()" {
      for i < inputLength - 3 {
        if i < inputLength - 3 && inputStr[i:i+4] == "do()" {
          break 
        }
        i++
      }
      if i >= inputLength - 3 {
        break outer
      }
      i += 3
      continue outer
    }
		if inputStr[i:i+4] == "mul(" {
			i += 4
      if i >= inputLength {
        break outer
      }
      firstNumBytes := make([]byte, 0, 3)
      for i < inputLength && isDigit(inputStr[i]) {
        firstNumBytes = append(firstNumBytes, inputStr[i])
        if len(firstNumBytes) > 3 {
          continue outer
        }
        i++
      }
      if i >= inputLength {
        break outer
      }
      if inputStr[i] != ',' {
        continue outer
      }
      i++
      firstNum, err := strconv.Atoi(string(firstNumBytes))
      if err != nil {
        panic(err)
      }
      // fmt.Printf("First num %v\n", firstNum)
      if i >= inputLength {
        break outer
      }
      secondNumBytes := make([]byte, 0, 3)
      for i < inputLength && isDigit(inputStr[i]) {
        secondNumBytes = append(secondNumBytes, inputStr[i])
        if len(secondNumBytes) > 3 {
          continue outer
        }
        i++
      }
      if i >= inputLength {
        break outer
      }
      if inputStr[i] != ')' {
        continue outer
      }
      secondNum, err := strconv.Atoi(string(secondNumBytes))
      if err != nil {
        panic(err)
      }
      result += firstNum * secondNum
		}
	}
  fmt.Printf("Day 2 result %v\n", result)
}
func isDigit(char byte) bool {
	return '0' <= char && '9' >= char
}
