package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	depths := getInputFromFile()

	fmt.Printf("Part one answer: %v\n", runPartOne(depths))
	fmt.Printf("Part one answer: %v\n", runPartTwo(depths))
}

func runPartOne(depths []int) int {
	return countDepthIncreases(depths)
}

func runPartTwo(depths []int) int {
	return countSlidingWindowDepthIncreases(depths)
}

func countDepthIncreases(depths []int) int {
	count := 0
	for i := 0; i < len(depths)-1; i++ {
		if depths[i] < depths[i+1] {
			count++
		}
	}
	return count
}

func countSlidingWindowDepthIncreases(depths []int) int {
	var previousSum, currentSum, increases int

	for i := range depths {
		if i == len(depths)-3 {
			// not enough values left to create a new window
			break
		} else {
			previousSum = currentSum
			currentSum = depths[i] + depths[i+1] + depths[i+2]
			if currentSum > previousSum {
				increases++
			}
		}

	}
	return increases
}

func getInputFromFile() []int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	depths := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		depths = append(depths, toInt(line))
	}

	return depths
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
