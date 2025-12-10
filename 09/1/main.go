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
	X int
	Y int
}

var points = []Point{}

var maxArea = 0

func findMaxArea(points []Point, point Point) {
	for _, p := range points {
		tempWidth := 0
		tempHeight := 0
		// assume a rectangle forming between two points (including the point itself)
		if p.X > point.X {
			tempWidth = (p.X - point.X) + 1
		} else if p.X < point.X {
			tempWidth = (point.X - p.X) + 1
		} else {
			tempWidth = 1
		}

		if p.Y > point.Y {
			tempHeight = (p.Y - point.Y) + 1
		} else if p.Y < point.Y {
			tempHeight = (point.Y - p.Y) + 1
		} else {
			tempHeight = 1
		}

		area := tempWidth * tempHeight
		if area > maxArea {
			maxArea = area
		}
	}
}

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
		var point Point
		_, err := fmt.Sscanf(line, "%d,%d", &point.X, &point.Y)
		if err != nil {
			log.Fatalf("Error parsing line: %v", err)
		}
		points = append(points, point)
	}

	fmt.Println("Points read from file:")
	for _, p := range points {
		fmt.Printf("(%d, %d)\n", p.X, p.Y)
	}

	for _, point := range points {
		findMaxArea(points, point)
	}
	fmt.Printf("Max area: %d\n", maxArea)
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
