package main

import (
	"aoc2021/shared"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var FrequencyCode = map[Signal]string{
	"89":      "1",
	"889":     "7",
	"6789":    "4",
	"47788":   "2",
	"67789":   "5",
	"77889":   "3",
	"467789":  "6",
	"467889":  "0",
	"677889":  "9",
	"4677889": "8",
}

type Signal string

type Display struct{ cipher, code []Signal }

func stringToSignal(s string) Signal {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return Signal(runes)
}

func (d Display) frequency() (freq map[rune]int) {
	freq = make(map[rune]int)
	for _, signal := range d.cipher {
		for _, r := range signal {
			freq[r]++
		}
	}
	return
}

func (d Display) countSimpleSignals() (count int) {
	for _, s := range d.code {
		switch len(s) {
		case 2, 3, 4, 7:
			count++
		}
	}
	return
}

func (d Display) decode() (digits string) {
	frequency := d.frequency()
	for _, signal := range d.code {
		id := identify(frequency, signal)
		digit := FrequencyCode[id]
		digits += digit
	}
	return
}

func identify(frequency map[rune]int, s Signal) (id Signal) {
	var frequencies string
	for _, r := range s {
		frequencies += strconv.Itoa(frequency[r])
	}
	return stringToSignal(frequencies)
}

func parseDisplays(lines []string) (displays []Display) {
	for _, line := range lines {
		halves := strings.Split(line, " | ")
		cipher, code := parseSignals(halves[0]), parseSignals(halves[1])
		displays = append(displays, Display{cipher, code})
	}
	return
}

func parseSignals(line string) (signals []Signal) {
	for _, s := range strings.Split(line, " ") {
		signal := stringToSignal(s)
		signals = append(signals, signal)
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	displays := parseDisplays(lines)

	sum, simpleSignals := 0, 0
	for _, display := range displays {
		simpleSignals += display.countSimpleSignals()
		number, _ := strconv.Atoi(display.decode())
		sum += number
	}
	fmt.Printf("%d digits could be decoded effortlessly.\n", simpleSignals)
	fmt.Printf("The sum of all digits is %d.\n", sum)
}
