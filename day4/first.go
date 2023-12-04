package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./example_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	games := []Game{}
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, parseGame(line))
	}

	total := 0
	for _, game := range games {
		total += gameToPoints(game)
	}

	fmt.Println(total)
}

func gameToPoints(game Game) int {
	matches := 0
	for _, selectedNum := range game.selectedNums {
		match := false
		for _, winningNum := range game.winningNums {
			if selectedNum == winningNum {
				match = true
				break
			}
		}

		if match {
			matches++
		}
	}

    if matches < 1 {
        return 0
    }

	return int(math.Pow(2.0, float64(matches-1)))
}

func parseGame(row string) Game {
	gameAndHeaderSplit := strings.Split(row, ": ")
	numsSplit := strings.Split(gameAndHeaderSplit[len(gameAndHeaderSplit)-1], " | ")

	winningNums := []int{}
	winningNumsStr := strings.Split(numsSplit[0], " ")
	for _, numStr := range winningNumsStr {
		num, err := strconv.Atoi(numStr)
        if err != nil {
            continue
        }
		winningNums = append(winningNums, num)
	}

	selectedNums := []int{}
	selectedNumsStr := strings.Split(numsSplit[len(numsSplit)-1], " ")
	for _, numStr := range selectedNumsStr {
		num, err := strconv.Atoi(numStr)
        if err != nil {
            continue
        }
		selectedNums = append(selectedNums, num)
	}

	return Game{
        winningNums: winningNums,
        selectedNums: selectedNums,
    }
}

type Game struct {
	winningNums  []int
	selectedNums []int
}
