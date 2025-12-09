package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

const inputTxt = "input.txt"
const maxConnections = 1000

type Point struct {
	X, Y, Z float64
}

type Edge struct {
	Dist float64
	I, J int // point indices
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // already in same set
	}

	// union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

func KShortestConnections(points []Point, K int) []Edge {
	var edges []Edge
	n := len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[i].X - points[j].X
			dy := points[i].Y - points[j].Y
			dz := points[i].Z - points[j].Z
			dist := math.Sqrt(dx*dx + dy*dy + dz*dz)
			edges = append(edges, Edge{Dist: dist, I: i, J: j})
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].Dist < edges[b].Dist
	})

	if K > len(edges) {
		K = len(edges)
	}

	return edges[:K]
}

func main() {
	points := []Point{}

	startTime := time.Now()
	defer func() {
		elapsed := time.Since(startTime)
		log.Printf("Execution time: %s\n", elapsed)
	}()

	file := readInputFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var p Point
		_, err := fmt.Sscanf(line, "%f,%f,%f", &p.X, &p.Y, &p.Z)
		if err != nil {
			log.Fatalf("Error parsing line '%s': %v", line, err)
		}
		points = append(points, p)
	}

	K := 999999999999999999
	result := KShortestConnections(points, K)

	uf := NewUnionFind(len(points))
	connectionsCount := 0

	for i := 0; i < maxConnections && i < len(result); i++ {
		e := result[i]
		if uf.Union(e.I, e.J) {
			connectionsCount++
		}
	}

	fmt.Println("Group count:", len(points)-connectionsCount) // ???
	fmt.Println("Edges processed:", maxConnections)
	fmt.Println("Actual connections made:", connectionsCount)

	// Count component sizes
	componentSize := make(map[int]int)
	for i := 0; i < len(points); i++ {
		root := uf.Find(i)
		componentSize[root]++
	}

	groupSizes := make([]int, 0, len(componentSize))
	for _, size := range componentSize {
		groupSizes = append(groupSizes, size)
	}

	sort.Slice(groupSizes, func(i, j int) bool {
		return groupSizes[i] > groupSizes[j]
	})

	fmt.Println("Three largest group sizes:", groupSizes[:3])

	product := groupSizes[0] * groupSizes[1] * groupSizes[2]
	fmt.Println("Product of three largest group sizes:", product)
}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
