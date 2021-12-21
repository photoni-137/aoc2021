package main

import (
	"aoc2021/shared"
	"fmt"
	"sort"
	"strings"
)

type Polymer string
type Rules map[Polymer]Polymer

func (p Polymer) Pairs() (pairs []Polymer) {
	for i := range p {
		if i+1 >= len(p) {
			return
		}
		pairs = append(pairs, p[i:i+2])
	}
	return
}

func (p Polymer) Build(rules Rules) (new Polymer) {
	new = p
	for i, pair := range p.Pairs() {
		insertHere := 2*i + 1
		new = new[:insertHere] + rules[pair] + new[insertHere:]
	}
	return
}

func (p Polymer) Count() (freq map[Polymer]int) {
	freq = make(map[Polymer]int)
	for _, atom := range p {
		freq[Polymer(atom)]++
	}
	return
}

func (p Polymer) CountPairs() (freq map[Polymer]int) {
	freq = make(map[Polymer]int)
	for _, pair := range p.Pairs() {
		freq[pair]++
	}
	return
}

func parseLines(lines []string) (start Polymer, rules Rules) {
	rules = make(Rules)
	for _, line := range lines {
		switch {
		case strings.Contains(line, "->"):
			items := strings.Split(line, " -> ")
			pair, insert := Polymer(items[0]), Polymer(items[1])
			rules[pair] = insert
		case len(line) > 0:
			start = Polymer(line)
		}
	}
	return
}

func main() {
	lines, err := shared.ParseInputFile("input.txt")
	shared.Handle(err)
	start, rules := parseLines(lines)
	polymer := start
	for i := 0; i < 10; i++ {
		polymer = polymer.Build(rules)
	}
	frequencies := polymer.Count()
	var counts sort.IntSlice
	for _, count := range frequencies {
		counts = append(counts, count)
	}
	counts.Sort()
	least, most := counts[0], counts[len(counts)-1]
	fmt.Println(polymer.Count())
	fmt.Println(most - least)
}
