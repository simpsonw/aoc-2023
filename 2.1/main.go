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
		possible := false
		for _, v := range showings {
			possible, err = gamePossible(v)
			if err != nil {
				log.Fatal(err)
			}

			if !possible {
				fmt.Printf("Game %d: %s is not possible (MaxRed: %d, MaxGreen: %d, MaxBlue: %d)\n", number, v, MaxRed, MaxGreen, MaxBlue)
				break
			}
		}
		if possible {
			sum += number
		}
	}
	fmt.Printf("The sum was %d\n", sum)
}

func gamePossible(game string) (bool, error){
	cubes := strings.Split(game, ", ")
	possible := true
	for _, round := range cubes {
		components := strings.Split(round, " ")
		if len(components) != 2 {
			return false, fmt.Errorf("Malformed components: %q", components)
		}
		number, err := strconv.Atoi(components[0])
		if err != nil {
			return false, fmt.Errorf("Malformed input: %q", game)
		}
		color := components[1]
		switch color {
		case "red":
			possible = number <= MaxRed
		case "blue":
			possible = number <= MaxBlue
		case "green":
			possible = number <= MaxGreen
		default:
			possible = false
		}
		if !possible {
			return possible, nil
		}
	}
	return possible, nil
}
