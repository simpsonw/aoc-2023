package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
	"strconv"
	"sort"
)

type Mapping struct{
	startSource int
	startDest int
	mapRange int
}

func (m Mapping) String() string{
	return fmt.Sprintf("%d %d %d", m.startDest, m.startSource, m.mapRange)
}

type Map struct{
	name string
	mappings []*Mapping
}

func (m Map) String() string{
	return m.name
}

func (m *Map) addMapping(startSource, startDest, mapRange int)  {
	mapping := &Mapping{
		startSource: startSource,
		startDest:   startDest,
		mapRange:    mapRange,
	}
	m.mappings = append(m.mappings, mapping)
}

func (m *Map) getDestValue(sourceValue int)  int {
	for _, mapping := range m.mappings {
		if sourceValue >= mapping.startSource && sourceValue <= mapping.startSource + mapping.mapRange {
			return  mapping.startDest - mapping.startSource + sourceValue
		}
	}
	return sourceValue
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
	seedsPattern := regexp.MustCompile(`seeds:\s(.+)`)
	mapPattern := regexp.MustCompile(`([a-z]+-to-[a-z]+)(\smap:)`)
	var seeds []string
	var maps []*Map
	inMap := false
	for _, line := range lines {
		if line == "" {
			inMap = false
			continue
		}
		if !inMap {
			if seedsPattern.MatchString(line) {
				submatches := seedsPattern.FindStringSubmatch(line)
				seeds = strings.Split(submatches[1], " ")
			} else if mapPattern.MatchString(line){
				submatches := mapPattern.FindAllStringSubmatch(line, -1)
				maps = append(maps, &Map{name: submatches[0][1]})
				inMap = true
			}
		} else {
			mappingValues := strings.Split(line, " ")
			source, _ := strconv.Atoi(mappingValues[1])
			dest, _ := strconv.Atoi(mappingValues[0])
			mapRange, _ := strconv.Atoi(mappingValues[2])
			maps[len(maps)-1].addMapping(source, dest, mapRange)
		}
	}
	var locations []int
	for _, seed := range seeds {
		sourceValue, _ := strconv.Atoi(seed)
		for _, m := range maps {
			destValue := m.getDestValue(sourceValue)
			sourceValue = destValue
		}
		locations = append(locations, sourceValue)
	}
	sort.Ints(locations)
	fmt.Printf("Part 1: The lowest location is %d\n", locations[0])

	var lowestLocation int
	for i := 0; i < len(seeds); i+=2 {
		startValue, _ := strconv.Atoi(seeds[i])
		seedRange, _ := strconv.Atoi(seeds[i+1])
		for j := 0; j < seedRange; j++ {
			sourceValue := startValue+j
			for _, m := range maps {
				destValue := m.getDestValue(sourceValue)
				sourceValue = destValue
			}
			if lowestLocation == 0 {
				lowestLocation = sourceValue
			} else if sourceValue < lowestLocation {
				fmt.Printf("source value: %d lowest location: %d\n", sourceValue, lowestLocation)
				lowestLocation = sourceValue
			}
		}
	}

	fmt.Printf("Part 2: The lowest location is %d\n", lowestLocation)

}