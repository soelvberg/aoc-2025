package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/kyroy/kdtree"
)

const inputTxt = "input_test.txt"

type Vec3 []float64

func (v Vec3) Dimensions() int {
	return len(v)
}

func (v Vec3) Dimension(i int) float64 {
	return v[i]
}

// group Vec3 points, I need to be able to add points to separate groups, so maybe int array containing arrays of Vec3
var pointGroups map[int][]Vec3 = make(map[int][]Vec3)
var tree = kdtree.New([]kdtree.Point{})

// var nearestDistances map[float64]Vec3 = make(map[float64]Vec3)

// slice with floats and points
var nearestDistances []struct {
	distance float64
	point    Vec3
}

func main() {
	// instantiate empty k-d

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

		// Parse line into Vec3
		var x, y, z float64
		_, err := fmt.Sscanf(line, "%f,%f,%f", &x, &y, &z)
		if err != nil {
			log.Fatalf("Error parsing point: %v", err)
		}
		point := Vec3{x, y, z}
		tree.Insert(point)
	}

	// iterate over points in tree and find nearest neighbor distance
	for _, point := range tree.Points() {
		findDistanceNearest(tree, point.(Vec3), 1)
	}

	// sort nearestDistances by distance ascending
	sort.Slice(nearestDistances, func(i, j int) bool {
		return nearestDistances[i].distance < nearestDistances[j].distance
	})

	// print nearestDistances
	fmt.Println("Nearest Distances:")
	for _, nd := range nearestDistances {
		fmt.Printf("Point: %v, Nearest Distance: %f\n", nd.point, nd.distance)
		groupPoints(tree, nd.point)
	}

	// firstPoint := tree.Points()[0].(Vec3)
	// groupPoints(tree, firstPoint)

	fmt.Println("Point Groups:", pointGroups)

	// count number of points in each group and store the count for the three largest groups
	groupSizes := make([]int, 0, len(pointGroups))
	for _, points := range pointGroups {
		groupSizes = append(groupSizes, len(points))
		fmt.Printf("Group size: %d\n", len(points))
	}

	// order groupSizes descending
	for i := 0; i < len(groupSizes)-1; i++ {
		for j := i + 1; j < len(groupSizes); j++ {
			if groupSizes[i] < groupSizes[j] {
				groupSizes[i], groupSizes[j] = groupSizes[j], groupSizes[i]
			}
		}
	}
	fmt.Println("Group Sizes Descending:", groupSizes)

	// multiply sizes of three largest groups
	product := groupSizes[0] * groupSizes[1] * groupSizes[2]
	fmt.Printf("Product of sizes of three largest groups: %d\n", product)

	fmt.Printf("Number of groups: %d\n", len(pointGroups))

}

func findDistanceNearest(tree *kdtree.KDTree, point Vec3, n int) {
	nearest := tree.KNN(point, n+1)
	nearestPoint := nearest[n].(Vec3)

	// calculate manhattan distance
	distance := 0.0
	for i := 0; i < 3; i++ {
		diff := point[i] - nearestPoint[i]
		if diff < 0 {
			diff = -diff
		}
		distance += diff
	}

	nearestDistances = append(nearestDistances, struct {
		distance float64
		point    Vec3
	}{
		distance: distance,
		point:    point,
	})
}

func groupPoints(tree *kdtree.KDTree, point Vec3) {
	// for _, _point := range tree.Points() {
	// point := _point.(Vec3)
	// find 1 nearest neighbor
	nearest := tree.KNN(point, 2)
	// nearest := twoNearest[1].(Vec3)
	nearestPoint := nearest[1].(Vec3)

	if len(pointGroups) == 0 {
		pointGroups[1] = []Vec3{point}
		pointGroups[1] = append(pointGroups[1], nearestPoint)
		return
	}

	// iterate over pointGroups to see if nearest neighbor is in any group
	for groupID, points := range pointGroups {
		for _, p := range points {
			// if point is already in group, continue
			if p[0] == point[0] && p[1] == point[1] && p[2] == point[2] {
				return
			}

			if p[0] == nearestPoint[0] && p[1] == nearestPoint[1] && p[2] == nearestPoint[2] {
				// add point to this group
				pointGroups[groupID] = append(pointGroups[groupID], point)
				return
			}
		}
	}
	newGroupID := len(pointGroups) + 1
	pointGroups[newGroupID] = []Vec3{point}
	// groupPoints(tree, nearestPoint)
	// }

}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
