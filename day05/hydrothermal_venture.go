package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

const gridSize = 1000

type Grid [gridSize][gridSize]int

type Point struct{ x, y int }

type Line struct{ start, end Point }

func (l Line) Direction() (direction string) {
	switch {
	case l.start.x == l.end.x && l.start.y == l.end.y:
		direction = ""
	case l.start.y == l.end.y:
		direction = "x"
	case l.start.x == l.end.x:
		direction = "y"
	default:
		direction = "diagonal"
	}
	return
}

func (g *Grid) Draw(l Line) {
	switch l.Direction() {
	case "x":
		g.drawX(l)
	case "y":
		g.drawY(l)
	case "diagonal":
		g.drawDiagonal(l)
	}
	return
}

func (g *Grid) drawX(l Line) {
	y := l.start.y
	for x := l.start.x; x <= l.end.x; x++ {
		g[x][y]++
	}
}

func (g *Grid) drawY(l Line) {
	x := l.start.x
	start, end := l.start.y, l.end.y
	if start > end {
		start, end = end, start
	}
	for y := start; y <= end; y++ {
		g[x][y]++
	}
}

func (g *Grid) drawDiagonal(l Line) {
	startX, endX := l.start.x, l.end.x
	startY, endY := l.start.y, l.end.y
	if startY < endY {
		for x, y := startX, startY; x <= endX && y <= endY; x, y = x+1, y+1 {
			g[x][y]++
		}
	} else {
		for x, y := startX, startY; x <= endX && y >= endY; x, y = x+1, y-1 {
			g[x][y]++
		}
	}
}

func (g Grid) CountFieldsOver(threshold int) (count int) {
	for _, row := range g {
		for _, value := range row {
			if value >= threshold {
				count++
			}
		}
	}
	return
}

func parseVents(lines []string) (vents []Line) {
	for _, line := range lines {
		var vent Line
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &vent.start.x, &vent.start.y, &vent.end.x, &vent.end.y)
		shared.Handle(err)
		if vent.start.x > vent.end.x {
			vent.start, vent.end = vent.end, vent.start
		}
		vents = append(vents, vent)
	}
	return
}

func shaveDiagonal(lines []Line) (shaved []Line) {
	for _, line := range lines {
		if d := line.Direction(); d == "x" || d == "y" {
			shaved = append(shaved, line)
		}
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	vents := parseVents(lines)

	withoutDiagonals := shaveDiagonal(vents)
	oceanFloor := Grid{}
	for _, line := range withoutDiagonals {
		oceanFloor.Draw(line)
	}
	fmt.Printf("Drew %d non-diagonal vent lines over the ocean floor. %d points show dangerous activity.\n",
		len(withoutDiagonals), oceanFloor.CountFieldsOver(2))

	fmt.Println("This is dangerously off, though. We need to start anew, this time including diagonal vents!")
	oceanFloor = Grid{}
	for _, line := range vents {
		oceanFloor.Draw(line)
	}
	fmt.Printf("Drew %d vent lines over the ocean floor. Now, %d points show dangerous activity.\n",
		len(vents), oceanFloor.CountFieldsOver(2))
}

func (g Grid) PrettyPrint() {
	for _, row := range g {
		var sb strings.Builder
		for _, value := range row {
			switch value {
			case 0:
				sb.WriteString(" ")
			default:
				sb.WriteString(strconv.Itoa(value))
			}
		}
		fmt.Println(sb.String())
	}
}
