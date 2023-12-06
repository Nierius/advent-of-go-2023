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

	intsPartOne := []int{}
	intsPartTwo := []int{}
	for scanner.Scan() {
		intsPartOne = append(intsPartOne, exportNums(scanner.Text()))
		intsPartTwo = append(intsPartTwo, exportNums(replaceNumStrings(scanner.Text())))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sumPartOne := 0
	sumPartTwo := 0
	for i := range intsPartOne {
		sumPartOne += intsPartOne[i]
		sumPartTwo += intsPartTwo[i]
	}

	fmt.Println("Sum part1 is ", sumPartOne)
	fmt.Println("Sum part2 is ", sumPartTwo)
}

func exportNums(input string) int {
	re := regexp.MustCompile("[0-9]")
	nums := re.FindAllString(input, -1)

	numStr := fmt.Sprintf("%s%s", nums[0], nums[len(nums)-1])

	val, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func replaceNumStrings(input string) string {
	input = strings.ReplaceAll(input, "one", "o1e")
	input = strings.ReplaceAll(input, "two", "t2o")
	input = strings.ReplaceAll(input, "three", "t3e")
	input = strings.ReplaceAll(input, "four", "4")
	input = strings.ReplaceAll(input, "five", "5e")
	input = strings.ReplaceAll(input, "six", "6")
	input = strings.ReplaceAll(input, "seven", "7n")
	input = strings.ReplaceAll(input, "eight", "e8t")
	input = strings.ReplaceAll(input, "nine", "n9e")

	return input
}
