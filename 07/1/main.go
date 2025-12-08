package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const inputTxt = "input.txt"

func main() {
	// Log execution time
	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		log.Printf("Execution time: %s\n", elapsed)
	}()

	// Read file
	file := readInputFile()
	defer file.Close()

	// Scan file
	isFirstLine := true
	previousLineBeamIndexes := map[int]bool{}
	countSplits := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if isFirstLine {
			isFirstLine = false

			for i, char := range line {
				if char == 'S' {
					previousLineBeamIndexes[i] = true
				}
			}
		} else {
			currentLineBeamIndexes := map[int]bool{}
			for prevBeamIndex := range previousLineBeamIndexes {

				for j, char := range line {
					if j == prevBeamIndex && char == '^' {
						countSplits++
						currentLineBeamIndexes[j-1] = true
						currentLineBeamIndexes[j+1] = true
					} else if j == prevBeamIndex && char == '.' {
						currentLineBeamIndexes[j] = true
					}
				}
			}
			previousLineBeamIndexes = currentLineBeamIndexes
		}
	}
	fmt.Printf("Number of splits: %d\n", countSplits)
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
