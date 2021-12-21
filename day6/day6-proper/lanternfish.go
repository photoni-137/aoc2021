package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInputFile(filePath string) (timers []int, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		timers, err = parseAges(scanner.Text())
	}
	return
}

func parseAges(input string) (timers []int, err error) {
	for _, s := range strings.Split(input, ",") {
		var timer int
		timer, err = strconv.Atoi(s)
		if err != nil {
			return
		}
		timers = append(timers, timer)
	}
	return
}

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

func main() {
	inputFile := "lanternfish.txt"
	timers, err := parseInputFile(inputFile)
	handle(err)

	limits := [2]int{80, 256}
	for _, limit := range limits {
		school := make(School)
		for _, timer := range timers {
			school[timer]++
		}
		for day := 0; day < limit; day++ {
			school.evolve()
		}
		fmt.Printf("Behind %d days, we have %d lanternfish.\n", limit, school.count())
	}
}
