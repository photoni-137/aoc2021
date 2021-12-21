package main

import (
	"aoc2021/shared"
	"fmt"
	"math"
	"sort"
	"strconv"
)

const MaxInt = math.MaxInt64

type Coordinates struct{ x, y int }

type Queue []Coordinates

func (q Queue) Push(c Coordinates) Queue {
	return append(q, c)
}

func (q Queue) Pop() (Coordinates, Queue) {
	return q[0], q[1:]
}

func (q Queue) Sort(risk map[Coordinates]int) {
	sort.Slice(q, func(i, j int) bool {
		return risk[q[i]] < risk[q[j]]
	})
}

type Cave struct {
	lenX, lenY int
	risk       map[Coordinates]int
}

func (c Cave) Neighbors(coordinates Coordinates) (neighbors []Coordinates) {
	x, y := coordinates.x, coordinates.y
	if x > 0 {
		neighbors = append(neighbors, Coordinates{x - 1, y})
	}
	if x < c.lenX-1 {
		neighbors = append(neighbors, Coordinates{x + 1, y})
	}
	if y > 0 {
		neighbors = append(neighbors, Coordinates{x, y - 1})
	}
	if y < c.lenY-1 {
		neighbors = append(neighbors, Coordinates{x, y + 1})
	}
	return
}

func parseLines(lines []string) (cave Cave) {
	cave.lenX, cave.lenY = len(lines), len(lines[0])
	cave.risk = make(map[Coordinates]int)
	for x, line := range lines {
		for y, r := range line {
			coordinates := Coordinates{x, y}
			risk, err := strconv.Atoi(string(r))
			shared.Handle(err)
			cave.risk[coordinates] = risk
		}
	}
	return
}

func parseAndExtendLines(lines []string) (cave Cave) {
	cave.lenX, cave.lenY = 5*len(lines), 5*len(lines[0])
	cave.risk = make(map[Coordinates]int)
	for x, line := range lines {
		for y, r := range line {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					coordinates := Coordinates{x + i*len(lines), y + j*len(lines[0])}
					risk, err := strconv.Atoi(string(r))
					shared.Handle(err)
					risk = modify(risk, i+j)
					cave.risk[coordinates] = risk
				}
			}
		}
	}
	return
}

func modify(risk, increase int) (newRisk int) {
	newRisk = risk + increase
	if newRisk > 9 {
		newRisk %= 9
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")

	cave := parseLines(lines)
	fmt.Println(cave.lowestRisk())

	extendedCave := parseAndExtendLines(lines)
	fmt.Println(extendedCave.lowestRisk())
}

func (c Cave) lowestRisk() int {
	neighbors := make(map[Coordinates][]Coordinates)
	visited := make(map[Coordinates]bool)
	risk := make(map[Coordinates]int)
	previous := make(map[Coordinates]Coordinates)

	start, end := Coordinates{0, 0}, Coordinates{c.lenX - 1, c.lenY - 1}
	for coordinates := range c.risk {
		risk[coordinates] = MaxInt
		neighbors[coordinates] = c.Neighbors(coordinates)
	}
	risk[start] = 0
	queue := Queue{start}

	for len(queue) > 0 {
		var current Coordinates
		queue.Sort(risk)
		current, queue = queue.Pop()
		if visited[current] {
			continue
		}
		visited[current] = true

		for _, neighbor := range neighbors[current] {
			if !visited[neighbor] {
				newRisk := risk[current] + c.risk[neighbor]
				if newRisk < risk[neighbor] {
					risk[neighbor] = newRisk
					previous[neighbor] = current
					queue = queue.Push(neighbor)
				}
			}
		}
	}
	return risk[end]
}
