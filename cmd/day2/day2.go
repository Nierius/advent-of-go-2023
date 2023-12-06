package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Round struct {
	blues  int
	greens int
	reds   int
}

type Game struct {
	id     int
	rounds []Round
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	games := []Game{}
	for scanner.Scan() {
		line := scanner.Text()
		game := lineIntoGame(line)
		games = append(games, game)
	}

	part1(games)
	part2(games)
}

func part1(games []Game) {
	total := 0
	for _, game := range games {
		gameOk := true
		for _, round := range game.rounds {
			if round.blues > 14 {
				gameOk = false
				break
			}
			if round.greens > 13 {
				gameOk = false
				break
			}
			if round.reds > 12 {
				gameOk = false
				break
			}
		}

		if !gameOk {
			continue
		}

		total += game.id
	}

	fmt.Println("Part 1: ", total)
}

func part2(games []Game) {
	total := 0
	for _, game := range games {
		gameMinBlues := 0
		gameMinReds := 0
		gameMinGreens := 0
		for _, round := range game.rounds {
			if round.blues > gameMinBlues {
				gameMinBlues = round.blues
			}
			if round.greens > gameMinGreens {
				gameMinGreens = round.greens
			}
			if round.reds > gameMinReds {
				gameMinReds = round.reds
			}
		}

		total += gameMinBlues * gameMinReds * gameMinGreens
	}

	fmt.Println("Part 2: ", total)
}

func lineIntoGame(line string) Game {
	game := Game{}

	splits := strings.Split(line, ";")

	idRe := regexp.MustCompile(`Game (\d+)`)
	redRe := regexp.MustCompile(`(\d+) red`)
	blueRe := regexp.MustCompile(`(\d+) blue`)
	greenRe := regexp.MustCompile(`(\d+) green`)

	for _, split := range splits {
		round := Round{}
		idMatches := idRe.FindStringSubmatch(split)
		if len(idMatches) > 0 {
			game.id, _ = strconv.Atoi(idMatches[len(idMatches)-1])
		}

		redMatches := redRe.FindStringSubmatch(split)
		if len(redMatches) > 0 {
			round.reds, _ = strconv.Atoi(redMatches[len(redMatches)-1])
		}

		blueMatches := blueRe.FindStringSubmatch(split)
		if len(blueMatches) > 0 {
			round.blues, _ = strconv.Atoi(blueMatches[len(blueMatches)-1])
		}

		greenMatches := greenRe.FindStringSubmatch(split)
		if len(greenMatches) > 0 {
			round.greens, _ = strconv.Atoi(greenMatches[len(greenMatches)-1])
		}

		game.rounds = append(game.rounds, round)
	}

	return game
}
