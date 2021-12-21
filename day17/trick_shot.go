package main

import (
	"aoc2021/shared"
	"fmt"
	"math"
)

type Coordinates struct{ x, y int }

type Area struct{ lowerLeft, upperRight Coordinates }

func (a Area) Points() (points []Coordinates) {
	for x := a.lowerLeft.x; x <= a.upperRight.x; x++ {
		for y := a.lowerLeft.y; y <= a.upperRight.y; y++ {
			points = append(points, Coordinates{x, y})
		}
	}
	return
}

type Probe struct{ position, velocity Coordinates }

func (p *Probe) Travel() {
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	if p.velocity.x > 0 {
		p.velocity.x -= 1
	} else if p.velocity.x < 0 {
		p.velocity.x += 1
	}
	p.velocity.y -= 1
}

func (p *Probe) Hit(target Area) bool {
	for {
		switch {
		case p.In(target):
			return true
		case p.Behind(target):
			return false
		default:
			p.Travel()
		}
	}
}

func (p Probe) In(target Area) bool {
	pos := p.position
	rightX := pos.x >= target.lowerLeft.x && pos.x <= target.upperRight.x
	rightY := pos.y >= target.lowerLeft.y && pos.y <= target.upperRight.y
	return rightX && rightY
}

func (p Probe) Behind(target Area) bool {
	pos := p.position
	overX := pos.x > target.upperRight.x
	overY := pos.y < target.lowerLeft.y
	return overX || overY
}

func MinimumVelocity(target Area) (v Coordinates) {
	v.x = int(solveVx(target.lowerLeft)) + 1
	v.y = target.lowerLeft.y
	return
}

func MaximumVelocity(target Area) (v Coordinates) {
	v.x = target.upperRight.x
	v.y = -target.lowerLeft.y - 1
	return
}

func solveVx(target Coordinates) float64 {
	targetX := target.x
	return math.Sqrt(2*float64(targetX)-0.25) - 0.5
}

func maxHeight(target Area) int {
	return GaussSum(MaximumVelocity(target).y)
}

func GaussSum(n int) int {
	return n * (n + 1) / 2
}

func parseLines(lines []string) (target Area) {
	_, err := fmt.Sscanf(lines[0], "target area: x=%d..%d, y=%d..%d",
		&target.lowerLeft.x, &target.upperRight.x, &target.lowerLeft.y, &target.upperRight.y)
	shared.Handle(err)
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	target := parseLines(lines)

	var velocities []Coordinates
	start := Coordinates{0, 0}
	min, max := MinimumVelocity(target), MaximumVelocity(target)
	for vx := min.x; vx <= max.x; vx++ {
		for vy := min.y; vy <= max.y; vy++ {
			velocity := Coordinates{vx, vy}
			probe := Probe{start, velocity}
			if probe.Hit(target) {
				velocities = append(velocities, velocity)
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println("Maximum height: ", maxHeight(target))
	fmt.Println(len(velocities))
}
