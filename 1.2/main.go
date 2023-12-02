package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	sum := 0
	for _, l := range lines {
		if l == "" {
			break
		}
		firstDigit := 0
		lastDigit := 0
		for k, _ := range l {
			digit, err := getDigit(l, k)
			if err == nil {
				if firstDigit == 0 {
					firstDigit = digit
					lastDigit = digit
				} else {
					lastDigit = digit
				}
			}
		}
		numberString := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		fmt.Printf("Number was %s (%q)\n", numberString, l)
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Printf("The sum was %d\n", sum)
}

func getDigit(line string, index int) (int, error) {
	r := rune(line[index])
	if unicode.IsDigit(r) {
		return int(r) - 48, nil
	} else {
		if isWordDigit(line, "one", index) {
			return 1, nil
		}
		if isWordDigit(line, "two", index) {
			return 2, nil
		}
		if isWordDigit(line, "three", index) {
			return 3, nil
		}
		if isWordDigit(line, "four", index) {
			return 4, nil
		}
		if isWordDigit(line, "five", index) {
			return 5, nil
		}
		if isWordDigit(line, "six", index) {
			return 6, nil
		}
		if isWordDigit(line, "seven", index) {
			return 7, nil
		}
		if isWordDigit(line, "eight", index) {
			return 8, nil
		}
		if isWordDigit(line, "nine", index) {
			return 9, nil
		}
	}
	return 0, fmt.Errorf("No digit found at index %d in %q", index, line)
}

func isWordDigit(line, wordDigit string, index int) bool {
	isWordDigit := false
	isLongEnough := len(line[index:]) >= len(wordDigit)
	if isLongEnough {
		candidateWord := line[index : index+len(wordDigit)]
		isWordDigit = isLongEnough && candidateWord == wordDigit
	}

	return isWordDigit
}
