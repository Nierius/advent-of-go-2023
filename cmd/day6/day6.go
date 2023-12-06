package main

import (
	"fmt"
	"log"
	"ourmodule/pkg/utils"
	"strconv"
	"strings"
)

func main() {
	contentsSplitted := strings.Split(utils.ReadEntireFileToString("input.txt"), "\n")
	racesPart1 := parseRacesPart1(contentsSplitted)
	racePart2 := parseRacePart2(contentsSplitted)

	part1(racesPart1)
    part2(racePart2)
}

func part1(races []Race) {
	winningChargeTimes := [][]int{}
	for _, race := range races {
		winningChargeTimes = append(winningChargeTimes, exportRecordChargeTimesForRace(race))
	}

	product := 1
	for _, winCharges := range winningChargeTimes {
		product *= len(winCharges)
	}

	fmt.Println("Part 1: ", product)
}

func part2(race Race) {
    fmt.Println("Part 2: ", len(exportRecordChargeTimesForRace(race)))
}

func parseRacesPart1(input []string) []Race {
	times := utils.Filter(strings.Split(input[0], " ")[1:], utils.StringIsNotEmpty)
	records := utils.Filter(strings.Split(input[1], " ")[1:], utils.StringIsNotEmpty)

	races := []Race{}
	for i := range times {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatal(err)
		}
		record, err := strconv.Atoi(records[i])
		if err != nil {
			log.Fatal(err)
		}

		races = append(races, Race{time: time, record: record})
	}

	return races
}

func parseRacePart2(input []string) Race {
	timeRaw := strings.Join(utils.Filter(strings.Split(strings.Split(input[0], ":")[1], " "), utils.StringIsNotEmpty), "")
	recordRaw := strings.Join(utils.Filter(strings.Split(strings.Split(input[1], ":")[1], " "), utils.StringIsNotEmpty), "")

	time, err := strconv.Atoi(timeRaw)
	if err != nil {
		log.Fatal(err)
	}
	record, err := strconv.Atoi(recordRaw)
	if err != nil {
		log.Fatal(err)
	}

	return Race{time: time, record: record}
}

func exportRecordChargeTimesForRace(race Race) []int {
	winningChargeTimes := []int{}

	// Data set is so small so bruteforce solution is good enough
	for i := 1; i < race.time; i++ {
		distance := (race.time - i) * i
		if distance > race.record {
			winningChargeTimes = append(winningChargeTimes, i)
		}
	}

	return winningChargeTimes
}

type Race struct {
	time   int
	record int
}
