package main

import (
	"aoc2021/shared"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type bingoField struct {
	value int
	hit   bool
}

type bingoBoard [5][5]bingoField

func newBingoBoard(numbers [][]int) (b bingoBoard) {
	for i, row := range numbers {
		for j, number := range row {
			b[i][j].value = number
		}
	}
	return
}

func (b *bingoBoard) playGame(numbers []int) (won bool, score int) {
	for _, number := range numbers {
		won, score = b.playRound(number)
		if won {
			return
		}
	}
	return
}

func (b *bingoBoard) playRound(number int) (won bool, score int) {
	if b.markIfHit(number) {
		if b.checkForWin() {
			won = true
			score = b.sumUnmarkedFields() * number
		}
	}
	return
}

func (b *bingoBoard) markIfHit(target int) (hit bool) {
	for i, row := range b {
		for j := range row {
			if b[i][j].value == target {
				b[i][j].hit = true
				hit = true
				return
			}
		}
	}
	return
}

func (b bingoBoard) checkForWin() (won bool) {
	for i := range b {
		if b.checkRowForWin(i) || b.checkColumnForWin(i) {
			won = true
			return
		}
	}
	return
}

func (b bingoBoard) sumUnmarkedFields() (sum int) {
	for _, row := range b {
		for _, field := range row {
			if !field.hit {
				sum += field.value
			}
		}
	}
	return
}

func (b bingoBoard) checkRowForWin(rowNumber int) (won bool) {
	won = true
	for _, field := range b[rowNumber] {
		won = won && field.hit
	}
	return
}

func (b bingoBoard) checkColumnForWin(columnNumber int) (won bool) {
	won = true
	for _, row := range b {
		won = won && row[columnNumber].hit
	}
	return
}

func parseLines(lines []string) (numbers []int, boards []bingoBoard) {
	stripped := removeEmptyLines(lines)
	numbers = parseNumbers(stripped[0], ",")

	remaining := stripped[1:]
	if len(remaining)%5 != 0 {
		shared.Handle(errors.New("garbage lines in input file"))
		return
	}
	for i := 0; i <= len(remaining)-5; i += 5 {
		boardLines := remaining[i : i+5]
		boards = append(boards, getBoardFromLines(boardLines))
	}
	return
}

func getBoardFromLines(lines []string) (b bingoBoard) {
	var rows [][]int
	for _, line := range lines {
		row := parseNumbers(line, " ")
		rows = append(rows, row)
	}
	b = newBingoBoard(rows)
	return
}

func parseNumbers(line, separator string) (numbers []int) {
	line = strings.Trim(line, " ")
	line = strings.ReplaceAll(line, "  ", " ")
	numberStrings := strings.Split(line, separator)
	for _, s := range numberStrings {
		number, err := strconv.Atoi(s)
		shared.Handle(err)
		numbers = append(numbers, number)
	}
	return
}

func removeEmptyLines(lines []string) (stripped []string) {
	for _, line := range lines {
		if line != "" {
			stripped = append(stripped, line)
		}
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	numbers, boards := parseLines(lines)

	fmt.Printf("Playing %d numbers for %d boards.\n", len(numbers), len(boards))
	remaining := make([]int, len(boards))
	for i := range remaining {
		remaining[i] = i
	}
	for round, number := range numbers {
		var notHit []int
		for _, boardIndex := range remaining {
			won, score := boards[boardIndex].playRound(number)
			if won {
				fmt.Printf("Board %d won after %d rounds with a score of %d points.\n", boardIndex, round+1, score)
			} else {
				notHit = append(notHit, boardIndex)
			}
		}
		remaining = append([]int(nil), notHit...)
		if len(remaining) == 0 {
			return
		}
	}
	return
}
