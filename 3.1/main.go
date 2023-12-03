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

const UnicodeDigitOffset = 48

func runeToIntValue(digit rune) int {
	return int(digit) - UnicodeDigitOffset
}

func main2(){
	str := "467..114.."
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindAllStringIndex(str, -1))
}

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
	var sum int
	for row, line := range lines {
		if line == "" {
			break
		}
		for column, character := range line {
			if !unicode.IsDigit(character) && character != '.' {
				fmt.Printf("Checking %c (%d, %d):\n", character, row, column)
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
		}
	}

	fmt.Printf("Sum: %d\n", sum)
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