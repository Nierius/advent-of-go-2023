package main

import (
	"fmt"
	"ourmodule/pkg/utils"
	"strings"
)

type Map struct {
	rows []string
}

type Galaxy struct {
	row, col int
}

func main() {
	mMap := Map{rows: strings.Split(utils.ReadEntireFileToString("input.txt"), "\n")}
	expanded := expandRows(mMap)
	expanded = expandCols(expanded)
	galaxies := exportGalaxies(expanded)

	sumOfDistances := 0
    for i, galaxy := range galaxies {
        for _, oGalaxy := range galaxies[i+1:] {
            sumOfDistances += utils.Abs(galaxy.col - oGalaxy.col)
            sumOfDistances += utils.Abs(galaxy.row - oGalaxy.row)
        }
    }

	fmt.Println("Part 1: ", sumOfDistances)
}

func exportGalaxies(inp Map) []Galaxy {
	galaxies := []Galaxy{}
	for i, row := range inp.rows {
		for j, chr := range row {
			if chr == '#' {
				galaxies = append(galaxies, Galaxy{row: i, col: j})
			}
		}
	}

	return galaxies
}

func expandCols(inp Map) Map {
	expandedMap := Map{rows: make([]string, len(inp.rows))}
	for i := range inp.rows[0] {
		emptyCol := true
		for _, row := range inp.rows {
			if row[i] == '#' {
				emptyCol = false
				break
			}
		}

		for j, row := range inp.rows {
			if len(expandedMap.rows) <= j {
				expandedMap.rows = append(expandedMap.rows, "")
			}
			if emptyCol {
				expandedMap.rows[j] += ".."
			} else {
				expandedMap.rows[j] += string(row[i])
			}
		}
	}

	return expandedMap
}

func expandRows(inp Map) Map {
	expandedMap := Map{}
	for _, row := range inp.rows {
		if !utils.StringIsNotEmpty(row) {
			continue
		}
		expandedMap.rows = append(expandedMap.rows, row)
		isEmptyRow := true
		for _, chr := range row {
			if chr == '#' {
				isEmptyRow = false
				break
			}
		}

		if isEmptyRow {
			expandedMap.rows = append(expandedMap.rows, row)
		}
	}

	return expandedMap
}
