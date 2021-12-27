package main

import (
	"aoc2021/shared"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Packet struct {
	version, typeId, value int
	children               []*Packet
}

func (p *Packet) Evaluate() {
	switch p.typeId {
	case 0:
		p.value = p.Sum()
	case 1:
		p.value = p.Product()
	case 2:
		p.value = p.Minimum()
	case 3:
		p.value = p.Maximum()
	case 5:
		p.value = p.GreaterThan()
	case 6:
		p.value = p.LessThan()
	case 7:
		p.value = p.EqualTo()
	}
}

func (p Packet) Sum() int {
	sum := 0
	for _, child := range p.children {
		sum += child.value
	}
	return sum
}

func (p Packet) Product() int {
	product := 1
	for _, child := range p.children {
		product *= child.value
	}
	return product
}

func (p Packet) Minimum() int {
	minimum := math.MaxInt64
	for _, child := range p.children {
		if child.value < minimum {
			minimum = child.value
		}
	}
	return minimum
}

func (p Packet) Maximum() int {
	maximum := -1
	for _, child := range p.children {
		if child.value > maximum {
			maximum = child.value
		}
	}
	return maximum
}

func (p Packet) GreaterThan() int {
	if p.children[0].value > p.children[1].value {
		return 1
	}
	return 0
}

func (p Packet) LessThan() int {
	if p.children[0].value < p.children[1].value {
		return 1
	}
	return 0
}

func (p Packet) EqualTo() int {
	if p.children[0].value == p.children[1].value {
		return 1
	}
	return 0
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
	return bitsToInt(popped)
}

func bitsToInt(bits string) int {
	number, err := strconv.ParseInt(bits, 2, 64)
	shared.Handle(err)
	return int(number)
}

func popPacket(bits *string) (p Packet, success bool) {
	if isOnlyZeros(*bits) {
		return
	}

	version := popInt(bits, 3)
	typeId := popInt(bits, 3)
	p = Packet{version, typeId, 0, []*Packet{}}

	switch typeId {
	case 4:
		var sb strings.Builder
		stop := false
		for !stop {
			stop = popInt(bits, 1) == 0
			sb.WriteString(popString(bits, 4))
		}
		p.value = bitsToInt(sb.String())
	default:
		lengthType := popInt(bits, 1)
		switch lengthType {
		case 0:
			totalLength := popInt(bits, 15)
			bitsToParse := popString(bits, totalLength)
			for {
				child, found := popPacket(&bitsToParse)
				if !found {
					break
				}
				p.children = append(p.children, &child)
			}
		case 1:
			targets := popInt(bits, 11)
			for i := 0; i < targets; i++ {
				child, _ := popPacket(bits)
				p.children = append(p.children, &child)
			}
		}
	}

	(&p).Evaluate()
	success = true
	return
}

func versionSum(p Packet) (sum int) {
	queue := []*Packet{&p}
	for len(queue) > 0 {
		packet := queue[0]
		sum += packet.version
		queue = append(queue[1:], packet.children...)
	}
	return
}

func parseTransmission(lines []string) string {
	var sb strings.Builder
	line := lines[0]
	for _, r := range line {
		number, err := strconv.ParseInt(string(r), 16, 64)
		shared.Handle(err)
		sb.WriteString(fmt.Sprintf("%04b", number))
	}
	return sb.String()
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	transmission := parseTransmission(lines)

	var packets []Packet
	for {
		packet, found := popPacket(&transmission)
		if !found {
			break
		}
		packets = append(packets, packet)
	}

	for i, packet := range packets {
		fmt.Printf("Packet number %d has a version sum of %d and evaluates to %d.\n",
			i+1, versionSum(packet), packet.value)
	}
}
