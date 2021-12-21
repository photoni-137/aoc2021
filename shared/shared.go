package shared

import (
	"bufio"
	"log"
	"os"
)

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func carefulClose(file *os.File) {
	err := file.Close()
	Handle(err)
}

func ParseInputFile(filePath string) (lines []string) {
	reader, err := os.Open(filePath)
	if err != nil {
		Handle(err)
	}
	defer carefulClose(reader)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
