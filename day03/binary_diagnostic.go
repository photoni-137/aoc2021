package main

import (
	"aoc2021/shared"
	"fmt"
	"strconv"
	"strings"
)

type bitSlice []int

func bitSliceToInt(bin bitSlice) int {
	var sb strings.Builder
	for _, bit := range bin {
		sb.WriteString(strconv.Itoa(bit))
	}
	number, err := strconv.ParseInt(sb.String(), 2, 64)
	shared.Handle(err)
	return int(number)
}

func parseBinaryStrings(lines []string) (bins []bitSlice) {
	for _, line := range lines {
		bin := make(bitSlice, len(line))
		for i, r := range line {
			number, err := strconv.Atoi(string(r))
			shared.Handle(err)
			bin[i] = number
		}
		bins = append(bins, bin)
	}
	return
}

func extractRates(bins []bitSlice) (mostCommon, leastCommon bitSlice) {
	mostCommon, leastCommon = make(bitSlice, len(bins[0])), make(bitSlice, len(bins[0]))
	binSum := make(bitSlice, len(bins[0]))
	for i := range binSum {
		for _, bin := range bins {
			binSum[i] += bin[i]
		}
		if binSum[i] >= len(bins)/2 {
			mostCommon[i], leastCommon[i] = 1, 0
		} else {
			mostCommon[i], leastCommon[i] = 0, 1
		}
	}
	return
}

func mostCommonBit(bins []bitSlice, position int) int {
	sum := 0
	for _, bin := range bins {
		sum += bin[position]
	}
	if sum*2 >= len(bins) {
		return 1
	} else {
		return 0
	}
}

func leastCommonBit(bins []bitSlice, position int) int {
	return (mostCommonBit(bins, position) + 1) % 2
}

func binarySieve(bins []bitSlice, targetBit func([]bitSlice, int) int) bitSlice {
	validBins := append([]bitSlice(nil), bins...)
	for bitIndex := range bins[0] {
		var remainingValid []bitSlice
		target := targetBit(validBins, bitIndex)
		for _, bin := range validBins {
			if bin[bitIndex] == target {
				remainingValid = append(remainingValid, bin)
			}
		}
		if len(remainingValid) == 1 {
			return remainingValid[0]
		}
		validBins = append([]bitSlice(nil), remainingValid...)
	}
	return validBins[0]
}

func main() {
	lines := shared.ParseInputFile("input.txt")
	bins := parseBinaryStrings(lines)

	mostCommonBits, leastCommonBits := extractRates(bins)
	epsilon, gamma := bitSliceToInt(mostCommonBits), bitSliceToInt(leastCommonBits)
	powerConsumption := epsilon * gamma
	fmt.Printf("ɛ = %d, γ = %d → Power consumption is %d.\n",
		epsilon, gamma, powerConsumption)

	o2Bits, co2Bits := binarySieve(bins, mostCommonBit), binarySieve(bins, leastCommonBit)
	o2Rating, co2Rating := bitSliceToInt(o2Bits), bitSliceToInt(co2Bits)
	lifeSupportRating := o2Rating * co2Rating
	fmt.Printf("O2 Generation = %d, CO2 Scrubbing = %d → Life Support Rating is %d.\n",
		o2Rating, co2Rating, lifeSupportRating)
}
