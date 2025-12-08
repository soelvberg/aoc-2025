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

	// Read file
	file := readInputFile()
	defer file.Close()

	var matrix [][]string

	// Scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		var row []string
		row = append(row, fields...)
		matrix = append(matrix, row)
	}

	sumColumns := make([]int, len(matrix[0]))
	lastRowIndex := len(matrix) - 1

	for i := 0; i < len(matrix[0]); i++ {
		sum := 0
		for j := 0; j < len(matrix)-1; j++ {
			var val int
			fmt.Sscanf(matrix[j][i], "%d", &val)
			operator := matrix[lastRowIndex][i]
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
		}
		fmt.Printf("Sum of column %d: %d\n", i, sumColumns[i])
	}

	sumTotal := 0
	for _, val := range sumColumns {
		sumTotal += val
	}

	fmt.Printf("Total sum of all columns: %d\n", sumTotal)

	// fmt.Println("Matrix:")
	// for _, row := range matrix {
	// 	fmt.Println(row)
	// }
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
