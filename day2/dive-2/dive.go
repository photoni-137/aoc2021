package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readCommands(r io.Reader) ([]Command, error) {
	scanner := bufio.NewScanner(r)
	var commands []Command
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		direction := words[0]
		amount, err := strconv.Atoi(words[1])
		if err != nil {
			return commands, err
		}
		commands = append(commands, Command{direction: direction, amount: amount})
	}
	return commands, scanner.Err()
}

func main() {
	inputFilePath := "commands.txt"
	reader, err := os.Open(inputFilePath)
	handle(err)
	defer reader.Close()

	commands, err := readCommands(reader)
	handle(err)

	sub := Submarine{position: 0, depth: 0}
	for _, command := range commands {
		sub.Execute(command)
	}
	fmt.Printf("The submarine is now at position %d and depth %d.\n This gives the product %d.\n", sub.position, sub.depth, sub.position*sub.depth)
}
