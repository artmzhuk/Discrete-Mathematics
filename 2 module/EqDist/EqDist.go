package main

import (
	"fmt"
	"sort"
)

type Queue struct {
	queue []int
}

type Graph struct {
	edges [][]int
	bases []int //опорные вершины
}

func (a *Queue) IsEmpty() bool {
	if len(a.queue) == 0 {
		return true
	} else {
		return false
	}
}

func (a *Queue) Push(x int) {
	a.queue = append(a.queue, x)
}

func (a *Queue) Pop() int {
	res := a.queue[0]
	a.queue = a.queue[1:]
	return res
}

func getInput() *Graph {
	var N, M, K int
	_, _ = fmt.Scan(&N, &M)
	edges := make([][]int, N)
	for i := 0; i < M; i++ {
		var from, to int
		_, _ = fmt.Scan(&from, &to)
		edges[from] = append(edges[from], to)
		edges[to] = append(edges[to], from)
	}
	_, _ = fmt.Scan(&K)
	bases := make([]int, K)
	for i := 0; i < K; i++ {
		var base int
		_, _ = fmt.Scan(&base)
		bases[i] = base
	}
	sort.Slice(bases, func(i, j int) bool { return bases[i] < bases[j] })
	g := new(Graph)
	g.bases = bases
	g.edges = edges
	return g
}

func bfs(graph *Graph) []int {
	previousDist := make([]int, len(graph.edges))
	isSameDist := make([]bool, len(graph.edges))
	for currentBaseId := 0; currentBaseId < len(graph.bases); currentBaseId++ {
		visited := make([]bool, len(graph.edges))
		distToCurrentBase := make([]int, len(graph.edges))
		queue := new(Queue)
		visited[graph.bases[currentBaseId]] = true
		queue.Push(graph.bases[currentBaseId])
		for !queue.IsEmpty() {
			popped := queue.Pop()
			visited[popped] = true
			for j := 0; j < len(graph.edges[popped]); j++ {
				neighbour := graph.edges[popped][j]
				if !visited[neighbour] {
					queue.Push(neighbour)
					visited[neighbour] = true
					distToCurrentBase[neighbour] = distToCurrentBase[popped] + 1
					if previousDist[neighbour] == 0 {
						previousDist[neighbour] = distToCurrentBase[popped] + 1
						isSameDist[neighbour] = true
					} else if isSameDist[neighbour] && previousDist[neighbour] != distToCurrentBase[popped]+1 {
						isSameDist[neighbour] = false
					}
				}
			}
		}
	}
	result := make([]int, 0)
	for i, j := 0, 0; i < len(isSameDist); i++ {
		if isSameDist[i] == true && graph.bases[j] != i {
			result = append(result, i)
		}
		if graph.bases[j] == i {
			j++
		}
	}
	return result
}

func main() {
	graph := getInput()
	result := bfs(graph)
	for i := range result {
		fmt.Println(result[i])
	}
}
