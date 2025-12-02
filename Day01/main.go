package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	startValue = 50
	dialLength = 100
)

func PartOne(lines []string) int {
	dialValue := startValue
	password := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		distance, err := strconv.Atoi(l[1:])
		if err != nil {
			fmt.Printf("Couldn't format %v as int: %v", l[1:], err)
			return -1
		}
		if l[0] == 'L' {
			dialValue -= distance
		} else {
			dialValue += distance
		}
		dialValue = (dialValue + dialLength) % dialLength
		if dialValue == 0 {
			password++
		}
	}
	return password
}

func PartTwo(lines []string) int {
	dialValue := startValue
	password := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		distance, err := strconv.Atoi(l[1:])
		if err != nil {
			fmt.Printf("Couldn't format %v as int: %v", l[1:], err)
			return -1
		}
		if distance >= dialLength {
			password += distance / dialLength
			distance = distance % dialLength
		}
		if distance == 0 {
			continue
		}
		wasZero := dialValue == 0
		switch l[0] {
		case 'L':
		case 'R':
		default:
		}
		if l[0] == 'L' {
			dialValue -= distance
		} else {
			dialValue += distance
		}
		if !wasZero && (dialValue > dialLength || dialValue < 0) {
			password++
		}
		dialValue = (dialValue + dialLength) % dialLength
		if dialValue == 0 {
			password++
		}
	}
	return password
}

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2025/Day01/input.txt"

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Couldn't open file %v: %v", filename, err)
		return
	}

	lines := strings.Split(string(contents), "\n")
	fmt.Printf("Part 1: %v\n", PartOne(lines))
	fmt.Printf("Part 2: %v\n", PartTwo(lines))
}
