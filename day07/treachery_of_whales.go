package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile := "crabs.txt"
	positions, err := parseInputFile(inputFile)
	handle(err)

	sortedPositions := sort.IntSlice(positions)
	sortedPositions.Sort()
	lower, upper := boundaries(sortedPositions)
	fuelCosts := make(map[int]int)
	for point := lower; point <= upper; point++ {
		fuelCosts[point] = totalFuel(sortedPositions, point)
	}
	minimum := fuelCosts[0]
	fmt.Printf("Initial fuel cost: %d at point %d.\n", fuel, 0)
	for point, fuel := range fuelCosts {
		if fuel < minimum {
			minimum = fuel
			fmt.Printf("New minimal fuel cost: %d at point %d.\n", fuel, point)
		}
	}
	fmt.Println("That's the best one!")
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInputFile(filePath string) (positions []int, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		positions, err = parsePositions(scanner.Text())
	}
	return
}

func parsePositions(input string) (positions []int, err error) {
	for _, s := range strings.Split(input, ",") {
		var position int
		position, err = strconv.Atoi(s)
		if err != nil {
			return
		}
		positions = append(positions, position)
	}
	return
}

func fuel(x, y int) (f int) {
	distance := x - y
	if distance < 0 {
		distance = -distance
	}
	f = distance * (distance + 1) / 2
	return
}

func totalFuel(positions []int, point int) (sum int) {
	for _, position := range positions {
		sum += fuel(position, point)
	}
	return
}

func boundaries(sortedPositions sort.IntSlice) (lo, hi int) {
	lo = sortedPositions[0]
	hi = sortedPositions[len(sortedPositions)-1]
	return
}
