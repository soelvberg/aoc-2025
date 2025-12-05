package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	readRangeMode := true

	var freshRanges [][]int
	// fresh := make(map[int]bool)
	// var available []int

	freshCount := 0

	// Scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readRangeMode = false
			continue
		}

		if readRangeMode {
			parts := strings.Split(line, "-")

			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			freshRanges = append(freshRanges, []int{start, end})

			// for i := start; i <= end; i++ {
			// 	fresh[i] = true
			// }

		} else {
			id, _ := strconv.Atoi(line)
			// available = append(available, id)

			for _, r := range freshRanges {
				if id >= r[0] && id <= r[1] {
					// fmt.Printf("%d in range\n", id)
					freshCount++
					break
				}
			}
		}

	}
	// fmt.Println(freshRanges)
	// fmt.Println(available)
	fmt.Printf("Fresh count: %d\n", freshCount)
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
