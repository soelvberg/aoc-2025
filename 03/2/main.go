package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const batteryCount = 12
const inputTxt = "input.txt"

// const inputTxt = "input_test.txt"

func main() {
	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		log.Printf("Execution time: %s", elapsed)
	}()

	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bank := make([]int, 0, len(line))

		for _, ch := range line {
			digit, err := strconv.Atoi(string(ch))
			if err != nil {
				continue
			}
			bank = append(bank, digit)
		}

		joltage := findLargestJoltage(batteryCount, bank)

		sum += joltage
	}

	fmt.Println("Sum:", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findLargestJoltage(count int, digits []int) int {
	numDigits := count

	selectedDigits := []int{}

	lastIndex := -1

	var high func(n int)
	high = func(n int) {

		if n == 0 {
			return
		}

		highestDigit := -1

		for i := lastIndex + 1; i < len(digits)-(n-1); i++ {
			if digits[i] > highestDigit {
				highestDigit = digits[i]
				lastIndex = i
			}
		}
		selectedDigits = append(selectedDigits, highestDigit)

		high(n - 1)
	}

	high(numDigits)

	resultStr := ""
	for _, digit := range selectedDigits {
		resultStr += strconv.Itoa(digit)
	}

	result, err := strconv.Atoi(resultStr)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
