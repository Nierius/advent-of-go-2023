package main

import (
	"fmt"
	"log"
	"ourmodule/pkg/utils"
	"strings"
)

const (
	NS    = '|'
	EW    = '-'
	NE    = 'L'
	NW    = 'J'
	SW    = '7'
	SE    = 'F'
	START = 'S'
)

const (
	North = "north"
	East  = "east"
	West  = "west"
	South = "south"
)

type Map struct {
	rows []string
}

func (self Map) getRowAndColumnValue(row, col int) rune {
	if row < 0 || row >= len(self.rows) || col < 0 || col >= len(self.rows[0]) {
		return '.'
	}
	return rune(self.rows[row][col])
}

func main() {
	rows := strings.Split(utils.ReadEntireFileToString("input.txt"), "\n")
	startingRow, startingCol := 0, 0

	for i, row := range rows {
		if !utils.StringIsNotEmpty(row) {
			continue
		}

		for j, tile := range row {
			if tile == 'S' {
				startingRow, startingCol = i, j
			}
		}
	}

	stepsInLoop := stepsInLoop(Map{rows: rows}, startingRow, startingCol)
    fmt.Println("Part 1: ", stepsInLoop / 2)
}

func stepsInLoop(mMap Map, startRow, startCol int) int {
	currentRow, currentCol := startRow, startCol
	nextDirection := getFirstMove(mMap, startRow, startCol)
	stepsTaken := 0
	for {
		// Step
		switch nextDirection {
		case North:
			currentRow--
		case East:
			currentCol++
		case West:
			currentCol--
		case South:
			currentRow++
		}
		stepsTaken++

		// Determine next move
		currentValue := mMap.getRowAndColumnValue(currentRow, currentCol)
		switch currentValue {
		case START:
			return stepsTaken // Loop found
		case NS:
			if nextDirection != South {
				nextDirection = North
			}
		case EW:
			if nextDirection != East {
				nextDirection = West
			}
		case NE:
			if nextDirection == South {
				nextDirection = East
			} else {
				nextDirection = North
			}
		case NW:
			if nextDirection == South {
				nextDirection = West
			} else {
				nextDirection = North
			}
		case SW:
			if nextDirection == North {
				nextDirection = West
			} else {
				nextDirection = South
			}
		case SE:
			if nextDirection == North {
				nextDirection = East
			} else {
				nextDirection = South
			}
		default:
			log.Fatal("Unreachable tile")
		}
	}
}

func getFirstMove(mMap Map, startRow, startCol int) string {
	// Is north possible first move
	northValue := mMap.getRowAndColumnValue(startRow-1, startCol)
	switch northValue {
	case NS:
		return North
	case SW:
		return North
	case SE:
		return North
	}

	// Is east possible first move
	eastValue := mMap.getRowAndColumnValue(startRow, startCol+1)
	switch eastValue {
	case SW:
		return East
	case EW:
		return East
	case NW:
		return East
	}

	// Is west possible first move
	westValue := mMap.getRowAndColumnValue(startRow, startCol-1)
	switch westValue {
	case NE:
		return West
	case EW:
		return West
	case SE:
		return West
	}

	// Is south possible first move
	southValue := mMap.getRowAndColumnValue(startRow+1, startCol)
	switch southValue {
	case NS:
		return South
	case NW:
		return South
	case NE:
		return South
	}

	return ""
}
