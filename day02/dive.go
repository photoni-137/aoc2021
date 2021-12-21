package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	direction string
	amount    int
}

type Submarine struct {
	position int
	depth    int
	aim      int
}

func (s *Submarine) Drive(command Command) {
	switch command.direction {
	case "forward":
		s.position += command.amount
	case "down":
		s.depth += command.amount
	case "up":
		s.depth -= command.amount
	}
}

func (s *Submarine) Execute(command Command) {
	switch command.direction {
	case "forward":
		s.position += command.amount
		s.depth += s.aim * command.amount
	case "down":
		s.aim += command.amount
	case "up":
		s.aim -= command.amount
	}
}

func parseCommands(lines []string) (commands []Command) {
	for _, line := range lines {
		words := strings.Split(line, " ")
		direction := words[0]
		amount, err := strconv.Atoi(words[1])
		shared.Handle(err)
		commands = append(commands, Command{direction: direction, amount: amount})
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	commands := parseCommands(lines)

	sub := Submarine{}
	for _, command := range commands {
		sub.Drive(command)
	}
	fmt.Printf("Interpreting the commands naively, the submarine ends up at position %d and depth %d.\n",
		sub.position, sub.depth)
	fmt.Printf("This gives the product %d.\n", sub.position*sub.depth)

	sub = Submarine{}
	for _, command := range commands {
		sub.Execute(command)
	}
	fmt.Printf("Interpreting the commands properly, the submarine ends up at position %d and depth %d.\n",
		sub.position, sub.depth)
	fmt.Printf("This gives the product %d.\n", sub.position*sub.depth)
}
