package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

type Element struct {
	value   int
	nesting int
}

type SnailNumber []*Element

func (sn SnailNumber) Copy() SnailNumber {
	return append(SnailNumber(nil), sn...)
}

func (sn *SnailNumber) Add(other SnailNumber) {
	*sn = append(*sn, other...)
	for _, e := range *sn {
		e.nesting++
	}
	sn.Reduce()
}

func (sn *SnailNumber) Reduce() {
	reducible := true
	for reducible {
		reducible = sn.ReduceOnce()
	}
}

func (sn *SnailNumber) ReduceOnce() bool {
	exploded := sn.Explode()
	if exploded {
		return true
	}
	split := sn.Split()
	if split {
		return true
	}
	return false
}

func (sn *SnailNumber) Explode() bool {
	for i, element := range *sn {
		if element.nesting > 4 {
			sn.ExplodeAt(i)
			return true
		}
	}
	return false
}

func (sn *SnailNumber) Split() bool {
	for i, element := range *sn {
		if element.value > 9 {
			sn.SplitAt(i)
			return true
		}
	}
	return false
}

func (sn *SnailNumber) ExplodeAt(i int) {
	before := (*sn)[:i].Copy()
	target, after := (*sn)[i:i+2].Copy(), SnailNumber{}
	if len(*sn) > i+2 {
		after = (*sn)[i+2:].Copy()
	}
	if len(before) > 0 {
		before[i-1].value += target[0].value
	}
	if len(after) > 0 {
		after[0].value += target[1].value
	}
	exploded := &Element{0, target[0].nesting - 1}
	*sn = append(append(before, exploded), after...)
}

func (sn *SnailNumber) SplitAt(i int) {
	before, target, after := (*sn)[:i].Copy(), (*sn)[i], SnailNumber{}
	if len(*sn) > i+1 {
		after = (*sn)[i+1:].Copy()
	}
	split := SnailNumber{
		{target.value / 2, target.nesting + 1},
		{(target.value + 1) / 2, target.nesting + 1},
	}
	*sn = append(append(before, split...), after...)
}

func (sn SnailNumber) Magnitude() int {
	contracted := sn.Copy()
	for len(contracted) > 2 {
		for i := range contracted[:len(contracted)-1] {
			left, right := contracted[i], contracted[i+1]
			if left.nesting == right.nesting {
				substitute := Element{
					value:   pairMagnitude(left.value, right.value),
					nesting: left.nesting - 1,
				}
				contracted = append(append(contracted[:i], &substitute), contracted[i+2:]...)
				break
			}
		}
	}
	return pairMagnitude(contracted[0].value, contracted[1].value)
}

func pairMagnitude(left, right int) int {
	return 3*left + 2*right
}

func (sn SnailNumber) Pretty() string {
	var sb strings.Builder
	for _, element := range sn {
		sb.WriteString(fmt.Sprintf("| %d (%d) |", element.value, element.nesting))
	}
	return sb.String()
}

func parseSnailNumber(s string) (sn SnailNumber) {
	var nesting int
	for _, r := range s {
		switch r {
		case '[':
			nesting++
		case ']':
			nesting--
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			value, _ := strconv.Atoi(string(r))
			sn = append(sn, &Element{value, nesting})
		}
	}
	return
}

func linesToNumbers(lines []string) (numbers []SnailNumber) {
	numbers = make([]SnailNumber, len(lines))
	for i, line := range lines {
		numbers[i] = parseSnailNumber(line)
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	numbers := linesToNumbers(lines)

	sum := numbers[0].Copy()
	for _, number := range numbers[1:] {
		sum.Add(number)
	}
	fmt.Println("All summed up!")
	fmt.Println("Magnitude:", sum.Magnitude())

	highestSum := 0
	for i, thisLine := range lines[:len(lines)-1] {
		for j, thatLine := range lines[i+1:] {
			fmt.Println(i, j)
			snailSum := parseSnailNumber(thisLine)
			fmt.Println(snailSum.Pretty())
			snailSum.Add(parseSnailNumber(thatLine))
			magnitude := snailSum.Magnitude()
			if magnitude > highestSum {
				highestSum = magnitude
			}
		}
	}
	fmt.Println("Highest possible sum of two snail numbers:", highestSum)
}
