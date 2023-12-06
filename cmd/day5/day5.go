package main

import (
	"fmt"
	"log"
	"math"
	"ourmodule/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

type Seed struct {
	id int
}

type SeedRange struct {
	rangeStart int
	rangeStop  int
}

type MapRange struct {
	sourceStart int
	targetStart int
	mRange      int
}

type Map struct {
	ranges []MapRange
}

func (self Map) sourceToTarget(source int) int {
	for _, rrange := range self.ranges {
		if rrange.includes(source) {
			return rrange.sourceToTarget(source)
		}
	}

	return source
}

func (self MapRange) includes(source int) bool {
	diff := source - self.sourceStart
	return diff >= 0 && diff < self.mRange
}

func (self MapRange) sourceToTarget(source int) int {
	diff := source - self.sourceStart
	return self.targetStart + diff
}


func main() {
	contents := utils.ReadEntireFileToString("example_input")
	conversionMaps := strings.Split(contents, "\n\n")

	seedsP1 := parseSeedsPart1(conversionMaps[0])
	seedsP2 := parseSeedsPart2(conversionMaps[0])

	maps := []Map{}
	for _, c := range conversionMaps[1:] {
		maps = append(maps, parseMap(c))
	}

	// Seed to loc
	part1(seedsP1, maps)
	part2(seedsP2, maps)
}

func part1(seeds []Seed, maps []Map) {
	smallestLocSoFar := math.MaxInt
	for _, seed := range seeds {
		loc := getSeedLoc(seed.id, maps)
		smallestLocSoFar = min(smallestLocSoFar, loc)
	}

	fmt.Println("Part 1: ", smallestLocSoFar)
}

func part2(seedRanges []SeedRange, maps []Map) {
	smallestLocSoFar := math.MaxInt
	for _, seedRange := range seedRanges {
		for i := seedRange.rangeStart; i < seedRange.rangeStop; i++ {
			loc := getSeedLoc(i, maps)
			smallestLocSoFar = min(smallestLocSoFar, loc)
		}
	}

	fmt.Println("Part 2: ", smallestLocSoFar)
}

func getSeedLoc(seedId int, maps []Map) int {
	currentPremilinaryValue := seedId
	for _, mMap := range maps {
		currentPremilinaryValue = mMap.sourceToTarget(currentPremilinaryValue)
	}

	return currentPremilinaryValue
}

func getSmallestLoc(seeds []Seed, maps []Map) int {
	results := []int{}
	for _, seed := range seeds {
		currentId := seed.id
		for _, mMap := range maps {
			currentId = mMap.sourceToTarget(currentId)
		}

		results = append(results, currentId)
	}

	sort.Ints(results)
	return results[0]
}

func parseSeedsPart1(line string) []Seed {
	seeds := []Seed{}
	seedSplit := strings.Split(line, " ")[1:]
	for _, s := range seedSplit {
		sId, e := strconv.Atoi(s)
		if e != nil {
			log.Fatal(e)
		}

		seeds = append(seeds, Seed{id: sId})
	}
	return seeds
}

func parseSeedsPart2(line string) []SeedRange {
	seedRanges := []SeedRange{}
	seedRangeParts := parseSeedsPart1(line)
	for i, s := range seedRangeParts {
		// Every second marks start of range. 0, 2, 4...
		if i%2 != 0 {
			continue
		}

		seedRanges = append(seedRanges, SeedRange{rangeStart: s.id, rangeStop: s.id + seedRangeParts[i+1].id - 1})
	}

	return seedRanges
}

func parseMap(lines string) Map {
	mapSplit := strings.Split(lines, "\n")[1:]
	ranges := []MapRange{}

	for _, rangeRaw := range mapSplit {
		if rangeRaw == "" {
			continue
		}
		ranges = append(ranges, parseRange(rangeRaw))
	}

	return Map{ranges: ranges}
}

func parseRange(line string) MapRange {
	vals := strings.Split(line, " ")
	convertedVals := []int{}

	for _, val := range vals {
		if val == "" {
			continue
		}

		converted, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		convertedVals = append(convertedVals, converted)
	}

	return MapRange{sourceStart: convertedVals[1], targetStart: convertedVals[0], mRange: convertedVals[2]}
}
