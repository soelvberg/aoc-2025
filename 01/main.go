package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Line  string
	Value int
}

const start = 50
const dialCount = 100

var currentPosition = start
var pointAt0Count = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var records []Record
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		record := parseLine(line)
		records = append(records, record)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	for _, record := range records {

		move(record.Value)

		// if currentPosition == 0 {
		// 	pointAt0Count++
		// }
	}

	log.Printf("Times Pointed at 0: %d", pointAt0Count)
}

func parseLine(line string) Record {
	direction := line[0]
	number, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatalf("Invalid number: %s", line)
	}

	value := number

	if direction == 'L' {
		value = -value
	}

	return Record{
		Line:  line,
		Value: value,
	}
}

func add(a, b int) int {
	return (a + b) % dialCount
}

func sub(a, b int) int {
	return (a - b + dialCount) % dialCount
}

func move(value int) {
	if value < 0 {
		// backward
		for i := 1; i <= -value; i++ {
			currentPosition = sub(currentPosition, 1)
			if currentPosition == 0 {
				pointAt0Count++
			}
		}
	} else {
		// forward
		for i := 1; i <= value; i++ {
			currentPosition = add(currentPosition, 1)
			if currentPosition == 0 {
				pointAt0Count++
			}
		}
	}

	//fmt.Printf("Moved from %d to %d by %d\n", position, currentPosition, value)

	//currentPosition = tempPosition
}
