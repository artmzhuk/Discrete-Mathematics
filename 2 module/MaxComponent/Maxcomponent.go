package main

import (
	"fmt"
)

type Graph struct {
	adjList          [][]int
	visited          []bool
	component        [][]int
	numOfEdgesToNode []int
}

func scanCreateGraph() (*Graph, []int) {
	var n, m int
	ans := make([]int, 0)
	_, _ = fmt.Scan(&n, &m)
	numOfEdgesToNode := make([]int, n)
	list := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		_, _ = fmt.Scan(&u, &v)
		if u != v {
			list[u] = append(list[u], v)
		}
		list[v] = append(list[v], u)
		numOfEdgesToNode[v]++
		ans = append(ans, u, v)
	}
	graph := new(Graph)
	graph.adjList = list
	graph.visited = make([]bool, n)
	graph.component = make([][]int, 0)
	graph.numOfEdgesToNode = numOfEdgesToNode
	return graph, ans
}

func dfs(g *Graph, at int) {
	g.visited[at] = true
	g.component[len(g.component)-1] = append(g.component[len(g.component)-1], at)
	for i := range g.adjList[at] {
		next := g.adjList[at][i]
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

func printAnswer(g *Graph, input []int) {
	var idToPrint int
	nextStep := true
	if nextStep {
		maxNumOfNodes := -1
		similarNums := 0
		for i := range g.component {
			if len(g.component[i]) > maxNumOfNodes {
				maxNumOfNodes = len(g.component[i])
				similarNums = 0
				idToPrint = i
			}
			if len(g.component[i]) == maxNumOfNodes {
				similarNums++
			}
		}
		if similarNums == 1 {
			nextStep = false
		}
	}
	if nextStep {
		edgeNum := make([]int, len(g.component))
		maxNumOfEdges := -1
		similarEdges := 0
		for i := range g.component {
			for j := range g.component[i] {
				edgeNum[i] += g.numOfEdgesToNode[g.component[i][j]]
				if edgeNum[i] > maxNumOfEdges {
					maxNumOfEdges = edgeNum[i]
					similarEdges = 0
					idToPrint = i
				}
				if edgeNum[i] == maxNumOfEdges {
					similarEdges++
				}
			}
		}
		if similarEdges == 1 {
			nextStep = false
		}
	}
	if nextStep {
		minIds := make([]int, len(g.component))
		for i := range g.component {
			minId := g.component[i][0]
			for j := range g.component[i] {
				if g.component[i][j] < minId {
					minId = g.component[i][j]
				}
			}
			minIds[i] = minId
		}
		idToPrint = 0
		for i := range minIds {
			if minIds[i] < minIds[idToPrint] {
				idToPrint = i
			}
		}
	}
	redNodes := make(map[int]bool)
	fmt.Print("Graph g {\n")
	for i := range g.component[idToPrint] {
		redNodes[g.component[idToPrint][i]] = true
		fmt.Print(g.component[idToPrint][i], " [color = red]\n")
	}

	for i := range input {
		if i%2 == 0 {
			fmt.Print(input[i], "--")
		} else {
			fmt.Print(input[i])
			if redNodes[input[i]] {
				fmt.Print(" [color = red]")
			}
			fmt.Print("\n")
		}
	}
	fmt.Print("}")
}

func main() {
	graph, input := scanCreateGraph()
	findComponents(graph)
	printAnswer(graph, input)
}
