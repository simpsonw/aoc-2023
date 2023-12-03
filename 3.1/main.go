package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatalf("Usage: %s <filename> \n", os.Args[0])
	}
	filename := os.Args[1]
	var content, err = os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var sum, gearRatioSum int
	for row, line := range lines {
		if line == "" {
			break
		}
		for column, character := range line {
			if !unicode.IsDigit(character) && character != '.' {
				// Above
				if row > 0 {
					sum += addNumbersForSymbol(lines[row-1], column)
				}
				//Below
				if row < len(lines) {
					sum += addNumbersForSymbol(lines[row+1], column)
				}
				// Current line
				sum += addNumbersForSymbol(line, column)
			}

			if character == '*' {
				var adjacentNumbers []int
				// Above
				if row > 0 {
					adjacentNumbers = append(adjacentNumbers, getAdjacentNumbers(lines[row-1], column)...)
				}
				//Below
				if row < len(lines) {
					adjacentNumbers = append(adjacentNumbers, getAdjacentNumbers(lines[row+1], column)...)
				}

				// Current line
				adjacentNumbers = append(adjacentNumbers, getAdjacentNumbers(line, column)...)
				if len(adjacentNumbers) == 2 {
					gearRatioSum += adjacentNumbers[0]*adjacentNumbers[1]
				}
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Gear Ratio Sum: %d\n", gearRatioSum)
}

func getAdjacentNumbers(line string, column int) (adjacentNumbers []int) {
	numbersPattern := regexp.MustCompile(`\d+`)
	numberIndicies := numbersPattern.FindAllStringIndex(line, -1)
	for _, v := range numberIndicies {
		if numberIsAdjacent(v, column){
			number, _ := strconv.Atoi(line[v[0]:v[1]])
			adjacentNumbers = append(adjacentNumbers, number)
		}
	}
	return
}

func addNumbersForSymbol(line string, column int) int {
	numbersPattern := regexp.MustCompile(`\d+`)
	var sum int
	numberIndicies := numbersPattern.FindAllStringIndex(line, -1)
	for _, v := range numberIndicies {
		if numberIsAdjacent(v, column){
			number, _ := strconv.Atoi(line[v[0]:v[1]])
			sum += number
		}
	}
	return sum
}

func numberIsAdjacent(indicies []int, column int) bool {
	return indicies[0] < column && indicies[1] >= column ||
		indicies[0] == column || indicies[0] == column+1
}