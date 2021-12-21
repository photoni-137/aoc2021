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
	length    int
}

type Submarine struct {
	position int
	depth    int
}

func (s *Submarine) Drive(command Command) {
	switch command.direction {
	case "forward":
		s.position += command.length
	case "down":
		s.depth += command.length
	case "up":
		s.depth -= command.length
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
		length, err := strconv.Atoi(words[1])
		if err != nil {
			return commands, err
		}
		commands = append(commands, Command{direction: direction, length: length})
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
		sub.Drive(command)
	}
	fmt.Printf("The submarine is now at position %d and depth %d.\n This gives the product %d.\n", sub.position, sub.depth, sub.position*sub.depth)
}
