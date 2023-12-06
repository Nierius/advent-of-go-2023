package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ourmodule/pkg/utils"
	"regexp"
	"strconv"
)

type PartNumber struct {
	number   int
	location Location
}

type Mark struct {
	location Location
	symbol   byte
}

type Location struct {
	row         int
	startcolumn int
	endcolumn   int
}

func main() {
	file, err := os.Open("example_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	marks := []Mark{}
	partnumbers := []PartNumber{}
	rowindex := 0
	for scanner.Scan() {
		line := scanner.Text()
		tparts, tmarks := exportPartNumsAndMarksFromRow(line, rowindex)
		marks = append(marks, tmarks...)
		partnumbers = append(partnumbers, tparts...)
		rowindex++
	}

	part1(marks, partnumbers)
	part2(marks, partnumbers)
}

func part1(marks []Mark, partnumbers []PartNumber) {
	total := 0
	for _, num := range partnumbers {
		numok := false
		for _, mark := range marks {
            if numberAndMarkTouch(num, mark) {
                numok = true
                break
            }
		}

		if numok {
			total += num.number
		}
	}

	fmt.Println("Part 1: ", total)
}

func part2(marks []Mark, partnumbers []PartNumber) {
	total := 0
	for _, gear := range utils.Filter(marks, func(mark Mark) bool {
		return mark.symbol == '*'
	}) {
		numbersnexttogear := []int{}
		for _, num := range partnumbers {
			if numberAndMarkTouch(num, gear) {
				numbersnexttogear = append(numbersnexttogear, num.number)
			}
		}

		if len(numbersnexttogear) != 2 {
			continue
		}

		total += numbersnexttogear[0] * numbersnexttogear[1]
	}

	fmt.Println("Part 2: ", total)
}

func numberAndMarkTouch(num PartNumber, mark Mark) bool {
	rowdiff := abs(num.location.row - mark.location.row)
	if rowdiff > 1 {
		return false
	}

	return num.location.startcolumn-1 <= mark.location.startcolumn && num.location.endcolumn+1 >= mark.location.startcolumn
}

func exportPartNumsAndMarksFromRow(row string, rowindex int) ([]PartNumber, []Mark) {
	partnumbers := []PartNumber{}
	numre := regexp.MustCompile(`\b\d+\b`)
	nummatches := numre.FindAllStringSubmatchIndex(row, -1)
	for _, nummatch := range nummatches {
		numstartindex := nummatch[0]
		numendindex := nummatch[1]

		value, _ := strconv.Atoi(row[numstartindex:numendindex])
		partnumbers = append(partnumbers, PartNumber{number: value, location: Location{row: rowindex, startcolumn: numstartindex, endcolumn: numendindex - 1}})
	}

	marks := []Mark{}
	markre := regexp.MustCompile(`[^0-9.]`)
	markmatches := markre.FindAllStringSubmatchIndex(row, -1)
	for _, markmatch := range markmatches {
		markindex := markmatch[0]
		symbol := row[markindex]
		marks = append(marks, Mark{location: Location{row: rowindex, startcolumn: markindex}, symbol: symbol})
	}

	return partnumbers, marks
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
