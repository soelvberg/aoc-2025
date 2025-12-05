package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var freshRanges [][]int
	var cleanRanges [][]int

	// freshCount := 0

	// Scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Split(line, "-")

		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		freshRanges = append(freshRanges, []int{start, end})
	}

	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})

	// fmt.Println(freshRanges)

	for _, fr := range freshRanges {
		if len(cleanRanges) == 0 {
			cleanRanges = append(cleanRanges, []int{fr[0], fr[1]})
			continue
		}

		freshStart := fr[0]
		freshEnd := fr[1]

		lastCleanRange := cleanRanges[len(cleanRanges)-1]
		// cleanStart := lastCleanRange[0]
		cleanEnd := lastCleanRange[1]

		if freshStart <= cleanEnd+1 {
			if freshEnd > cleanEnd {
				lastCleanRange[1] = freshEnd
			}
		} else {
			cleanRanges = append(cleanRanges, []int{freshStart, freshEnd})
		}
	}

	totalFresh := 0
	for _, cr := range cleanRanges {
		start := cr[0]
		end := cr[1]
		totalFresh += (end - start + 1)
	}
	fmt.Println("Total fresh IDs:", totalFresh)

	// 342018167474526
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
