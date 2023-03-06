package main

import "fmt"

type Edge struct {
	node, weight int
}

func getEdges() {
	var n, m int
	fmt.Scan(&n, &m)
	edges := make([][]Edge, n)
	for i := 0; i < m; i++ {
		var from, to, weight int
		fmt.Scan(&from, &to, &weight)
		edge1 := Edge{node: to, weight: weight}
		edge2 := Edge{node: from, weight: weight}
		edges[from] = append(edges[from], edge1)
		edges[to] = append(edges[to], edge2)
	}
}

func main() {
	getEdges()
}
