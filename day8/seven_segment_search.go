package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Signal []rune

func (s Signal) Contains(target rune) (contains bool) {
	for _, r := range s {
		if r == target {
			contains = true
			return
		}
	}
	return
}

type Code map[string]Signal
type Cipher map[string]string
type Wiring map[rune]rune
type Frequency map[rune]int

func main() {
	inputFile := "signals.txt"
	inputs, outputs, err := parseInputFile(inputFile)
	handle(err)

	count := countSimpleOutputs(outputs)
	fmt.Println(count)

	numbers, err := decipherOutputs(inputs, outputs)
	handle(err)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}

func decipherOutputs(inputs [][]Signal, outputs [][]Signal) (numbers []int, err error) {
	if len(inputs) != len(outputs) {
		err = errors.New("input and output have different length")
		return
	}
	for i, in := range inputs {
		code := solve(in)
		cipher := invert(code)
		var sb strings.Builder
		for _, signal := range outputs[i] {
			sb.WriteString(cipher[string(signal)])
		}
		var number int
		number, err = strconv.Atoi(sb.String())
		if err != nil {
			return
		}
		numbers = append(numbers, number)
	}
	return
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func countSimpleOutputs(outputs [][]Signal) int {
	count := 0
	for _, out := range outputs {
		for _, signal := range out {
			switch len(signal) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}

func invert(code Code) (cipher Cipher) {
	cipher = make(Cipher)
	for digit, signal := range code {
		cipher[string(signal)] = digit
	}
	return
}

func parseInputFile(filePath string) (inputs [][]Signal, outputs [][]Signal, err error) {
	reader, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		halves := strings.Split(line, " | ")
		rawInput, rawOutput := halves[0], halves[1]
		input, output := stringToSignals(rawInput), stringToSignals(rawOutput)
		inputs, outputs = append(inputs, input), append(outputs, output)
	}
	return
}

func frequency(signals []Signal) (freq Frequency) {
	freq = make(Frequency)
	for _, signal := range signals {
		for _, r := range signal {
			freq[r]++
		}
	}
	return
}

func solve(input []Signal) (code Code) {
	code = make(Code)
	var fivers, sixers []Signal
	for _, signal := range input {
		switch len(signal) {
		case 2:
			code["1"] = signal
		case 3:
			code["7"] = signal
		case 4:
			code["4"] = signal
		case 5:
			fivers = append(fivers, signal)
		case 6:
			sixers = append(sixers, signal)
		case 7:
			code["8"] = signal
		}
	}
	wire := make(Wiring)
	for r, freq := range frequency(input) {
		switch freq {
		case 4:
			wire['e'] = r
		case 6:
			wire['b'] = r
		case 8:
			if code["1"].Contains(r) {
				wire['c'] = r
			}
		case 9:
			wire['f'] = r
		}
	}
	var remainingFivers []Signal
	for _, signal := range fivers {
		if !signal.Contains(wire['f']) {
			code["2"] = signal
		} else {
			remainingFivers = append(remainingFivers, signal)
		}
	}
	for _, signal := range remainingFivers {
		if signal.Contains(wire['b']) {
			code["5"] = signal
		} else {
			code["3"] = signal
		}
	}
	var remainingSixers []Signal
	for _, signal := range sixers {
		if !signal.Contains(wire['e']) {
			code["9"] = signal
		} else {
			remainingSixers = append(remainingSixers, signal)
		}
	}
	for _, signal := range remainingSixers {
		if signal.Contains(wire['c']) {
			code["0"] = signal
		} else {
			code["6"] = signal
		}
	}
	return
}

func stringToSignals(rawInput string) (signals []Signal) {
	for _, signalString := range strings.Split(rawInput, " ") {
		signal := Signal(signalString)
		sort.Slice(signal, func(i, j int) bool { return signal[i] < signal[j] })
		signals = append(signals, signal)
	}
	return
}
