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

func evolve(timers []int) (newTimers []int) {
	newTimers = make([]int, len(timers))
	for i, timer := range timers {
		if timer == 0 {
			newTimers = append(newTimers, 8)
			newTimers[i] = 6
		} else {
			newTimers[i] = timer - 1
		}
	}
	return
}

func spawn(daysSinceBirth int, cache map[int]int) (spawned int) {
	spawned, found := cache[daysSinceBirth]
	if found {
		return
	}
	for day := daysSinceBirth - 9; day >= 0; day -= 7 {
		spawned += 1 + spawn(day, cache)
	}
	cache[daysSinceBirth] = spawned
	return
}

func main() {
	inputFile := "lanternfish.txt"
	timers, err := parseInputFile(inputFile)
	handle(err)

	limit := 80
	for day := 0; day < limit; day++ {
		timers = evolve(timers)
	}
	fmt.Printf("Behind %d days, we have %d lanternfishes.\n", limit, len(timers))
	dummy := []int{8}
	dummyCache := make(map[int]int)
	for day := 0; day < limit; day++ {
		dummyCache[day] = len(dummy) - 1
		dummy = evolve(dummy)
	}

	fmt.Println("Let's do it recursive now!")
	limit = 256
	timers, _ = parseInputFile(inputFile)
	cache := fillCache(limit)
	fishes := 0
	for _, birthday := range Birthdays(timers, limit) {
		fishes += 1 + spawn(birthday, cache)
	}
	fmt.Printf("Behind %d days, we have %d lanternfishes.\n", limit, fishes)
}

func Birthdays(timers []int, limit int) []int {
	birthdays := make([]int, len(timers))
	for i, timer := range timers {
		birthdays[i] = limit + 8 - timer
	}
	return birthdays
}

func fillCache(limit int) map[int]int {
	cache := make(map[int]int)
	for day := 0; day < limit; day++ {
		spawn(day, cache)
	}
	return cache
}
