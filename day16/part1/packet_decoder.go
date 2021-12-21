package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

type Packet struct{ version, typeId, content int }

func parseLines(lines []string) string {
	var sb strings.Builder
	line := lines[0]
	for _, r := range line {
		number, err := strconv.ParseInt(string(r), 16, 64)
		shared.Handle(err)
		sb.WriteString(fmt.Sprintf("%04b", number))
	}
	return sb.String()
}

func popPackets(bits *string) (packets []Packet) {
	if isOnlyZeros(*bits) {
		return
	}

	fmt.Println((*bits)[0:3], (*bits)[3:6])
	version := popInt(bits, 3)
	typeId := popInt(bits, 3)
	packets = []Packet{{version, typeId, 0}}
	fmt.Println(packets[0])

	switch typeId {
	case 4:
		var sb strings.Builder
		stop := false
		for !stop {
			stop = popInt(bits, 1) == 0
			sb.WriteString(popString(bits, 4))
		}
		packets[0].content = toInt(sb.String())
	default:
		var subpackets []Packet
		lengthType := popInt(bits, 1)
		switch lengthType {
		case 0:
			totalLength := popInt(bits, 15)
			bitsToParse := popString(bits, totalLength)
			for {
				found := popPackets(&bitsToParse)
				if len(found) == 0 {
					break
				}
				subpackets = append(subpackets, found...)
			}
		case 1:
			targets := popInt(bits, 11)
			for i := 0; i < targets; i++ {
				found := popPackets(bits)
				subpackets = append(subpackets, found...)
			}
		}
		packets = append(packets, subpackets...)
	}
	return
}

func isOnlyZeros(s string) bool {
	for _, r := range s {
		if r != '0' {
			return false
		}
	}
	return true
}

func popString(bits *string, n int) string {
	popped := (*bits)[:n]
	*bits = (*bits)[n:]
	return popped
}

func popInt(bits *string, n int) int {
	popped := popString(bits, n)
	return toInt(popped)
}

func toInt(bits string) int {
	number, err := strconv.ParseInt(bits, 2, 64)
	shared.Handle(err)
	return int(number)
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	bitstream := parseLines(lines)

	var packets []Packet
	for {
		found := popPackets(&bitstream)
		if len(found) == 0 {
			break
		}
		packets = append(packets, found...)
	}

	sum := 0
	for _, p := range packets {
		sum += p.version
	}
	fmt.Println(sum)
}
