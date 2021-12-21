package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readNumbers(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	var numbers []int
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}
	return numbers, scanner.Err()
}

func slidingSums(numbers []int, windowSize int) []int {
	var sums []int
	for i := 0; i <= len(numbers)-windowSize; i++ {
		sum := 0
		for _, number := range numbers[i : i+windowSize] {
			sum += number
		}
		sums = append(sums, sum)
	}
	return sums
}

func countIncreases(numbers []int) int {
	counter := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			counter++
		}
	}
	return counter
}

func main() {
	inputFilePath := "sweep_report.txt"
	reader, err := os.Open(inputFilePath)
	defer reader.Close()
	handle(err)
	depths, err := readNumbers(reader)
	handle(err)
	fmt.Printf("Depth increased %d times in %d measurements.\n", countIncreases(depths), len(depths))

	depthWindows := slidingSums(depths, 3)
	fmt.Printf("Depth windows increased %d times in %d aggregated measurements.\n", countIncreases(depthWindows), len(depthWindows))
}
