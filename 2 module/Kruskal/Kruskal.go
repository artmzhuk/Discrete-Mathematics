package main

import (
	"fmt"
	"math"
	"sort"
)

type edge struct {
	from, to int
	weight   float64
}

type rollercoaster struct {
	x, y float64
}

func (a rollercoaster) getR(b rollercoaster) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func getEdges() []edge {
	n := 0
	_, _ = fmt.Scan(&n)
	rollers := make([]rollercoaster, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&rollers[i].x, &rollers[i].y)
	}
	edges := make([]edge, 0, (n*(n-1))/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			toAppend := edge{
				from:   i,
				to:     j,
				weight: rollers[i].getR(rollers[j]),
			}
			edges = append(edges, toAppend)
		}
	}
	return edges
}

func findParent(parent []int, component int) int {
	if parent[component] == component {
		return component
	} else {
		parent[component] = findParent(parent, parent[component])
		return parent[component]
	}
}

func unionSet(u, v int, parent, rank []int) {
	u = findParent(parent, u)
	v = findParent(parent, v)
	if rank[u] < rank[v] {
		parent[u] = v
	} else if rank[u] < rank[v] {
		parent[v] = u
	} else {
		parent[v] = u
		rank[u]++
	}
}

func kruskal(edges []edge) float64 {
	sort.Slice(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight })
	var minCost float64
	parent := make([]int, len(edges))
	rank := make([]int, len(edges))
	for i := range parent {
		parent[i] = i
	}
	for i := 0; i < len(edges); i++ {
		v1 := findParent(parent, edges[i].from)
		v2 := findParent(parent, edges[i].to)
		if v1 != v2 {
			unionSet(v1, v2, parent, rank)
			minCost += edges[i].weight
			//fmt.Printf("from %d to %d is %f\n", edges[i].from, edges[i].to, edges[i].weight)
		}
	}
	return minCost
}

func main() {
	edges := getEdges()
	fmt.Printf("%.2f\n", kruskal(edges))
}
