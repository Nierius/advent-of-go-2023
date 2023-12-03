package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
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

	total := 0
    for _, gear := range marks {
        numbersnexttogear := []int{}
        for _, num := range partnumbers {
			rowdiff := abs(num.location.row - gear.location.row)
			if rowdiff > 1 {
				continue
			}

			if num.location.startcolumn-1 <= gear.location.startcolumn && num.location.endcolumn+1 >= gear.location.startcolumn {
                numbersnexttogear = append(numbersnexttogear, num.number)
			}
        }

        if len(numbersnexttogear) != 2 {
            continue
        }

        total += numbersnexttogear[0] * numbersnexttogear[1]
    }
	fmt.Println(total)
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
	markre := regexp.MustCompile(`\*`)
	markmatches := markre.FindAllStringSubmatchIndex(row, -1)
	for _, markmatch := range markmatches {
		markindex := markmatch[0]
		marks = append(marks, Mark{location: Location{row: rowindex, startcolumn: markindex}})
	}

	return partnumbers, marks
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

type PartNumber struct {
	number   int
	location Location
}

type Mark struct {
	location Location
}

type Location struct {
	row         int
	startcolumn int
	endcolumn   int
}
