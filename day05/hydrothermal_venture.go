package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

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

const gridSize = 1000

type SquareGrid [gridSize][gridSize]int

func (g *SquareGrid) Draw(line Line) {
	switch line.Direction() {
	case "x":
		g.drawX(line)
	case "y":
		g.drawY(line)
	case "diagonal":
		g.drawDiagonal(line)
	}
	return
}

func (g *SquareGrid) drawX(line Line) {
	start := line.start.x
	end := line.end.x
	y := line.start.y
	if start > end {
		start, end = end, start
	}
	for x := start; x <= end; x++ {
		g[x][y]++
	}
}

func (g *SquareGrid) drawY(line Line) {
	start := line.start.y
	end := line.end.y
	x := line.start.x
	if start > end {
		start, end = end, start
	}
	for y := start; y <= end; y++ {
		g[x][y]++
	}
}

func (g *SquareGrid) drawDiagonal(line Line) {
	startX := line.start.x
	endX := line.end.x
	startY := line.start.y
	endY := line.end.y
	if startX > endX {
		startX, endX = endX, startX
		startY, endY = endY, startY
	}
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

func (g SquareGrid) CountFieldsOver(threshold int) (count int) {
	for _, row := range g {
		for _, value := range row {
			if value >= threshold {
				count++
			}
		}
	}
	return
}

func (g SquareGrid) PrettyPrint() {
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

func parseInputFile(filePath string) (lines []Line, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line, err := parseLine(scanner.Text())
		if err == nil {
			lines = append(lines, line)
		}
	}
	return
}

func parseLine(input string) (line Line, err error) {
	_, err = fmt.Sscanf(input, "%d,%d -> %d,%d", &line.start.x, &line.start.y, &line.end.x, &line.end.y)
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
	inputFile := "vents.txt"
	ventLines, err := parseInputFile(inputFile)
	handle(err)

	shavedLines := shaveDiagonal(ventLines)
	oceanFloor := SquareGrid{}
	for _, line := range shavedLines {
		oceanFloor.Draw(line)
	}
	fmt.Printf("Drew %d non-diagonal vent lines over the ocean floor. %d points show dangerous activity.\n",
		len(shavedLines), oceanFloor.CountFieldsOver(2))

	fmt.Println("This is dangerously off, though. We need to start anew, this time including diagonal vents!")
	oceanFloor = SquareGrid{}
	for _, line := range ventLines {
		oceanFloor.Draw(line)
	}
	fmt.Printf("Drew %d vent lines over the ocean floor. Now, %d points show dangerous activity.\n",
		len(ventLines), oceanFloor.CountFieldsOver(2))
}
