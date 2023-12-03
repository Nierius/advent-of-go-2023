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

    total := 0
    for _, game := range games {
        gameok := true
        for _, round := range game.rounds {
            if round.blues > 14 {
                gameok = false
                break
            }
            if round.greens > 13 {
                gameok = false
                break
            }
            if round.reds > 12 {
                gameok = false
                break
            }
        }

        if !gameok {
            continue
        }

        total += game.id
    }

    fmt.Println(total)
}

type Round struct {
	blues  int
	greens int
	reds   int
}

type Game struct {
	id     int
	rounds []Round
}

func lineIntoGame(line string) Game {
	game := Game{}

	splits := strings.Split(line, ";")

	idre := regexp.MustCompile(`Game (\d+)`)
	redre := regexp.MustCompile(`(\d+) red`)
	bluere := regexp.MustCompile(`(\d+) blue`)
	greenre := regexp.MustCompile(`(\d+) green`)

	for _, split := range splits {
		fmt.Println("Row", split)
		round := Round{}
		idmatches := idre.FindStringSubmatch(split)
		if len(idmatches) > 0 {
			game.id, _ = strconv.Atoi(idmatches[len(idmatches)-1])
		}

		redmatches := redre.FindStringSubmatch(split)
		if len(redmatches) > 0 {
			round.reds, _ = strconv.Atoi(redmatches[len(redmatches)-1])
		}

		bluematches := bluere.FindStringSubmatch(split)
		if len(bluematches) > 0 {
			round.blues, _ = strconv.Atoi(bluematches[len(bluematches)-1])
		}

		greenmatches := greenre.FindStringSubmatch(split)
		if len(greenmatches) > 0 {
			round.greens, _ = strconv.Atoi(greenmatches[len(greenmatches)-1])
		}

		game.rounds = append(game.rounds, round)
	}

	return game
}
