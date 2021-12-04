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
	lines := getLinesFromFile()
	fmt.Printf("Part one answer: %v\n", runPartOne(lines))
	fmt.Printf("Part two answer: %v\n", runPartTwo(lines))
}

func runPartTwo(lines []string) int {
	var forward, up, down, aim int
	for _, line := range lines {
		// split line by space char
		// first half is the string key ("forward", "up", "down")
		// second half is the int value
		splitLine := strings.Split(line, " ")

		v := toInt(splitLine[1])
		if splitLine[0] == "forward" {
			forward += v
			down += aim * v
		}
		if splitLine[0] == "up" {
			aim -= v
		}
		if splitLine[0] == "down" {
			aim += v
		}
	}

	answer := forward * (down - up)
	return answer
}

func runPartOne(lines []string) int {
	positions := make(map[string]int)
	for _, line := range lines {
		// split line by space char
		// first half is the string key ("forward", "up", "down")
		// second half is the int value
		splitLine := strings.Split(line, " ")
		positions[splitLine[0]] += toInt(splitLine[1])
	}

	answer := positions["forward"] * (positions["down"] - positions["up"])
	return answer
}

func getLinesFromFile() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
