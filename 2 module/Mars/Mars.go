package main

import (
	"fmt"
	"math"
	"sort"
)

type Graph struct {
	adjList [][]int
}

type solutions struct {
	colors [][]int
}

type colorSplit struct {
	color1 []int
	color2 []int
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

func checkIsBipartiteAndSplit(g *Graph) solutions {
	visited := make([]bool, len(g.adjList))
	possibleSolutions := make([][]int, 0)
	isBipartite := true
	for i := 0; i < len(g.adjList); i++ {
		if visited[i] {
			continue
		}
		isBipartite = true
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
		} else {
			break
		}
	}
	if len(possibleSolutions) == 0 || !isBipartite {
		return solutions{
			colors: nil,
		}
	} else {
		return solutions{
			colors: possibleSolutions,
		}
	}
}

func selectSolution(s solutions) {
	possible := make([]colorSplit, len(s.colors))
	for i := range possible {
		possible[i].color1 = make([]int, 0)
		possible[i].color2 = make([]int, 0)
		for j := range s.colors[i] {
			if s.colors[i][j] == 1 {
				possible[i].color1 = append(possible[i].color1, j)
			} else if s.colors[i][j] == 2 {
				possible[i].color2 = append(possible[i].color2, j)
			}
		}
	}

	ultraNode := make([][]int, 1)
	for i := len(possible) - 1; i >= 0; i-- {
		if len(possible[i].color1) == 1 && len(possible[i].color2) == 0 {
			ultraNode[0] = append(ultraNode[0], possible[i].color1[0])
			possible[i].color1 = possible[i].color1[:0]
			possibleTmp := possible[:i]
			possibleTmp2 := possible[i+1:]
			possible = possible[0:0]
			possible = append(possible, append(possibleTmp, possibleTmp2...)...)
		}
	}
	if len(ultraNode[0]) > 0 {
		for i := 1; i < int(math.Pow(2, float64(len(ultraNode[0]))))-1; i++ {
			toAppend := make([]int, 0)
			for j := 0; j < len(ultraNode[0]); j++ {
				if (i>>j)&1 == 1 {
					toAppend = append(toAppend, ultraNode[0][j])
				}
			}
			ultraNode = append(ultraNode, toAppend)
		}
	}

	possible2 := make([][]int, 0)
	prev := make([][]int, 0)
	if len(possible) > 0 {
		recursiveGenerationOfSolutions(&possible2, &possible, prev, 0, 0)
		recursiveGenerationOfSolutions(&possible2, &possible, prev, 0, 1)
		for i := range possible2 {
			for j := range ultraNode {
				if len(possible2[i]) > 0 && len(ultraNode[0]) > 0 {
					toCopy := make([]int, len(possible2[i]))
					possible2 = append(possible2, toCopy)
					copy(possible2[len(possible2)-1], possible2[i])
					possible2[len(possible2)-1] = append(possible2[len(possible2)-1], ultraNode[j]...)
				}
			}
		}
	} else {
		for i := range ultraNode {
			possible2 = append(possible2, nil)
			possible2[i] = append(possible2[i], ultraNode[i]...)
		}
	}
	for k := range possible2 {
		sort.Slice(possible2[k], func(i, j int) bool {
			return possible2[k][i] < possible2[k][j]
		})
	}

	for i2 := len(possible2) - 1; i2 >= 0; i2-- {
		toAppend := make([]int, 0)
		j2 := 0
		for k := 0; j2 < len(s.colors[0]) || k < len(possible2[i2]); {
			if k >= len(possible2[i2]) || possible2[i2][k] != j2 {
				toAppend = append(toAppend, j2)
			} else {
				k++
			}
			j2++
		}
		possible2 = append(possible2, toAppend)
	}

	possible3 := make([][]int, 0)
	for diff := 0; diff < len(s.colors[0]); diff++ {
		for i := range possible2 {
			if len(possible2[i]) == len(s.colors[0])/2-diff || len(possible2[i]) == len(s.colors[0])/2+diff {
				possible3 = append(possible3, possible2[i])
			}
		}
		if len(possible3) > 0 {
			break
		}
	}
	if len(possible3) == 1 {
		sort.Slice(possible3[0], func(i, j int) bool { return possible3[0][i] < possible3[0][j] })
		for i := range possible3[0] {
			fmt.Printf("%d ", possible3[0][i]+1)
		}
		return
	}
	for k := range possible3 {
		sort.Slice(possible3[k], func(i, j int) bool {
			return possible3[k][i] < possible3[k][j]
		})
	}
	sort.Slice(possible3, func(i, j int) bool {
		if len(possible3[i]) < len(possible3[j]) {
			return true
		} else if len(possible3[i]) > len(possible3[j]) {
			return false
		} else {
			for k := range possible3[i] {
				if possible3[i][k] < possible3[j][k] {
					return true
				} else if possible3[i][k] > possible3[j][k] {
					return false
				}
			}
		}
		return true
	})
	for i := range possible3[0] {
		fmt.Printf("%d ", possible3[0][i]+1)
	}
}

func recursiveGenerationOfSolutions(array *[][]int, parent *[]colorSplit, prev [][]int, parent1, parent2 int) {
	if parent2 == 0 {
		prev = append(prev, (*parent)[parent1].color1)
	} else {
		prev = append(prev, (*parent)[parent1].color2)
	}
	if parent1+1 < len(*parent) {
		recursiveGenerationOfSolutions(array, parent, prev, parent1+1, 0)
		if len((*parent)[parent1+1].color2) != 0 {
			recursiveGenerationOfSolutions(array, parent, prev, parent1+1, 1)
		}
	} else {
		*array = append(*array, nil)
		for i := range prev {
			(*array)[len(*array)-1] = append((*array)[len(*array)-1], prev[i]...)
		}
	}
}

func main() {
	g := inputGraph()
	splittedGraph := checkIsBipartiteAndSplit(g)
	if len(splittedGraph.colors) > 0 {
		selectSolution(splittedGraph)
	} else {
		fmt.Println("No solution")
	}
}
