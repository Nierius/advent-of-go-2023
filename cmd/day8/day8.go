package main

import (
	"fmt"
	"log"
	"ourmodule/pkg/utils"
	"strings"
)

type Node struct {
	self  string
	left  string
	right string
}

func main() {
	contents := utils.ReadEntireFileToString("./input.txt")
	contentsSplitted := strings.Split(contents, "\n\n")

	nodes := parseNodes(contentsSplitted[1])

	part1(contentsSplitted[0], nodes)
	part2(contentsSplitted[0], nodes)
}

func part1(instructions string, nodes []Node) {
	currentNodeValue := "AAA"
	steps := 0
	for currentNodeValue != "ZZZ" {
		currentNode := utils.Filter(nodes, func(n Node) bool {
			return n.self == currentNodeValue
		})[0]
		instruction := instructions[steps%len(instructions)]
		switch instruction {
		case 'L':
			currentNodeValue = currentNode.left
		case 'R':
			currentNodeValue = currentNode.right
		default:
			log.Fatal("out of range")
		}
		steps++
	}

	fmt.Println("Part 1: ", steps)
}

func part2(instructions string, nodes []Node) {
	startingNodes := utils.Filter(nodes, func(n Node) bool {
		return n.self[len(n.self)-1] == 'A'
	})

	allSteps := []int{}
	for _, currentNode := range startingNodes {
		steps := 0
		for currentNode.self[2] != 'Z' {
			instruction := instructions[steps%len(instructions)]
			switch instruction {
			case 'L':
				currentNode = utils.Filter(nodes, func(n Node) bool {
					return n.self == currentNode.left
				})[0]
			case 'R':
				currentNode = utils.Filter(nodes, func(n Node) bool {
					return n.self == currentNode.right
				})[0]
			default:
				log.Fatal("out of range")
			}
			steps++
		}

		allSteps = append(allSteps, steps)
	}

	fmt.Println("Part 2: ", utils.Lcm(allSteps))
}

func parseNodes(input string) []Node {
	nodes := []Node{}
	for _, rawNode := range strings.Split(input, "\n") {
		if !utils.StringIsNotEmpty(rawNode) {
			continue
		}

		rawNodeSplit := strings.Split(rawNode, " = ")
		leftAndRightSplit := strings.Split(rawNodeSplit[1], ", ")
		left := strings.Trim(leftAndRightSplit[0], "()")
		right := strings.Trim(leftAndRightSplit[1], "()")

		nodes = append(nodes, Node{self: rawNodeSplit[0], left: left, right: right})
	}

	return nodes
}
