package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
)

func parseDepths(lines []string) (numbers []int) {
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		shared.Handle(err)
		numbers = append(numbers, number)
	}
	return
}

func slidingSums(numbers []int, windowSize int) []int {
	var sums []int
	for i := 0; i <= len(numbers)-windowSize; i++ {
		sum := 0
		for _, number := range numbers[i : i+windowSize] {
			sum += number
		}
		sums = append(sums, sum)
	}
	return sums
}

func countIncreases(numbers []int) int {
	counter := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			counter++
		}
	}
	return counter
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	depths := parseDepths(lines)
	fmt.Printf("Depth increased %d times in %d measurements.\n",
		countIncreases(depths), len(depths))

	depthWindows := slidingSums(depths, 3)
	fmt.Printf("Depth windows increased %d times in %d aggregated measurements.\n",
		countIncreases(depthWindows), len(depthWindows))
}
