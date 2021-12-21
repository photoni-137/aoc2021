package main

import (
	"aoc2021/shared"
	"fmt"
	"strings"
)

type Point struct{ x, y int }
type Paper map[string]int
type Fold struct {
	direction rune
	line      int
}

func (p *Point) Fold(fold Fold) {
	switch fold.direction {
	case 'x':
		if p.x > fold.line {
			p.x = 2*fold.line - p.x
		}
		return
	case 'y':
		if p.y > fold.line {
			p.y = 2*fold.line - p.y
		}
		return
	}
}

func (p Point) ToString() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p Paper) Draw(point Point) {
	p[point.ToString()]++
}

func (p Paper) Count() (count int) {
	for _, points := range p {
		if points > 0 {
			count++
		}
	}
	return
}

func parseLines(lines []string) (points []Point, folds []Fold, err error) {
	for _, line := range lines {
		switch {
		case strings.Contains(line, ","):
			var point Point
			_, err = fmt.Sscanf(line, "%d,%d", &point.x, &point.y)
			if err != nil {
				return
			}
			points = append(points, point)
		case strings.Contains(line, "fold along"):
			var fold Fold
			_, err = fmt.Sscanf(line, "fold along %c=%d", &fold.direction, &fold.line)
			if err != nil {
				return
			}
			folds = append(folds, fold)
		}
	}
	return
}

func main() {
	lines := shared.ParseInputFile("paper.txt")
	points, folds, err := parseLines(lines)
	shared.Handle(err)

	paper := make(Paper)
	for _, fold := range folds[:1] {
		for _, point := range points {
			point.Fold(fold)
			paper.Draw(point)
		}
	}
	fmt.Println(paper.Count())

	for _, fold := range folds {
		var newPoints []Point
		for _, point := range points {
			point.Fold(fold)
			newPoints = append(newPoints, point)
		}
		points = newPoints
	}
	var solution [40][8]string
	for i, row := range solution {
		for j := range row {
			solution[i][j] = " "
		}
	}
	for _, point := range points {
		solution[point.x][point.y] = "x"
	}
	for _, row := range solution {
		fmt.Println(strings.Join(row[:], " "))
	}
}
