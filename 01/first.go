package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./example_input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

    ints := []int {}
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
	re := regexp.MustCompile("[0-9]")
    nums := re.FindAllString(input, -1)

    numstr := fmt.Sprintf("%s%s", nums[0], nums[len(nums)-1])

	val, err := strconv.Atoi(numstr)
	if err != nil {
		log.Fatal(err)
	}

	return val
}
