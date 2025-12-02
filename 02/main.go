package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var sum int = 0

type Range struct {
	Start int
	End   int
}

func main() {

	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		log.Printf("Execution time: %s", elapsed)
	}()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error open file: %v", err)
	}
	defer file.Close()

	ranges := parseFile(file)

	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			if checkRepeatePattern(i) {
				sum = sum + i
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}

func checkRepeatePattern(num int) bool {
	productIdStr := fmt.Sprintf("%d", num)
	productIdLength := len(productIdStr)

	for i := 1; i < productIdLength; i++ {
		if productIdLength%i != 0 {
			continue
		}

		partCount := productIdLength / i

		parts := make([]string, partCount)
		for j := 0; j < partCount; j++ {
			parts[j] = productIdStr[j*i : (j+1)*i]
		}
		allEqual := true
		for k := 1; k < len(parts); k++ {
			if parts[k] != parts[0] {
				allEqual = false
				break
			}
		}
		if allEqual {
			return true
		}

	}
	return false
}

func parseFile(file *os.File) []Range {
	var ranges []Range
	var content []byte = make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil {
		log.Fatalf("Error read file: %v", err)
	}
	data := string(content[:n])
	parts := strings.SplitSeq(data, ",")
	for part := range parts {
		var r Range
		_, err := fmt.Sscanf(part, "%d-%d", &r.Start, &r.End)
		if err != nil {
			log.Fatalf("Error parsing: %v", err)
		}
		ranges = append(ranges, r)
	}
	return ranges
}
