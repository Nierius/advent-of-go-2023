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
	startColumn int
	endColumn   int
}

func main() {
	file, err := os.Open("example_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	marks := []Mark{}
	partNumbers := []PartNumber{}
	rowIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		tparts, tmarks := exportPartNumsAndMarksFromRow(line, rowIndex)
		marks = append(marks, tmarks...)
		partNumbers = append(partNumbers, tparts...)
		rowIndex++
	}

	part1(marks, partNumbers)
	part2(marks, partNumbers)
}

func part1(marks []Mark, partNumbers []PartNumber) {
	total := 0
	for _, num := range partNumbers {
		numOk := false
		for _, mark := range marks {
			if numberAndMarkTouch(num, mark) {
				numOk = true
				break
			}
		}

		if numOk {
			total += num.number
		}
	}

	fmt.Println("Part 1: ", total)
}

func part2(marks []Mark, partNumbers []PartNumber) {
	total := 0
	for _, gear := range utils.Filter(marks, func(mark Mark) bool {
		return mark.symbol == '*'
	}) {
		numbersNextToGear := []int{}
		for _, num := range partNumbers {
			if numberAndMarkTouch(num, gear) {
				numbersNextToGear = append(numbersNextToGear, num.number)
			}
		}

		if len(numbersNextToGear) != 2 {
			continue
		}

		total += numbersNextToGear[0] * numbersNextToGear[1]
	}

	fmt.Println("Part 2: ", total)
}

func numberAndMarkTouch(num PartNumber, mark Mark) bool {
	rowDiff := utils.Abs(num.location.row - mark.location.row)
	if rowDiff > 1 {
		return false
	}

	return num.location.startColumn-1 <= mark.location.startColumn && num.location.endColumn+1 >= mark.location.startColumn
}

func exportPartNumsAndMarksFromRow(row string, rowIndex int) ([]PartNumber, []Mark) {
	partNumbers := []PartNumber{}
	numRe := regexp.MustCompile(`\b\d+\b`)
	numMatches := numRe.FindAllStringSubmatchIndex(row, -1)
	for _, numMatch := range numMatches {
		numStartIndex := numMatch[0]
		numEndIndex := numMatch[1]

		value, _ := strconv.Atoi(row[numStartIndex:numEndIndex])
		partNumbers = append(partNumbers, PartNumber{number: value, location: Location{row: rowIndex, startColumn: numStartIndex, endColumn: numEndIndex - 1}})
	}

	marks := []Mark{}
	markRe := regexp.MustCompile(`[^0-9.]`)
	markMatches := markRe.FindAllStringSubmatchIndex(row, -1)
	for _, markmatch := range markMatches {
		markIndex := markmatch[0]
		symbol := row[markIndex]
		marks = append(marks, Mark{location: Location{row: rowIndex, startColumn: markIndex}, symbol: symbol})
	}

	return partNumbers, marks
}

