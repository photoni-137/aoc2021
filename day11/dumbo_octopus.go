package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
)

type Dumbo struct {
	energy  int
	flashed bool
}

type Point struct{ x, y int }

type Grid [][]Dumbo

func (d Dumbo) CanFlash() bool {
	return !d.flashed && d.energy > 9
}

func (g Grid) At(p Point) *Dumbo {
	return &g[p.x][p.y]
}

func (g Grid) Neighbors(p Point) (neighbors []Point) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			neighbor := Point{p.x + dx, p.y + dy}
			if g.ValidPoint(neighbor) && neighbor != p {
				neighbors = append(neighbors, Point{p.x + dx, p.y + dy})
			}
		}
	}
	return
}

func (g Grid) ValidPoint(p Point) bool {
	xMax, yMax := len(g), len(g[0])
	validX := p.x >= 0 && p.x < xMax
	validY := p.y >= 0 && p.y < yMax
	return validX && validY
}

func (g Grid) Reset() {
	for x, row := range g {
		for y, dumbo := range row {
			if dumbo.flashed {
				g[x][y] = Dumbo{}
			}
		}
	}
}

func (g Grid) increaseEnergy() {
	for x, row := range g {
		for y := range row {
			g[x][y].energy++
		}
	}
}

func (g Grid) Flash(p Point) {
	for _, neighbor := range g.Neighbors(p) {
		g.At(neighbor).energy++
	}
	g.At(p).flashed = true
}

func (g Grid) Evolve() (flashes int) {
	g.increaseEnergy()
	loop := true
	for loop {
		newFlashes := g.FlashAll()
		flashes += newFlashes
		loop = newFlashes > 0
	}
	g.Reset()
	return
}

func (g Grid) FlashAll() (flashes int) {
	for x, row := range g {
		for y, dumbo := range row {
			if dumbo.CanFlash() {
				g.Flash(Point{x, y})
				flashes++
			}
		}
	}
	return
}

func (g Grid) IsSynchronized() bool {
	synchronizedLevel := g[0][0].energy
	for _, row := range g {
		for _, dumbo := range row {
			if dumbo.energy != synchronizedLevel {
				return false
			}
		}
	}
	return true
}

func parseGrid(lines []string) (grid Grid) {
	grid = make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = make([]Dumbo, len(line))
		for j, r := range line {
			number, err := strconv.Atoi(string(r))
			shared.Handle(err)
			grid[i][j] = Dumbo{number, false}
		}
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	grid := parseGrid(lines)

	flashes, step := 0, 0
	for !grid.IsSynchronized() {
		if step == 100 {
			fmt.Printf("%d flashes after 100 steps.\n", flashes)
		}
		flashes += grid.Evolve()
		step++
	}
	fmt.Printf("Synchronized after %d steps!\n", step)
}
