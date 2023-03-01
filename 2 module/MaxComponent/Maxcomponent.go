package main

import (
	"fmt"
	"math"
	"strconv"
)

type Graph struct {
	adjList           [][]int
	visited           []bool
	component        [][]int
	componentNodeNum []int
	componentEdgeNum []int
	//edgesCountForNode []int
}

func scanCreateGraph() (*Graph, string) {
	var n, m int
	ans := "graph G {\n"
	_, _ = fmt.Scan(&n, &m)
	howManyEdges := make([]int, n)
	list := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		_, _ = fmt.Scan(&u, &v)
		if u != v {
			list[u] = append(list[u], v)
		}
		list[v] = append(list[v], u)
		howManyEdges[v]++
		ans = ans + strconv.Itoa(u) + "--" + strconv.Itoa(v) + "!\n"
	}
	ans = ans + "}"
	graph := new(Graph)
	graph.adjList = list
	graph.visited = make([]bool, n)
	graph.component = make([][]int, 0)
	graph.componentEdgeNum = make([]int, 0)
	graph.componentNodeNum = make([]int, 0)
	//graph.edgesCountForNode = howManyEdges
	return graph, ans
}

func dfs(g *Graph, at int) {
	g.visited[at] = true
	g.component[len(g.component)-1] = append(g.component[len(g.component)-1], at)
	for nextIndex := range g.adjList[at] {
		next := g.adjList[at][nextIndex]
		if !g.visited[next] {
			dfs(g, next)
		}
	}
}

func findComponents(g *Graph) *Graph {
	count := -1

	for i := range g.adjList {
		if !g.visited[i] {
			count++
			g.component = append(g.component, nil)
			dfs(g, i)
		}
	}
	return g
}

func selectAnswer(graph *Graph, answer string) int {
	component := graph.component
	maxNumOfNodes := -1
	//maxIndex := -1
	similarNums := make([]int, 0)
	for i := range component {
		if len(component[i]) > maxNumOfNodes {
			maxNumOfNodes = len(component[i])
			//maxIndex = i
		}
	}
	for i := range component {
		if len(component[i]) == maxNumOfNodes {
			similarNums = append(similarNums, i)
		}
	}
	if len(similarNums) != 1 {
		edges := make([]int, len(similarNums))
		maxNumOfEdges := 0
		for i := range similarNums {
			for j := range component[i] {
				edges[i] += graph.edgesCountForNode[component[i][j]]
				if edges[i] > maxNumOfEdges {
					maxNumOfEdges = edges[i]
				}
			}
		}
		countMaxEdges :=
		for i := range edges {
			if edges[i] == maxNumOfEdges {
				countMaxEdges++
			}
		}
		if countMaxEdges != 1{
			minNumber := make([]int, len(edges))
			for i := range minNumber{
				minNumber[i] = math.MaxInt
			}
			for i := range edges{
				for j := range component[edges[i]]{
					if component[edges[i]][j] < minNumber[	]
				}
			}
		}
	}
}

func main() {
	graph, ans := scanCreateGraph()
	findComponents(graph)
	selectAnswer(graph, ans)
}
