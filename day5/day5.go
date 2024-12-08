package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	textInput, err := os.ReadFile("./day5.input")
	if err != nil {
		panic(err)
	}
	splitSections := bytes.Split(textInput, []byte{'\n', '\n'})
	orderingRules, updates := splitSections[0], splitSections[1]
	mustComeBeforeMap := make(map[string][]string)
	splitOrderingRules := bytes.Split(orderingRules, []byte{'\n'})
	for _, orderingRule := range splitOrderingRules {
		splitRule := bytes.Split(orderingRule, []byte{'|'})
		if len(splitRule) == 0 {
			continue
		}
		pageBefore, pageAfter := splitRule[0], splitRule[1]
		if curSlice, ok := mustComeBeforeMap[string(pageBefore)]; ok {
			mustComeBeforeMap[string(pageBefore)] = append(curSlice, string(pageAfter))
		} else {
			mustComeBeforeMap[string(pageBefore)] = []string{string(pageAfter)}
		}
	}
	// fmt.Println(mustComeBeforeMap)
	splitUpdates := bytes.Split(updates, []byte{'\n'})
	result, validUpdates := 0, 0
	part2Res, invalidUpdates := 0, 0
	for _, update := range splitUpdates {
		if len(update) == 0 {
			continue
		}
		splitPageNumbers := bytes.Split(update, []byte{','})
		validUpdate := true
		seenPages := make(map[string]bool)
	currentUpdateValidation:
		for _, pageNumber := range splitPageNumbers {
			mustComeBeforeSlice, _ := mustComeBeforeMap[string(pageNumber)]
			for _, mustComeBeforePage := range mustComeBeforeSlice {
				if _, ok := seenPages[mustComeBeforePage]; ok { //If we have seen any of the pages this page must come before already, this update is invalid
					validUpdate = false
					break currentUpdateValidation
				}
			}
			seenPages[string(pageNumber)] = true
		}
		if validUpdate {
			middleIndex := (len(splitPageNumbers) - 1) / 2
			middlePageNumber, err := strconv.Atoi(string(splitPageNumbers[middleIndex]))
			if err != nil {
				panic(err)
			}
			validUpdates++
			result += middlePageNumber
		} else {
			slices.SortFunc(splitPageNumbers, func(a, b []byte) int {
				aBeforeSlice, _ := mustComeBeforeMap[string(a)]
				bBeforeSlice, _ := mustComeBeforeMap[string(b)]
				if slices.Contains(aBeforeSlice, string(b)) {
					return -1
				} else if slices.Contains(bBeforeSlice, string(a)) {
					return 1
				}
				return 0
			})
			middleIndex := (len(splitPageNumbers) - 1) / 2
			middlePageNumber, err := strconv.Atoi(string(splitPageNumbers[middleIndex]))
			if err != nil {
				panic(err)
			}
			invalidUpdates++
			part2Res += middlePageNumber
		}
	}
	fmt.Printf("Part 1 had %v valid updates result %v\n", validUpdates, result)
	fmt.Printf("Part 2 had %v invalid page numbers result %v", invalidUpdates, part2Res)
}
