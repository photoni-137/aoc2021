package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dumbo struct {
	energy  int
	flashed bool
}

func (d Dumbo) CanFlash() bool {
	return !d.flashed && d.energy > 9
}

type Point struct{ x, y int }

func (p Point) IsValid() bool {
	validX := p.x >= 0 && p.x < 10
	validY := p.y >= 0 && p.y < 10
	return validX && validY
}

func (p Point) Neighbors() (neighbors []Point) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			neighbor := Point{p.x + dx, p.y + dy}
			if neighbor.IsValid() && neighbor != p {
				neighbors = append(neighbors, Point{p.x + dx, p.y + dy})
			}
		}
	}
	return
}

type Row []Dumbo
type Grid []Row

func (g Grid) At(p Point) *Dumbo {
	return &g[p.x][p.y]
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
	for _, neighbor := range p.Neighbors() {
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

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInputFile(filePath string) (grid Grid, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var row Row
		for _, r := range scanner.Text() {
			var number int
			number, err = strconv.Atoi(string(r))
			if err != nil {
				return
			}
			row = append(row, Dumbo{number, false})
		}
		grid = append(grid, row)
	}
	return
}

func main() {
	inputFile := "grid.txt"
	grid, err := parseInputFile(inputFile)
	handle(err)

	flashes := 0
	step := 0
	for !grid.IsSynchronized() {
		flashes += grid.Evolve()
		if step == 100 {
			fmt.Printf("%d flashes after 100 steps.\n", flashes)
		}
		step++
	}
	fmt.Printf("Synchronized after %d steps!\n", step)
}
