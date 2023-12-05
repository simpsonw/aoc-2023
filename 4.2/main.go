package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
	re := regexp.MustCompile(`Card\s+([0-9]+): (.+)`)
	numbersPattern := regexp.MustCompile(`\d+`)
	total := 0
	cardInstances := make(map[int]int, len(lines))
	for _, line := range lines {
		if line == "" {
			break
		}
		subMatches := re.FindStringSubmatch(line)
		if len(subMatches) < 3 {
			log.Fatalf("Malformed input: %q", line)
		}
		cardNumber, _ := strconv.Atoi(subMatches[1])
		gameValues := subMatches[2]
		gameNumbers := strings.Split(gameValues, "|")
		winningNumbers := numbersPattern.FindAllString(gameNumbers[0], -1)
		drawnNumbers := numbersPattern.FindAllString(gameNumbers[1], -1)
		var matchingNumbers []string

		for _, number := range drawnNumbers {
			if contains(winningNumbers, number) {
				matchingNumbers = append(matchingNumbers, number)
			}
		}
		fmt.Printf("Adding original instance of card %d\n", cardNumber)
		fmt.Printf("\tcard %d contained %d matching numbers\n", cardNumber, len(matchingNumbers))
		cardInstances[cardNumber] += 1
		for i:= 1; i <= len(matchingNumbers); i++ {
			if cardInstances[cardNumber+i] == 0 {
				fmt.Printf("\t\tCard %d has no instances, adding 1\n", cardNumber+i)
				cardInstances[cardNumber+i] += cardInstances[cardNumber]
			} else {
				fmt.Printf("\t\tCard %d has %d instances, adding %d instances\n", cardNumber+i, cardInstances[cardNumber+i], cardInstances[cardNumber])
				cardInstances[cardNumber+i] += cardInstances[cardNumber]
			}
		}
	}
	for k, v := range cardInstances {
		fmt.Printf("Card %d has %d instances\n", k, v)
		total += v
	}
	fmt.Printf("Total cards: %d\n", total)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}