package main

import (
	"aoc2021/shared"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func parsePositions(line string) (positions []int) {
	for _, s := range strings.Split(line, ",") {
		position, err := strconv.Atoi(s)
		shared.Handle(err)
		positions = append(positions, position)
	}
	return
}

func distance(x, y int) (dist int) {
	dist = x - y
	if dist < 0 {
		dist = -dist
	}
	return
}

func totalFuel(positions []int, target int, fuelCost func(int) int) (sum int) {
	for _, position := range positions {
		dist := distance(position, target)
		sum += fuelCost(dist)
	}
	return
}

func boundaries(positions []int) (lo, hi int) {
	sortedPositions := sort.IntSlice(positions)
	sortedPositions.Sort()
	lo = sortedPositions[0]
	hi = sortedPositions[len(sortedPositions)-1]
	return
}

func optimalPosition(positions []int, fuelCost func(int) int) (target, minimalFuel int) {
	lower, upper := boundaries(positions)
	fuelCosts := make(map[int]int)
	for point := lower; point <= upper; point++ {
		fuelCosts[point] = totalFuel(positions, point, fuelCost)
	}

	target, minimalFuel = -1, math.MaxInt
	for point, fuel := range fuelCosts {
		if fuel < minimalFuel {
			target, minimalFuel = point, fuel
		}
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	positions := parsePositions(lines[0])

	simpleFuel := func(d int) int { return d }
	target, fuel := optimalPosition(positions, simpleFuel)
	fmt.Printf("Assuming the simplest model, the optimal point is at x = %d, costing %d units of fuel.\n",
		target, fuel)

	gaussFuel := func(d int) int { return d * (d + 1) / 2 }
	target, fuel = optimalPosition(positions, gaussFuel)
	fmt.Printf("Given the more complex model, the optimal point is at x = %d, costing %d units of fuel.\n",
		target, fuel)
}
