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

	ints := []int{}
	for scanner.Scan() {
		ints = append(ints, exportNums(scanner.Text()))
		fmt.Println(exportNums(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, v := range ints {
		sum += v
	}

	fmt.Printf("Sum is %d", sum)
}

func exportNums(input string) int {
	input = strings.ReplaceAll(input, "one", "o1e")
	input = strings.ReplaceAll(input, "two", "t2o")
	input = strings.ReplaceAll(input, "three", "t3e")
	input = strings.ReplaceAll(input, "four", "4")
	input = strings.ReplaceAll(input, "five", "5e")
	input = strings.ReplaceAll(input, "six", "6")
	input = strings.ReplaceAll(input, "seven", "7n")
	input = strings.ReplaceAll(input, "eight", "e8t")
	input = strings.ReplaceAll(input, "nine", "n9e")

	re := regexp.MustCompile("[0-9]")
	nums := re.FindAllString(input, -1)

	numstr := fmt.Sprintf("%s%s", nums[0], nums[len(nums)-1])

	val, err := strconv.Atoi(numstr)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

