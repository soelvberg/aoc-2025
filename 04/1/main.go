package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const inputTxt = "input.txt"

func main() {
	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		fmt.Printf("Execution time: %s\n", elapsed)
	}()

	file, err := os.Open(inputTxt)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	matrix := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for _, ch := range line {
			row = append(row, ch)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	traverseMatrix(matrix)
}

func traverseMatrix(matrix [][]rune) {
	countAccessibleRolls := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			_ = matrix[y][x]
			if isRoll(x, y, matrix) {
				numRolls := checkAdjacentNumRolls(x, y, matrix)
				// fmt.Printf("(%d, %d): %c - %d adjacent @\n", x, y, matrix[y][x], numRolls)
				if numRolls < 4 {
					countAccessibleRolls++
				}
				continue
			} else {
				// fmt.Printf("(%d, %d): %c - N/A \n", x, y, matrix[y][x])
			}
		}
	}
	fmt.Printf("Count accessible rolls: %d\n", countAccessibleRolls)
}

func checkAdjacentNumRolls(x, y int, matrix [][]rune) int {
	numRolls := 0
	directions := [8][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		if newY >= 0 && newY < len(matrix) && newX >= 0 && newX < len(matrix[newY]) {
			if matrix[newY][newX] == '@' {
				numRolls++
			}
		}
	}
	return numRolls
}

func isRoll(x, y int, matrix [][]rune) bool {
	return matrix[y][x] == '@'
}
