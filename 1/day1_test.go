package main

import "testing"

func TestPartOne(t *testing.T) {
	depths := getInputFromFile()
	runPartOne(depths)
}

func TestPartTwo(t *testing.T) {
	depths := getInputFromFile()
	runPartTwo(depths)
}
