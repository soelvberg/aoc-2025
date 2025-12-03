package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		log.Printf("Execution time: %s", elapsed)
	}()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int = 0
	// var lines [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]int, 0, len(line))

		for _, ch := range line {
			digit, err := strconv.Atoi(string(ch))
			if err != nil {
				continue
			}
			digits = append(digits, digit)
		}

		// lines = append(lines, digits)

		joltage := findLargestJoltage(digits)
		fmt.Println("Largest joltage:", joltage)

		sum += joltage
	}

	fmt.Println("Sum:", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findLargestJoltage(digits []int) int {
	first := 0
	second := 0
	indexFirst := -1

	for i, num := range digits[:len(digits)-1] {
		if num > first {
			first = num
			indexFirst = i
		}
	}

	for i := indexFirst + 1; i < len(digits); i++ {
		if digits[i] > second {
			second = digits[i]
		}
	}

	resultStr := fmt.Sprintf("%d%d", first, second)
	result, err := strconv.Atoi(resultStr)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
