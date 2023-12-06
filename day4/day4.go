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

type Game struct {
	winningNums  []int
	selectedNums []int
}

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

	part1(games)
	part2(games)
}

func part1(games []Game) {
	total := 0
	for _, game := range games {
		total += gameToPoints(game)
	}

	fmt.Println("Part 1: ", total)
}

func part2(games []Game) {
	cards := make([]int, len(games))
	for i := 0; i < len(games); i++ {
		cards = append(cards, 0)
	}

	total := 0
	for i, game := range games {
		cards[i]++
		matches := gameToMatches(game)
		for j := 1; j <= matches; j++ {
			cards[i+j] += 1 * cards[i]
		}

		total += cards[i]
	}

	fmt.Println("Part 2: ", total)
}

func gameToPoints(game Game) int {
	matches := gameToMatches(game)

	if matches < 1 {
		return 0
	}

	return int(math.Pow(2.0, float64(matches-1)))
}

func gameToMatches(game Game) int {
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

	return matches
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
		winningNums:  winningNums,
		selectedNums: selectedNums,
	}
}
