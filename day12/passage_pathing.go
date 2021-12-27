package main

import (
	"aoc2021/shared"
	"fmt"
	"strings"
)

type Path []string

type ExtendedPath struct {
	path          Path
	usedException bool
}

type Cave struct {
	neighbors []string
	isLarge   bool
}

type CaveSystem map[string]*Cave

func (p Path) Current() string {
	return p[len(p)-1]
}

func (p Path) Copy() Path {
	return append(Path(nil), p...)
}

func (cs CaveSystem) AddIfNew(names []string) {
	for _, name := range names {
		if _, found := cs[name]; !found {
			cs[name] = newCave(name)
		}
	}
}

func (cs CaveSystem) Connect(this, that string) {
	cs[this].neighbors = append(cs[this].neighbors, that)
	cs[that].neighbors = append(cs[that].neighbors, this)
}

func (cs CaveSystem) ConstructPaths() (finishedPaths []Path) {
	unfinishedPaths := []Path{{"start"}}
	for len(unfinishedPaths) > 0 {
		var newPaths []Path
		for _, p := range unfinishedPaths {
			current := p.Current()
			for _, neighbor := range cs[current].neighbors {
				if cs.ValidPath(p, neighbor) {
					newPath := append(p.Copy(), neighbor)
					if neighbor == "end" {
						finishedPaths = append(finishedPaths, newPath)
					} else {
						newPaths = append(newPaths, newPath)
					}
				}
			}
		}
		unfinishedPaths = newPaths
	}
	return
}

func (cs CaveSystem) ValidPath(p Path, target string) bool {
	if cs[target].isLarge {
		return true
	}
	for _, name := range p {
		if name == target {
			return false
		}
	}
	return true
}

func (cs CaveSystem) ConstructExtendedPaths() (finishedPaths []ExtendedPath) {
	unfinishedPaths := []ExtendedPath{{Path{"start"}, false}}
	for len(unfinishedPaths) > 0 {
		var newPaths []ExtendedPath
		for _, ep := range unfinishedPaths {
			current := ep.path.Current()
			for _, neighbor := range cs[current].neighbors {
				isValid, usedException := cs.ValidExtendedPath(ep, neighbor)
				if isValid {
					newPath := ExtendedPath{
						path:          append(ep.path.Copy(), neighbor),
						usedException: usedException,
					}
					if neighbor == "end" {
						finishedPaths = append(finishedPaths, newPath)
					} else {
						newPaths = append(newPaths, newPath)
					}
				}
			}
		}
		unfinishedPaths = newPaths
	}
	return
}

func (cs CaveSystem) ValidExtendedPath(ep ExtendedPath, target string) (isValid, usedException bool) {
	usedException = ep.usedException
	if target == "start" {
		return
	}
	if cs[target].isLarge || target == "end" {
		isValid = true
		return
	}
	for _, name := range ep.path {
		if name == target {
			if usedException {
				return
			}
			usedException = true
		}
	}
	isValid = true
	return
}

func newCave(name string) (cave *Cave) {
	cave = new(Cave)
	cave.isLarge = isUppercase(name)
	return
}

func isUppercase(s string) bool {
	if strings.ToUpper(s) == s {
		return true
	}
	return false
}

func parseCaveSystem(lines []string) CaveSystem {
	caves := make(CaveSystem)
	for _, line := range lines {
		names := strings.Split(line, "-")
		from, to := names[0], names[1]
		caves.AddIfNew([]string{from, to})
		caves.Connect(from, to)
	}
	return caves
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	caveSystem := parseCaveSystem(lines)

	paths := caveSystem.ConstructPaths()
	fmt.Printf("Found %d valid paths.\n", len(paths))

	extendedPaths := caveSystem.ConstructExtendedPaths()
	fmt.Printf("Extending the rules, we found %d valid paths.\n", len(extendedPaths))
}
