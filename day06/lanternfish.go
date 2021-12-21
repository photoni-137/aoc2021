package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

type School map[int]int

func (s School) evolve() {
	birthing := s[0]
	for timer := 0; timer < 8; timer++ {
		s[timer] = s[timer+1]
	}
	s[6] += birthing
	s[8] = birthing
}

func (s School) count() (sum int) {
	for _, fish := range s {
		sum += fish
	}
	return
}

func parseTimers(line string) (timers []int) {
	timers = parseAges(line)
	return
}

func parseAges(input string) (timers []int) {
	for _, s := range strings.Split(input, ",") {
		timer, err := strconv.Atoi(s)
		shared.Handle(err)
		timers = append(timers, timer)
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	timers := parseTimers(lines[0])

	limits := []int{80, 256}
	for _, limit := range limits {
		school := make(School)
		for _, timer := range timers {
			school[timer]++
		}
		for day := 0; day < limit; day++ {
			school.evolve()
		}
		fmt.Printf("After %d days, we have %d lanternfish.\n", limit, school.count())
	}
}
