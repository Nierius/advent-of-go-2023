package main

import (
	"fmt"
	"log"
	"ourmodule/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

const JOKER = 1

type Hand struct {
	cardValues []int
	bid        int
}

func (self Hand) isFiveOfKind() bool {
	valueFreqs := self.getValueFrequencies()
	return len(utils.MapFilter(valueFreqs, func(key int, val int) bool { return val > 4 })) > 0
}

func (self Hand) isFourOfKind() bool {
	valueFreqs := self.getValueFrequencies()
	return len(utils.MapFilter(valueFreqs, func(key int, val int) bool { return val > 3 })) > 0
}

func (self Hand) isFullHouse() bool {
	valueFreqs := self.getValueFrequencies()
	return utils.MapEvery(valueFreqs, func(key int, val int) bool { return val == 3 || val == 2 || key == JOKER })
}

func (self Hand) isThreeOfKind() bool {
	valueFreqs := self.getValueFrequencies()
	return len(utils.MapFilter(valueFreqs, func(key int, val int) bool { return val > 2 })) > 0
}

func (self Hand) isTwoPair() bool {
	valueFreqs := self.getValueFrequencies()
	return len(utils.MapFilter(valueFreqs, func(key int, val int) bool { return val > 1 })) > 1
}

func (self Hand) isOnePair() bool {
	valueFreqs := self.getValueFrequencies()
	return len(utils.MapFilter(valueFreqs, func(key int, val int) bool { return val > 1 })) > 0
}

func (self Hand) getValueFrequencies() map[int]int {
	valueFrequencies := make(map[int]int)
	for _, val := range self.cardValues {
		_, exists := valueFrequencies[val]
		if exists {
			valueFrequencies[val]++
			continue
		}

		valueFrequencies[val] = 1
	}

	jokersAmount, ok := valueFrequencies[JOKER]
	if !ok {
		jokersAmount = 0
	}

	// Add joker to the one there is most of
	if jokersAmount > 0 {
		mostOfKey := 0
		mostOfVal := 0
		for key, val := range valueFrequencies {
			if key == JOKER {
				continue
			}

			if val > mostOfVal {
				mostOfVal = val
				mostOfKey = key
			}
		}
		valueFrequencies[mostOfKey] += jokersAmount
	}

	return valueFrequencies
}

type SortByValue []Hand

func (a SortByValue) Len() int           { return len(a) }
func (a SortByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByValue) Less(i, j int) bool { return compareHands(a[i], a[j]) < 0 }

func main() {
	rawHands := strings.Split(utils.ReadEntireFileToString("example_input"), "\n")

	handsNoJoker := []Hand{}
	for _, hand := range rawHands {
		if utils.StringIsNotEmpty(hand) {
			handsNoJoker = append(handsNoJoker, parseHand(hand, false))
		}
	}

	handsJoker := []Hand{}
	for _, hand := range rawHands {
		if utils.StringIsNotEmpty(hand) {
			handsJoker = append(handsJoker, parseHand(hand, true))
		}
	}

	sort.Sort(SortByValue(handsNoJoker))
	sort.Sort(SortByValue(handsJoker))
	part1(handsNoJoker)
	part2(handsJoker)
}

func part1(sortedHands []Hand) {
	total := 0
	for i, hand := range sortedHands {
		total += (i + 1) * hand.bid
	}

	fmt.Println("Part 1: ", total)
}

func part2(sortedHands []Hand) {
	total := 0
	for i, hand := range sortedHands {
		total += (i + 1) * hand.bid
	}

	fmt.Println("Part 2: ", total)
}

func parseHand(raw string, useJokers bool) Hand {
	splittedToCardsAndBid := strings.Split(raw, " ")
	bid, err := strconv.Atoi(splittedToCardsAndBid[1])
	if err != nil {
		log.Fatal(err)
	}

	cardValues := []int{}
	for _, rawCardValue := range splittedToCardsAndBid[0] {
		switch rawCardValue {
		case 'A':
			cardValues = append(cardValues, 14)
		case 'K':
			cardValues = append(cardValues, 13)
		case 'Q':
			cardValues = append(cardValues, 12)
		case 'J':
			if useJokers {
				cardValues = append(cardValues, JOKER)
				continue
			}

			cardValues = append(cardValues, 11)
		case 'T':
			cardValues = append(cardValues, 10)

		default:
			val, err := strconv.Atoi(string(rawCardValue))
			if err != nil {
				log.Fatal(err)
			}

			cardValues = append(cardValues, val)
		}
	}

	return Hand{cardValues: cardValues, bid: bid}
}

func compareHands(handOne Hand, handTwo Hand) int {
	if handOne.isFiveOfKind() || handTwo.isFiveOfKind() {
		if !handOne.isFiveOfKind() {
			return -1
		}
		if !handTwo.isFiveOfKind() {
			return 1
		}

		return comparisonTieBreaker(handOne, handTwo)
	}
	if handOne.isFourOfKind() || handTwo.isFourOfKind() {
		if !handOne.isFourOfKind() {
			return -1
		}
		if !handTwo.isFourOfKind() {
			return 1
		}

		return comparisonTieBreaker(handOne, handTwo)
	}
	if handOne.isFullHouse() || handTwo.isFullHouse() {
		if !handOne.isFullHouse() {
			return -1
		}
		if !handTwo.isFullHouse() {
			return 1
		}

		return comparisonTieBreaker(handOne, handTwo)
	}
	if handOne.isThreeOfKind() || handTwo.isThreeOfKind() {
		if !handOne.isThreeOfKind() {
			return -1
		}
		if !handTwo.isThreeOfKind() {
			return 1
		}

		return comparisonTieBreaker(handOne, handTwo)
	}
	if handOne.isTwoPair() || handTwo.isTwoPair() {
		if !handOne.isTwoPair() {
			return -1
		}
		if !handTwo.isTwoPair() {
			return 1
		}

		return comparisonTieBreaker(handOne, handTwo)
	}
	if handOne.isOnePair() || handTwo.isOnePair() {
		if !handOne.isOnePair() {
			return -1
		}
		if !handTwo.isOnePair() {
			return 1
		}

		return comparisonTieBreaker(handOne, handTwo)
	}

	return comparisonTieBreaker(handOne, handTwo)
}

func comparisonTieBreaker(handOne Hand, handTwo Hand) int {
	for i := range handOne.cardValues {
		diff := handOne.cardValues[i] - handTwo.cardValues[i]
		if diff != 0 {
			return diff
		}
	}

	return 0
}
