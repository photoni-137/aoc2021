package main

import (
	"aoc2021/shared"
	"fmt"
	"sort"
)

var corruptScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incompleteScore = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var Closing = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

type Stack []rune

func (s Stack) Push(r rune) Stack {
	return append(s, r)
}

func (s Stack) Pop() (rune, Stack) {
	final := len(s) - 1
	return s[final], s[:final]
}

func isOpening(r rune) bool {
	for bracket := range Closing {
		if r == bracket {
			return true
		}
	}
	return false
}

func isClosing(r rune) bool {
	for _, bracket := range Closing {
		if r == bracket {
			return true
		}
	}
	return false
}

func sum(numbers sort.IntSlice) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func middle(numbers sort.IntSlice) int {
	sort.Sort(numbers)
	return numbers[len(numbers)/2]
}

func validate(line string) (status string, score int) {
	var stack Stack
	for _, r := range line {
		if isOpening(r) {
			stack = stack.Push(r)
		} else if isClosing(r) {
			var currentOpen rune
			currentOpen, stack = stack.Pop()
			if r != Closing[currentOpen] {
				status = "corrupted"
				score = corruptScore[r]
				return
			}
		}
	}

	if len(stack) == 0 {
		status = "valid"
	} else {
		status = "incomplete"
		score = scoreIncomplete(stack)
	}
	return
}

func scoreIncomplete(stack Stack) (score int) {
	for len(stack) > 0 {
		var r rune
		r, stack = stack.Pop()
		score *= 5
		score += incompleteScore[r]
	}
	return
}

func main() {
	lines := shared.ParseInputFile("input.txt")

	scores := make(map[string]sort.IntSlice)
	for _, line := range lines {
		status, score := validate(line)
		scores[status] = append(scores[status], score)
	}

	fmt.Printf("Corrupted lines score for %d points, incomplete lines for %d points.\n",
		sum(scores["corrupted"]), middle(scores["incomplete"]))
}
