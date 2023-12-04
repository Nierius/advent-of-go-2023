package main

import (
	"bufio"
	"fmt"
	"log"
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
	cards := make([]int, len(games))
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, parseGame(line))
	}

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

    fmt.Println(cards)
	fmt.Println(total)
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

type Game struct {
	winningNums  []int
	selectedNums []int
}
