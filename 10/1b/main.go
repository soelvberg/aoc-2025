package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const inputTxt = "input.txt"

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

		// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
		// [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
		// [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}

		// target = strings.Trim(parts[0], "[]")

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
			// s := strings.Trim(sStr, "()")
			// sParts := strings.Split(s, ",")
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

		// Begin
		state := strings.Repeat("0", len(target))
		count := pushButtons(state, target, buttons, 1)
		sum += count
		// fmt.Println("Minimum button pushes to reach target:", count)
	}
	fmt.Println("Sum of minimum button pushes for all lines:", sum)
}

func pushButtons(state string, target string, buttons []string, pushCount int) int {
	if state == target {
		return 0
	}

	visited := make(map[string]bool)
	queue := []struct {
		state string
		count int
	}{{state, 0}}
	visited[state] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.state == target {
			// fmt.Println("Reached target with push count:", current.count)
			return current.count
		}

		// Try each button
		for i := 0; i < len(buttons); i++ {
			newState := ""
			for j := 0; j < len(current.state); j++ {
				if buttons[i][j] == '1' {
					// Toggle bit
					if current.state[j] == '0' {
						newState += "1"
					} else {
						newState += "0"
					}
				} else {
					newState += string(current.state[j])
				}
			}

			if !visited[newState] {
				visited[newState] = true
				queue = append(queue, struct {
					state string
					count int
				}{newState, current.count + 1})
			}
		}
	}

	return -1 // Target not reachable
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
