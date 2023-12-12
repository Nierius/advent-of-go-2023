package main

import (
	"fmt"
	"log"
	"ourmodule/pkg/utils"
	"strconv"
	"strings"
)

type History = []int

func getDiffSlice(vals []int) []int {
	diffSlice := []int{}
	for i := range vals {
		if i == 0 {
			continue
		}
		diffSlice = append(diffSlice, vals[i]-vals[i-1])
	}

	return diffSlice
}

func main() {
	contents := utils.ReadEntireFileToString("input.txt")
	histories := []History{}
	part1 := 0
	part2 := 0
	for _, line := range strings.Split(contents, "\n") {
		if !utils.StringIsNotEmpty(line) {
			continue
		}

		histories = append(histories, parseHistory(line))
		part1 += extrapolateRight(parseHistory(line))
		part2 += extrapolateLeft(parseHistory(line))
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func extrapolateRight(history History) int {
    diffSlices := getAllDiffSlices(history)

	lastValue := 0
	for i := len(diffSlices) - 1; i >= 0; i-- {
		slice := diffSlices[i]
		lastValue = lastValue + slice[len(slice)-1]
	}

	return lastValue
}

func extrapolateLeft(history History) int {
    diffSlices := getAllDiffSlices(history)

	lastValue := 0
	for i := len(diffSlices) - 1; i >= 0; i-- {
		slice := diffSlices[i]
		lastValue = slice[0] - lastValue
	}

	return lastValue
}

func getAllDiffSlices(history History) [][]int {
	prevDiffSlice := history
	diffSlices := [][]int{history}
	for {
		diffSlice := getDiffSlice(prevDiffSlice)
		if utils.Every(diffSlice, func(val int) bool {
			return val == 0
		}) {
			break
		}

		prevDiffSlice = diffSlice
		diffSlices = append(diffSlices, diffSlice)
	}

	return diffSlices
}

func parseHistory(line string) History {
	history := History{}
	for _, val := range strings.Split(line, " ") {
		if !utils.StringIsNotEmpty(val) {
			continue
		}

		val, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		history = append(history, val)
	}

	return history
}
