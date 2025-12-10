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
	X, Y int
}

type Line struct {
	A, B Point
}

type Polygon struct {
	Points []Point
	Lines  []Line
}

var maxArea = 0
var verticals []Line
var horizontals []Line

func createPolygon(points []Point) Polygon {
	n := len(points)
	lines := make([]Line, 0, n)

	for i := range n {
		a := points[i]
		b := points[(i+1)%n] // wrap around to the first point

		lines = append(lines, Line{A: a, B: b})
	}

	return Polygon{
		Points: points,
		Lines:  lines,
	}
}

func pip(point Point) bool {
	inside := false

	// check point on edge - Horizontal lines
	for _, line := range horizontals {
		if point.Y == line.A.Y &&
			point.X >= min(line.A.X, line.B.X) && point.X <= max(line.A.X, line.B.X) {
			return true
		}
	}

	// check point on edge - Vertical lines
	for _, line := range verticals {
		if point.X == line.A.X &&
			point.Y >= min(line.A.Y, line.B.Y) && point.Y <= max(line.A.Y, line.B.Y) {
			return true
		}
	}

	// ray-cast intersect vertical lines (ray is horizontal)
	for _, line := range verticals {
		if point.X < line.A.X && point.Y >= min(line.A.Y, line.B.Y) && point.Y < max(line.A.Y, line.B.Y) {
			inside = !inside
		}
	}

	return inside
}

func process(polygon Polygon) {
	for _, point := range polygon.Points {
		for _, other := range polygon.Points {
			if point == other {
				continue
			}
			// rectangle bounds
			xmin := min(point.X, other.X)
			xmax := max(point.X, other.X)
			ymin := min(point.Y, other.Y)
			ymax := max(point.Y, other.Y)

			// rectangle corners
			c1 := Point{xmin, ymin}
			c2 := Point{xmin, ymax}
			c3 := Point{xmax, ymin}
			c4 := Point{xmax, ymax}

			// check corners inside
			if !pip(c1) ||
				!pip(c2) ||
				!pip(c3) ||
				!pip(c4) {
				continue
			}

			// check if any other points are inside the rectangle
			invalid := false
			// for _, p := range polygon.Points {
			// 	if p.X > xmin && p.X < xmax && p.Y > ymin && p.Y < ymax {
			// 		invalid = true
			// 		break
			// 	}
			// }
			// if invalid {
			// 	continue
			// }

			// line-rectangle intersection check
			for _, line := range horizontals {
				y := line.A.Y
				if y > ymin && y < ymax {
					lineXMin := min(line.A.X, line.B.X)
					lineXMax := max(line.A.X, line.B.X)
					if lineXMin < xmax && lineXMax > xmin {
						invalid = true
						break
					}
				}
			}
			if !invalid {
				for _, line := range verticals {
					x := line.A.X
					if x > xmin && x < xmax {
						lineYMin := min(line.A.Y, line.B.Y)
						lineYMax := max(line.A.Y, line.B.Y)
						if lineYMin < ymax && lineYMax > ymin {
							invalid = true
							break
						}
					}
				}
			}
			if invalid {
				continue
			}

			// valid rectangle, compute area
			area := (xmax - xmin + 1) * (ymax - ymin + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
}

// 4583207265 -- too high
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

	var points = []Point{}

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

	polygon := createPolygon(points)

	for _, line := range polygon.Lines {
		if line.A.X == line.B.X {
			verticals = append(verticals, line)
		}
	}

	for _, line := range polygon.Lines {
		if line.A.Y == line.B.Y {
			horizontals = append(horizontals, line)
		}
	}

	process(polygon)

	fmt.Printf("Max area: %d\n", maxArea)

	// fmt.Printf("Constructed polygon: %+v\n", polygon)
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
