package main

import (
	"aoc2021/shared"
	"fmt"
	"sort"
	"strings"
)

type Pair string
type Polymer struct {
	start, end string
	pairs      map[Pair]int
}
type Rules map[Pair]string

func NewPolymer(s string) Polymer {
	start, end := string(s[0]), string(s[len(s)-1])
	pairs := toPairs(s)
	return Polymer{start, end, pairs}
}

func toPairs(s string) (pairs map[Pair]int) {
	pairs = make(map[Pair]int)
	for i := range s {
		if i+1 >= len(s) {
			return
		}
		pair := Pair(s[i : i+2])
		pairs[pair]++
	}
	return
}

func (p Pair) Insert(rules Rules) []Pair {
	insert := rules[p]
	return []Pair{
		Pair(string(p[0]) + insert),
		Pair(insert + string(p[1])),
	}
}

func (p Polymer) Build(rules Rules) Polymer {
	newPairs := make(map[Pair]int)
	for pair, count := range p.pairs {
		for _, newPair := range pair.Insert(rules) {
			newPairs[newPair] += count
		}
	}
	return Polymer{p.start, p.end, newPairs}
}

func (p Polymer) Score() int {
	var counts sort.IntSlice
	for _, count := range p.Atoms() {
		counts = append(counts, count)
	}
	counts.Sort()
	least, most := counts[0], counts[len(counts)-1]
	return most - least
}

func (p Polymer) Atoms() map[string]int {
	atoms := make(map[string]int)
	for pair, count := range p.pairs {
		for _, atom := range pair {
			atoms[string(atom)] += count
		}
	}
	atoms[p.start]++
	atoms[p.end]++
	for atom := range atoms {
		atoms[atom] /= 2
	}
	return atoms
}

func parseLines(lines []string) (start Polymer, rules Rules) {
	rules = make(Rules)
	for _, line := range lines {
		switch {
		case strings.Contains(line, "->"):
			items := strings.Split(line, " -> ")
			pair, insert := Pair(items[0]), items[1]
			rules[pair] = insert
		case len(line) > 0:
			start = NewPolymer(line)
		}
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	start, rules := parseLines(lines)
	polymer := start
	for i := 0; i < 40; i++ {
		if i == 10 {
			fmt.Println(polymer.Score())
		}
		polymer = polymer.Build(rules)
	}
	fmt.Println(polymer.Score())
}
