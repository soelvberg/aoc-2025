package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const inputTxt = "input.txt"

type Point struct {
	x int
	y int
}

var memo = map[Point]int{}

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

	grid := [][]rune{}
	var start Point

	// Scan file
	isFirstLine := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if isFirstLine {
			isFirstLine = false

			for i, char := range line {
				if char == 'S' {
					start = Point{x: i, y: 0}
				}
			}
		}

		gridLine := []rune(line)
		grid = append(grid, gridLine)
	}

	totalPaths := dfs(grid, start)

	fmt.Println("Starting point:", start)
	// for _, line := range grid {
	// 	fmt.Println(string(line))
	// }

	fmt.Println("Total paths:", totalPaths)
}

func dfs(grid [][]rune, point Point) int {
	if v, ok := memo[point]; ok {
		return v
	}

	rows := len(grid)

	// Path end, +1
	if point.y == rows-1 {
		memo[point] = 1
		return 1
	}

	moves := nextMoves(grid, point)
	total := 0
	for _, move := range moves {
		total += dfs(grid, move)
	}

	memo[point] = total

	return total
}

func nextMoves(grid [][]rune, point Point) []Point {
	rows := len(grid)
	cols := len(grid[0])
	cell := grid[point.y][point.x]

	moves := []Point{}

	switch cell {
	case 'S', '.':
		// Down
		if point.y+1 < rows {
			moves = append(moves, Point{x: point.x, y: point.y + 1})
		}
	case '^':
		// Splitter: left and right
		if point.x-1 >= 0 {
			moves = append(moves, Point{x: point.x - 1, y: point.y})
		}
		if point.x+1 < cols {
			moves = append(moves, Point{x: point.x + 1, y: point.y})
		}
	default:
		log.Fatalf("Invalid cell (%d, %d)", point.x, point.y)
	}

	return moves
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
