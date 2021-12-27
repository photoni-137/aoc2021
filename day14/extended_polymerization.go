package main

import (
	"aoc2021/shared"
	"fmt"
	"sort"
	"strings"
)

type AtomPair string

type Polymer struct {
	start, end string
	pairs      map[AtomPair]int
}

type InsertRules map[AtomPair]string

func NewPolymer(s string) Polymer {
	start, end := string(s[0]), string(s[len(s)-1])
	pairs := stringToPair(s)
	return Polymer{start, end, pairs}
}

func stringToPair(s string) (pairs map[AtomPair]int) {
	pairs = make(map[AtomPair]int)
	for i := range s {
		if i+1 >= len(s) {
			return
		}
		pair := AtomPair(s[i : i+2])
		pairs[pair]++
	}
	return
}

func (ap AtomPair) Insert(rules InsertRules) []AtomPair {
	insert := rules[ap]
	return []AtomPair{
		AtomPair(string(ap[0]) + insert),
		AtomPair(insert + string(ap[1])),
	}
}

func (p Polymer) Build(rules InsertRules) Polymer {
	newPairs := make(map[AtomPair]int)
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

func parseLines(lines []string) (start Polymer, rules InsertRules) {
	rules = make(InsertRules)
	for _, line := range lines {
		switch {
		case strings.Contains(line, "->"):
			items := strings.Split(line, " -> ")
			pair, insert := AtomPair(items[0]), items[1]
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
			fmt.Printf("After %d days, the polymer scores for %d points.\n", i, polymer.Score())
		}
		polymer = polymer.Build(rules)
	}
	fmt.Printf("After %d days, the polymer scores for %d points. Wew lad!\n", 40, polymer.Score())
}
