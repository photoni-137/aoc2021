package main

import (
	"aoc2021/shared"
	"fmt"
	"strings"
)

type Path []string
type ExtendedPath struct {
	path      Path
	exception bool
}
type Caves map[string]*Cave
type Cave struct {
	neighbors []string
	isSmall   bool
}

func (c Caves) addIfNew(names []string) {
	for _, name := range names {
		if _, found := c[name]; !found {
			c[name] = newCave(name)
		}
	}
}

func newCave(name string) (cave *Cave) {
	cave = new(Cave)
	cave.isSmall = isLowercase(name)
	return
}

func isLowercase(s string) bool {
	if strings.ToLower(s) == s {
		return true
	}
	return false
}

func (c Caves) connect(thisName, thatName string) {
	this, that := c[thisName], c[thatName]
	this.neighbors = append(this.neighbors, thatName)
	that.neighbors = append(that.neighbors, thisName)
}

func (c Caves) constructPaths() (finishedPaths []Path) {
	var unfinishedPaths []Path
	unfinishedPaths = append(unfinishedPaths, Path{"start"})
	for len(unfinishedPaths) > 0 {
		var newPaths []Path
		for _, path := range unfinishedPaths {
			currentName := path[len(path)-1]
			for _, neighbor := range c[currentName].neighbors {
				if c.validPath(path, neighbor) {
					newPath := append(Path(nil), path...)
					newPath = append(newPath, neighbor)
					if neighbor == "end" {
						finishedPaths = append(finishedPaths, newPath)
					} else {
						newPaths = append(newPaths, newPath)
					}
				}
			}
		}
		unfinishedPaths = append([]Path(nil), newPaths...)
	}
	return
}

func (c Caves) validPath(p Path, target string) bool {
	if !c[target].isSmall {
		return true
	}
	for _, name := range p {
		if name == target {
			return false
		}
	}
	return true
}

func (c Caves) constructExtendedPaths() (finishedPaths []ExtendedPath) {
	var unfinishedPaths []ExtendedPath
	unfinishedPaths = append(unfinishedPaths, ExtendedPath{Path{"start"}, false})
	for len(unfinishedPaths) > 0 {
		var newPaths []ExtendedPath
		for _, path := range unfinishedPaths {
			currentName := path.path[len(path.path)-1]
			for _, neighbor := range c[currentName].neighbors {
				if isValid, usedException := c.validExtendedPath(path, neighbor); isValid {
					newPath := ExtendedPath{append(Path(nil), path.path...), usedException}
					newPath.path = append(newPath.path, neighbor)
					if neighbor == "end" {
						finishedPaths = append(finishedPaths, newPath)
					} else {
						newPaths = append(newPaths, newPath)
					}
				}
			}
		}
		unfinishedPaths = append([]ExtendedPath(nil), newPaths...)
	}
	return
}

func (c Caves) validExtendedPath(p ExtendedPath, target string) (isValid, usedException bool) {
	usedException = p.exception
	if !c[target].isSmall {
		isValid = true
		return
	}
	switch target {
	case "start":
		return
	case "end":
		isValid = true
		return
	default:
		for _, name := range p.path {
			if name == target {
				if usedException {
					return
				}
				usedException = true
			}
		}
	}
	isValid = true
	return
}

func linesToCaves(lines []string) Caves {
	caves := make(Caves)
	for _, line := range lines {
		names := strings.Split(line, "-")
		from, to := names[0], names[1]
		caves.addIfNew([]string{from, to})
		caves.connect(from, to)
	}
	return caves
}

func main() {
	lines := shared.ParseInputFile("paths.txt")
	caves := linesToCaves(lines)

	paths := caves.constructPaths()
	fmt.Printf("Found %d valid paths.\n", len(paths))

	extendedPaths := caves.constructExtendedPaths()
	fmt.Printf("Extending the rules, we found %d valid paths.\n", len(extendedPaths))
}
