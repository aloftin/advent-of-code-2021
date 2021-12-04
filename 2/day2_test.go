package main

import "testing"

func TestPartOne(t *testing.T) {
	lines := getLinesFromFile()
	runPartOne(lines)
}

func TestPartTwo(t *testing.T) {
	lines := getLinesFromFile()
	runPartTwo(lines)
}
