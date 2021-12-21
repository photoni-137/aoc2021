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

type bitSlice []int

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func bitSliceToInt(bin bitSlice) (number int64, err error) {
	var sb strings.Builder
	for _, bit := range bin {
		sb.WriteString(strconv.Itoa(bit))
	}
	number, err = strconv.ParseInt(sb.String(), 2, 64)
	return
}

func readBinaryStrings(r io.Reader) (bins []bitSlice, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		bin := make(bitSlice, len(line))
		for i, char := range line {
			bin[i], err = strconv.Atoi(string(char))
			if err != nil || !(bin[i] == 0 || bin[i] == 1) {
				return
			}
		}
		bins = append(bins, bin)
	}
	return
}

func extractRates(bins []bitSlice) (mostCommon, leastCommon bitSlice) {
	binSum := make(bitSlice, len(bins[0]))
	mostCommon = make(bitSlice, len(bins[0]))
	leastCommon = make(bitSlice, len(bins[0]))
	for i := range binSum {
		for _, bin := range bins {
			binSum[i] += bin[i]
		}
		if binSum[i] >= len(bins)/2 {
			mostCommon[i] = 1
			leastCommon[i] = 0
		} else {
			mostCommon[i] = 0
			leastCommon[i] = 1
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
	fmt.Printf("Valid Binaries: %d out of %d total, ratio: %f\n", len(validBins), len(bins), float64(len(validBins))/float64(len(bins)))
	for bitIndex := range bins[0] {
		var remainingValid []bitSlice
		target := targetBit(validBins, bitIndex)
		for _, bin := range validBins {
			if bin[bitIndex] == target {
				remainingValid = append(remainingValid, bin)
			}
		}
		fmt.Printf("Valid Binaries: %d out of %d total, ratio: %f\n", len(remainingValid), len(validBins), float64(len(remainingValid))/float64(len(validBins)))
		if len(remainingValid) == 1 {
			return remainingValid[0]
		}
		validBins = append([]bitSlice(nil), remainingValid...)
	}
	return validBins[0]
}

func main() {
	inputFilePath := "diagnostic_report.txt"
	reader, err := os.Open(inputFilePath)
	handle(err)
	defer reader.Close()

	bins, err := readBinaryStrings(reader)
	handle(err)

	mostCommonBits, leastCommonBits := extractRates(bins)
	fmt.Println(mostCommonBits)
	fmt.Println(leastCommonBits)
	epsilon, err := bitSliceToInt(mostCommonBits)
	handle(err)
	gamma, err := bitSliceToInt(leastCommonBits)
	handle(err)
	powerConsumption := epsilon * gamma
	fmt.Printf("ɛ = %d, γ = %d → Power consumption is %d.", epsilon, gamma, powerConsumption)

	o2Bits := binarySieve(bins, mostCommonBit)
	co2Bits := binarySieve(bins, leastCommonBit)
	fmt.Println(o2Bits)
	fmt.Println(co2Bits)
	o2Rating, err := bitSliceToInt(o2Bits)
	handle(err)
	co2Rating, err := bitSliceToInt(co2Bits)
	handle(err)
	lifeSupportRating := o2Rating * co2Rating
	fmt.Printf("O2 Generation = %d, CO2 Scrubbing = %d → Life Support Rating is %d.",
		o2Rating, co2Rating, lifeSupportRating)
}
