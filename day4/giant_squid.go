package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

type bingoField struct {
	value int
	hit   bool
}

type bingoBoard [5][5]bingoField

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

func newBingoBoard(numbers [][]int) (b bingoBoard) {
	for i, row := range numbers {
		for j, number := range row {
			b[i][j].value = number
		}
	}
	return
}

func parseInputFile(filePath string) (numbers []int, boards []bingoBoard, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer reader.Close()

	lines := readLines(reader)
	stripped := removeEmptyLines(lines)
	numbers = getNumbersFromLine(stripped[0], ",")

	boardLines := stripped[1:]
	if len(boardLines)%5 != 0 {
		err = errors.New("garbage lines in input file")
		return
	}
	for i := 0; i <= len(boardLines)-5; i += 5 {
		nextFiveLines := boardLines[i : i+5]
		boards = append(boards, getBoardFromLines(nextFiveLines))
	}
	return
}

func readLines(reader io.Reader) (lines []string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func getNumbersFromLine(line string, separator string) (numbers []int) {
	numberStrings := strings.Split(line, separator)
	for _, s := range numberStrings {
		number, err := strconv.Atoi(s)
		if err == nil {
			numbers = append(numbers, number)
		}
	}
	return
}

func getBoardFromLines(lines []string) (b bingoBoard) {
	var rows [][]int
	for _, line := range lines {
		row := getNumbersFromLine(line, " ")
		rows = append(rows, row)
	}
	b = newBingoBoard(rows)
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
	inputFile := "bingo.txt"
	numbers, boards, err := parseInputFile(inputFile)
	handle(err)

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
