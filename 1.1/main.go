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
		for _, c := range l {
			if unicode.IsDigit(c) {
				if firstDigit == 0 {
					firstDigit = int(c)
					lastDigit = int(c)
				} else {
					lastDigit = int(c)
				}
			}
		}
		number, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit-48, lastDigit-48))
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Printf("The sum was %d\n", sum)
}
