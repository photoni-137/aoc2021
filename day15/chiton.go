package main

import (
	"aoc2021/shared"
	"container/heap"
	"fmt"
	"math"
	"strconv"
)

const MaxInt = math.MaxInt64

type Coordinates struct{ x, y int }

type Node struct {
	coordinates Coordinates
	risk        int
	totalRisk   int
	index       int
}

type Cave struct {
	lenX, lenY int
	nodes      map[Coordinates]*Node
}

func (c Cave) Neighbors(n *Node) (neighbors []*Node) {
	x, y := n.coordinates.x, n.coordinates.y
	if x > 0 {
		neighbors = append(neighbors, c.nodes[Coordinates{x - 1, y}])
	}
	if x < c.lenX-1 {
		neighbors = append(neighbors, c.nodes[Coordinates{x + 1, y}])
	}
	if y > 0 {
		neighbors = append(neighbors, c.nodes[Coordinates{x, y - 1}])
	}
	if y < c.lenY-1 {
		neighbors = append(neighbors, c.nodes[Coordinates{x, y + 1}])
	}
	return
}

func parseLines(lines []string) (cave Cave) {
	cave.lenX, cave.lenY = len(lines), len(lines[0])
	cave.nodes = make(map[Coordinates]*Node)
	index := 0
	for x, line := range lines {
		for y, r := range line {
			coordinates := Coordinates{x, y}
			risk, err := strconv.Atoi(string(r))
			shared.Handle(err)
			cave.nodes[coordinates] = &Node{
				coordinates: coordinates,
				risk:        risk,
				totalRisk:   MaxInt,
				index:       index,
			}
			index++
		}
	}
	return
}

func parseAndExtendLines(lines []string, repeat int) (cave Cave) {
	cave.lenX, cave.lenY = 5*len(lines), 5*len(lines[0])
	cave.nodes = make(map[Coordinates]*Node)
	index := 0
	for x, line := range lines {
		for y, r := range line {
			for i := 0; i < repeat; i++ {
				for j := 0; j < repeat; j++ {
					coordinates := Coordinates{x + i*len(lines), y + j*len(lines[0])}
					risk, err := strconv.Atoi(string(r))
					shared.Handle(err)
					risk = modify(risk, i+j)
					cave.nodes[coordinates] = &Node{
						coordinates: coordinates,
						risk:        risk,
						totalRisk:   MaxInt,
						index:       index,
					}
					index++
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

func (c Cave) lowestRisk() int {
	visited := make(map[Coordinates]bool)

	start := c.nodes[Coordinates{0, 0}]
	end := c.nodes[Coordinates{c.lenX - 1, c.lenY - 1}]

	start.totalRisk = 0
	queue := PriorityQueue{start}
	heap.Init(&queue)

	for queue.Len() > 0 {
		node := heap.Pop(&queue).(*Node)
		current := node.coordinates

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, neighbor := range c.Neighbors(node) {
			if !visited[neighbor.coordinates] {
				newRisk := node.totalRisk + neighbor.risk
				if newRisk < neighbor.totalRisk {
					neighbor.totalRisk = newRisk
					heap.Push(&queue, neighbor)
				}
			}
		}
	}
	return end.totalRisk
}

func main() {
	lines := shared.ParseInputFile("input.txt")

	cave := parseLines(lines)
	fmt.Println(cave.lowestRisk())

	extendedCave := parseAndExtendLines(lines, 5)
	fmt.Println(extendedCave.lowestRisk())
}
