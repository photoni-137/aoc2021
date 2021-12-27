package main

import (
	"aoc2021/shared"
	"fmt"
	"strings"
)

type Dot struct{ x, y int }

type Paper map[Dot]int

type Fold struct {
	direction rune
	line      int
}

func (d *Dot) Fold(fold Fold) {
	switch fold.direction {
	case 'x':
		if d.x > fold.line {
			d.x = 2*fold.line - d.x
		}
		return
	case 'y':
		if d.y > fold.line {
			d.y = 2*fold.line - d.y
		}
		return
	}
}

func (d Dot) ToString() string {
	return fmt.Sprintf("%d,%d", d.x, d.y)
}

func (p Paper) Draw(d Dot) {
	p[d]++
}

func (p Paper) Count() (count int) {
	for _, dots := range p {
		if dots > 0 {
			count++
		}
	}
	return
}

func parseLines(lines []string) (dots []Dot, folds []Fold) {
	for _, line := range lines {
		if strings.Contains(line, ",") {
			dots = append(dots, parseDot(line))
		} else if strings.Contains(line, "fold along") {
			folds = append(folds, parseFold(line))
		}
	}
	return
}

func parseDot(line string) (dot Dot) {
	_, err := fmt.Sscanf(line, "%d,%d", &dot.x, &dot.y)
	shared.Handle(err)
	return
}

func parseFold(line string) (fold Fold) {
	_, err := fmt.Sscanf(line, "fold along %c=%d", &fold.direction, &fold.line)
	shared.Handle(err)
	return
}

func prettyPrint(dots []Dot) {
	var solution [6][39]string
	for i, row := range solution {
		for j := range row {
			solution[i][j] = " "
		}
	}
	for _, dot := range dots {
		solution[dot.y][dot.x] = "x"
	}

	fmt.Println()
	for _, row := range solution {
		fmt.Println(strings.Join(row[:], " "))
	}
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	dots, folds := parseLines(lines)

	paper := make(Paper)
	for _, fold := range folds[:1] {
		for _, dot := range dots {
			dot.Fold(fold)
			paper.Draw(dot)
		}
	}
	fmt.Printf("After the first fold, %d distinct dots remain.\n", paper.Count())

	for _, fold := range folds {
		var newDots []Dot
		for _, dot := range dots {
			dot.Fold(fold)
			newDots = append(newDots, dot)
		}
		dots = newDots
	}

	prettyPrint(dots)
}
