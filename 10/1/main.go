package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const inputTxt = "input_test.txt"

var sum int = 0

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
		// fmt.Println(line)

		target := ""
		buttons := []string{}
		joltages := []int{}

		indexOfLastSquareBracket := strings.Index(line, "]")
		indexOfFirstCurlyBracket := strings.Index(line, "{")

		// Target
		targetPart := line[1:indexOfLastSquareBracket]
		for _, c := range targetPart {
			switch c {
			case '.':
				target += "0"
			case '#':
				target += "1"
			default:
				log.Fatalf("Unexpected character in target: %c", c)
			}
		}

		// Buttons
		buttonsPart := line[indexOfLastSquareBracket+2 : indexOfFirstCurlyBracket-1]
		buttonsStrings := strings.Split(buttonsPart, " ")
		for _, s := range buttonsStrings {
			sHelper := ""
			for i := 0; i < len(target); i++ {
				if strings.Contains(s, fmt.Sprintf("%d", i)) {
					sHelper += "1"
				} else {
					sHelper += "0"
				}
			}
			buttons = append(buttons, sHelper)
		}
		// _ = buttons

		// Joltages
		joltagesPart := line[indexOfFirstCurlyBracket+1 : len(line)-1]
		joltagesStrings := strings.Split(joltagesPart, ",")
		for _, j := range joltagesStrings {
			var joltage int
			fmt.Sscanf(j, "%d", &joltage)
			joltages = append(joltages, joltage)
		}
		_ = joltages

		// fmt.Println("Target:", target)
		// fmt.Println("Buttons:", buttons)
		// fmt.Println("Joltages:", joltages)

		// begin
		state := strings.Repeat("0", len(target))
		count := pushButtons(state, target, buttons)
		sum += count
		// fmt.Println("Minimum button pushes to reach target:", count)
	}

	fmt.Println("Sum of minimum button pushes for all lines:", sum)
}

func pushButtons(state string, target string, buttons []string) int {
	if state == target {
		return 0
	}

	maxPushes := len(buttons)

	// 1 button, 2 buttons etc.
	for numPushes := 1; numPushes <= maxPushes; numPushes++ {
		// Generate all combinations of numPushes buttons
		result := tryCombos(state, target, buttons, numPushes, 0, []int{0})
		if result >= 0 {
			return result
		}
	}

	return -1 // target is not matched
}

func tryCombos(state string, target string, buttons []string, numPushes int, startIndex int, combination []int) int {
	// base case, combination ready
	if len(combination) == numPushes {
		testState := testCombo(state, buttons, combination)
		if testState == target {
			// found combo
			return len(combination)
		}
		return -1
	}

	// recursion, add more buttons
	for i := startIndex; i < len(buttons); i++ {
		result := tryCombos(state, target, buttons, numPushes, i+1, append(combination, i))
		if result >= 0 {
			return result
		}
	}

	return -1
}

func testCombo(state string, buttons []string, combination []int) string {
	toggleCount := make([]int, len(state))

	// apply button/add toggle counts
	for _, buttonIndex := range combination {
		button := buttons[buttonIndex]
		for j := 0; j < len(button); j++ {
			if button[j] == '1' {
				toggleCount[j]++
			}
		}
	}

	// mod 2 for final state
	result := ""
	for i := 0; i < len(state); i++ {
		currentBit := int(state[i] - '0')
		newBit := (currentBit + toggleCount[i]) % 2
		result += fmt.Sprintf("%d", newBit)
	}

	return result
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
