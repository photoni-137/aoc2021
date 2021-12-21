package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const maxHeight = 9

type Point struct{ x, y int }
type PointSet []Point
type Heightmap [][]int

func (p Point) Neighbors() PointSet {
	x, y := p.x, p.y
	return PointSet{
		{x: x - 1, y: y},
		{x: x + 1, y: y},
		{x: x, y: y - 1},
		{x: x, y: y + 1},
	}
}

func (s PointSet) Contains(p Point) bool {
	for _, point := range s {
		if point == p {
			return true
		}
	}
	return false
}

func (h Heightmap) At(p Point) int {
	x, y := p.x, p.y
	if x < 0 || x >= len(h) {
		return maxHeight
	}
	if y < 0 || y >= len(h[0]) {
		return maxHeight
	}
	return h[x][y]
}

func main() {
	inputFile := "floor.txt"
	floor, err := parseInputFile(inputFile)
	handle(err)

	lowPoints, risk := floor.identifyLowPoints()
	fmt.Println(risk)

	var basinSizes sort.IntSlice
	for _, lowPoint := range lowPoints {
		size := floor.basinSize(lowPoint)
		fmt.Println(lowPoint, size)
		basinSizes = append(basinSizes, size)
	}
	sort.Sort(sort.Reverse(basinSizes))
	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}

func (h Heightmap) identifyLowPoints() (lowPoints PointSet, riskLevel int) {
	for x, row := range h {
		for y, height := range row {
			point := Point{x: x, y: y}
			isLower := true
			for _, neighbor := range point.Neighbors() {
				isLower = isLower && height < h.At(neighbor)
			}
			if isLower {
				lowPoints = append(lowPoints, point)
				riskLevel += height + 1
			}
		}
	}
	return
}

func (h Heightmap) basinSize(lowPoint Point) int {
	basin := PointSet(nil)
	newPoints := PointSet{lowPoint}
	for len(newPoints) > 0 {
		basin = append(basin, newPoints...)
		newPoints = h.scan(newPoints, basin)
	}
	return len(basin)
}

func (h Heightmap) scan(startPoints PointSet, basin PointSet) (newPoints PointSet) {
	for _, point := range startPoints {
		for _, neighbor := range point.Neighbors() {
			if h.At(neighbor) < maxHeight && !basin.Contains(neighbor) && !newPoints.Contains(neighbor) {
				newPoints = append(newPoints, neighbor)
			}
		}
	}
	return
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInputFile(filePath string) (floor Heightmap, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer reader.Close()

	var lines []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	floor = make(Heightmap, len(lines))
	for i, line := range lines {
		floor[i] = make([]int, len(line))
		for j, r := range line {
			var height int
			height, err = strconv.Atoi(string(r))
			if err != nil {
				return
			}
			floor[i][j] = height
		}
	}
	return
}
