package main

import "fmt"

type Pair struct {
	a, b int
}

type Graph struct {
	adjList [][]int
	ids     []int
	low     []int
	visited []bool
}

func scanCreateGraph() *Graph {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	list := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		_, _ = fmt.Scan(&u, &v)
		list[u] = append(list[u], v)
		list[v] = append(list[v], u)
	}
	graph := new(Graph)
	graph.adjList = list
	graph.ids = make([]int, m)
	graph.low = make([]int, m)
	graph.visited = make([]bool, m)
	return graph
}

func dfs(g *Graph, currentId, at, parent int, bridges *[]Pair) {
	g.visited[at] = true
	currentId++
	g.low[at] = currentId
	g.ids[at] = currentId
	for toIndex := range g.adjList[at] {
		toValue := g.adjList[at][toIndex]
		if toValue == parent {
			continue
		}
		if !g.visited[toValue] {
			dfs(g, currentId, toValue, at, bridges)
			if g.low[toValue] < g.low[at] {
				g.low[at] = g.low[toValue]
			}
			if g.ids[at] < g.low[toValue] {
				var bridge Pair
				bridge.a = at
				bridge.b = toValue
				*bridges = append(*bridges, bridge)
			}
		} else {
			if g.ids[toValue] < g.low[at] {
				g.low[at] = g.ids[toValue]
			}
		}
	}
}

func findBridges(graph *Graph) []Pair {
	bridges := make([]Pair, 0, 1)
	for i := 0; i < len(graph.ids); i++ {
		if !graph.visited[i] {
			dfs(graph, 0, i, -1, &bridges)
		}
	}
	return bridges
}

func main() {
	graph := scanCreateGraph()
	bridges := findBridges(graph)
	fmt.Println(len(bridges))
}
