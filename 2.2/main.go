package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"regexp"
)

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

func main2(){
	re := regexp.MustCompile(`Game ([0-9]+): (.+)`)
	str := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	fmt.Printf("%#v\n", re.FindStringSubmatch(str))

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
	sum := 0
	re := regexp.MustCompile(`Game ([0-9]+): (.+)`)
	for _, l := range lines {
		if l == "" {
			break
		}
		subMatches := re.FindStringSubmatch(l)
		if subMatches == nil  || len(subMatches) != 3{
			log.Fatalf("Malformed line input: %q", subMatches)
		}
		number, err := strconv.Atoi(subMatches[1])
		if err != nil {
			log.Fatal(err)
		}
		showings := strings.Split(subMatches[2], "; ")
		var minRed, minGreen, minBlue int
		for _, v := range showings {
			cubes := strings.Split(v, ", ")
			for _, round := range cubes {
				components := strings.Split(round, " ")
				color := components[1]
				numCubes, _ := strconv.Atoi(components[0])
				switch color {
				case "red":
					if numCubes > minRed {
						minRed = numCubes
					}
				case "blue":
					if numCubes > minBlue {
						minBlue = numCubes
					}
				case "green":
					if numCubes > minGreen {
						minGreen = numCubes
					}
				default:
					log.Fatalf("Unknown color: %q", color)
				}
			}
		}
		power := minRed*minGreen*minBlue
		fmt.Printf("Game %d: min red %d, min blue %d, min green %d, power %d\n", number, minRed, minBlue, minGreen, power)
		sum += minBlue*minGreen*minRed
	}
	fmt.Printf("The sum was %d\n", sum)
}