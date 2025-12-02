package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type IdRange struct {
	Low       int
	LowFront  int
	High      int
	HighFront int
}

func PartOne(ranges []IdRange) int {
	invalidSum := 0
	for _, idRange := range ranges {
		if idRange.LowFront > idRange.HighFront {
			idRange.HighFront = idRange.LowFront
		}
		for i := idRange.LowFront; i <= idRange.HighFront; i++ {
			potentialIdStr := strconv.Itoa(i) + strconv.Itoa(i)
			potentialId, err := strconv.Atoi(potentialIdStr)
			if err != nil {
				fmt.Printf("Couldn't parse doubled value %v: %v", potentialIdStr, err)
				return -1
			}
			if potentialId >= idRange.Low && potentialId <= idRange.High {
				invalidSum += potentialId
			}
		}
	}
	return invalidSum
}

func PartTwo(ranges []IdRange) int {
	invalidIds := map[int]struct{}{}
	for _, idRange := range ranges {
		fmt.Printf("Current range %+v\n", idRange)
		if idRange.HighFront < idRange.LowFront {
			idRange.HighFront = idRange.LowFront + ((idRange.LowFront%10 - 10) * -1)
		}
		for i := idRange.LowFront; i <= idRange.HighFront; i++ {
			if i == 0 {
				continue
			}
			lowStr := strconv.Itoa(i)
			for j := 1; j <= len(lowStr); j++ {
				piece, err := strconv.Atoi(lowStr[0:j])
				if err != nil {
					fmt.Printf("Couldn't convert %v to int: %v", lowStr[0:j], err)
					return -1
				}
				potentialId := piece
				for potentialId < idRange.Low {
					potentialId = potentialId*int(math.Pow(10, float64(j))) + piece
				}
				if potentialId <= idRange.High {
					fmt.Printf("Invalid ID: %v from %v\n", potentialId, piece)
					invalidIds[potentialId] = struct{}{}
				}
			}
		}
	}
	invalidSum := 0
	for i, _ := range invalidIds {
		invalidSum += i
	}
	return invalidSum
}

// 22471547034
// 22471583899

func main() {
	//filename := "./test.txt"
	filename := "../../AdventOfCodeInputs/2025/Day02/input.txt"
	lines, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Couldn't open file %v: %v", filename, err)
		return
	}
	rangesStr := strings.Split(strings.Trim(string(lines), "\n"), ",")
	var ranges []IdRange
	for _, s := range rangesStr {
		edges := strings.Split(s, "-")
		if len(edges) != 2 {
			fmt.Printf("- split of %v resulted in %v tokens!", s, len(edges))
			return
		}
		idRange := IdRange{}
		idRange.Low, err = strconv.Atoi(edges[0])
		if err != nil {
			fmt.Printf("Couldn't convert %v to int: %v", edges[0], err)
			return
		}
		if idRange.Low > 9 {
			lowFrontStr := edges[0][:(len(edges[0]))/2]
			idRange.LowFront, err = strconv.Atoi(lowFrontStr)
			if err != nil {
				fmt.Printf("Couldn't get int from %v: %v", lowFrontStr, err)
				return
			}
		}
		idRange.High, err = strconv.Atoi(edges[1])
		if err != nil {
			fmt.Printf("Couldn't convert %v to int: %v", edges[0], err)
			return
		}
		if idRange.High > 9 {
			highFrontStr := edges[1][:(len(edges[1])+1)/2]
			idRange.HighFront, err = strconv.Atoi(highFrontStr)
			if err != nil {
				fmt.Printf("Couldn't get int from %v: %v", highFrontStr, err)
				return
			}
		}
		ranges = append(ranges, idRange)
	}
	fmt.Printf("Part 1: %v\n", PartOne(ranges))
	fmt.Printf("Part 2: %v\n", PartTwo(ranges))
}

// 13108240966
// 13108371860
