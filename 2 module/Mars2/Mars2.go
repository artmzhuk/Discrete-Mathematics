package main

import (
	"fmt"
)

type Graph struct {
	adjList [][]int
}

type solutions struct {
	colors  [][]int
	balance []int
}

type Queue struct {
	array []int
}

func (q *Queue) Empty() bool {
	if len(q.array) == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) Pop() int {
	res := q.array[0]
	q.array = q.array[1:]
	return res
}

func (q *Queue) Push(x int) {
	q.array = append(q.array, x)
}

func inputGraph() *Graph {
	var n int
	_, _ = fmt.Scan(&n)
	list := make([][]int, n)
	for i := 0; i < n; i++ {
		list[i] = make([]int, 0)
		for j := 0; j < n; j++ {
			var sign string
			_, _ = fmt.Scan(&sign)
			if ([]rune(sign))[0] == '+' {
				list[i] = append(list[i], j)
			}
		}
	}
	g := new(Graph)
	g.adjList = list
	return g
}

func findSolutions(g *Graph) solutions {
	visited := make([]bool, len(g.adjList))
	possibleSolutions := make([][]int, 0)
	for i := 0; i < len(g.adjList); i++ {
		if visited[i] {
			continue
		}
		isBipartite := true
		color := make([]int, len(g.adjList))
		queue := make([]int, 0)
		color[i] = 1
		visited[i] = true
		queue = append(queue, i)
		for len(queue) != 0 && isBipartite {
			popped := queue[0]
			queue = queue[1:]
			for j := range g.adjList[popped] {
				if color[g.adjList[popped][j]] == 0 {
					if color[popped] == 1 {
						color[g.adjList[popped][j]] = 2
					} else {
						color[g.adjList[popped][j]] = 1
					}
					visited[g.adjList[popped][j]] = true
					queue = append(queue, g.adjList[popped][j])
				} else {
					if color[g.adjList[popped][j]] == color[popped] {
						isBipartite = false
						break
					}
				}
			}
		}
		if isBipartite {
			possibleSolutions = append(possibleSolutions, color)
		}
	}
	if len(possibleSolutions) == 0 {
		fmt.Println("No solution")
		return solutions{
			colors: nil,
		}
	} else {
		return solutions{
			colors: possibleSolutions,
		}
	}
}

func main() {
	g := inputGraph()
	sol := findSolutions(g)
	fmt.Println(sol)
}

/*func biPart(g *Graph){
	color := make([]int, len(g.adjList))
	color[0] = 1
	q := Queue{array: make([]int, 0)}
	q.Push(0)
	for !q.Empty(){
		u := q.Pop()
		i=
	}
}*/
