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

// const maxConnections = 1000

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

	groupCount := 0
	for i := 0; i < len(result); i++ {
		e := result[i]
		if uf.Union(e.I, e.J) {
			connectionsCount++
		}

		groupCount = len(points) - connectionsCount
		if groupCount == 1 {
			pointI_Xvalue := points[e.I].X
			pointJ_Xvalue := points[e.J].X
			fmt.Println("Final connection made between points with X values:", pointI_Xvalue, "and", pointJ_Xvalue)

			product := pointI_Xvalue * pointJ_Xvalue
			fmt.Println("Product of X values:", product)

			break
		}
	}

	fmt.Println("Group count:", len(points)-connectionsCount) // ???
	// fmt.Println("Edges processed:", maxConnections)
	fmt.Println("Actual connections made:", connectionsCount)

}

func readInputFile() *os.File {
	file, err := os.Open(inputTxt)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}
