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

func main() {
	// Log execution time
	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		log.Printf("Execution time: %s\n", elapsed)
	}()

	file := readInputFile()
	defer file.Close()

	blankspaceIndexes := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for i, char := range line {
			if char == ' ' {
				// if already false, keep it false
				if val, exists := blankspaceIndexes[i]; exists {
					blankspaceIndexes[i] = val && true
				} else {
					blankspaceIndexes[i] = true
				}
			} else {
				blankspaceIndexes[i] = false
			}
		}
	}

	stringMatrix := [][]string{}
	scanner = bufio.NewScanner(file)
	file.Seek(0, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		runeLine := []rune(line)
		for i, blanks := range blankspaceIndexes {
			if blanks && i < len(runeLine) {
				runeLine[i] = ';'
			}
		}
		row := strings.FieldsFunc(string(runeLine), func(r rune) bool {
			return r == ';'
		})
		stringMatrix = append(stringMatrix, row)
	}
	fmt.Println("Processed Matrix:")
	for _, row := range stringMatrix {
		fmt.Println(row)
	}

	sumColumns := make([]int, len(stringMatrix[0]))
	lastRowIndex := len(stringMatrix) - 1

	for i := 0; i < len(stringMatrix[0]); i++ {
		sum := 0
		var newNumStr []string
		for j := 0; j < len(stringMatrix)-1; j++ {
			str := stringMatrix[j][i]
			for k, char := range str {
				if len(newNumStr) <= k {
					newNumStr = append(newNumStr, "")
				}
				newNumStr[k] += string(char)
			}

		}
		operator := strings.TrimSpace(stringMatrix[lastRowIndex][i])

		for l, s := range newNumStr {
			newNumStr[l] = strings.TrimSpace(s)
			var val int
			fmt.Sscanf(newNumStr[l], "%d", &val)
			switch operator {
			case "+":
				sum += val
			case "*":
				if sum == 0 {
					sum = 1
				}
				sum *= val
			}
			sumColumns[i] = sum
			fmt.Printf("Sum of col %d: %d\n", i, sumColumns[i])
		}
	}

	sumTotal := 0
	for _, val := range sumColumns {
		sumTotal += val
	}

	fmt.Printf("Total sum of all cols: %d\n", sumTotal)
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
