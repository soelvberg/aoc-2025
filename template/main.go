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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
