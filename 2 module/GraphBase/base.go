package main

import (
	"fmt"
	"sort"
)

func getInput() ([][]int, [][]int) {
	var n, m int
	fmt.Scan(&n, &m)
	graph := make([][]int, n)
	graphTr := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		graph[a] = append(graph[a], b)
		graphTr[b] = append(graphTr[b], a)
	}
	return graph, graphTr
}

func dfs1(v int, g *[][]int, used *[]bool, order *[]int) {
	(*used)[v] = true
	for i := range (*g)[v] {
		if !(*used)[(*g)[v][i]] {
			dfs1((*g)[v][i], g, used, order)
		}
	}
	*order = append(*order, v)
}

func dfs2(v int, gr *[][]int, used *[]bool, component *[]int) {
	(*used)[v] = true
	*component = append(*component, v)
	for i := range (*gr)[v] {
		if !(*used)[(*gr)[v][i]] {
			dfs2((*gr)[v][i], gr, used, component)
		}
	}
}

func buildCond(g, gr [][]int) []int {
	n := len(g)
	used := make([]bool, n)
	order := make([]int, 0)
	component := make([]int, 0)
	components := make([][]int, 0)
	for i := 0; i < n; i++ {
		if !used[i] {
			dfs1(i, &g, &used, &order)
		}
	}
	for i := range used {
		used[i] = false
	}
	for i := 0; i < n; i++ {
		v := order[n-1-i]
		if !used[v] {
			dfs2(v, &gr, &used, &component)
			components = append(components, component)
			component = nil
		}
	}
	nodesByComponent := make([]int, n)
	for i := range components {
		for j := range components[i] {
			nodesByComponent[components[i][j]] = i
		}
	}
	condensationGraphTr := make([][]int, len(components))
	for i := range components {
		for j := range components[i] {
			for k := range g[components[i][j]] {
				if nodesByComponent[g[components[i][j]][k]] != i {
					condensationGraphTr[nodesByComponent[g[components[i][j]][k]]] =
						append(condensationGraphTr[nodesByComponent[g[components[i][j]][k]]], i)
				}
			}
		}
	}
	result := make([]int, 0)
	for k := range condensationGraphTr {
		if condensationGraphTr[k] == nil {
			sort.Ints(components[k])
			result = append(result, components[k][0])
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	g, gr := getInput()
	res := buildCond(g, gr)
	for i := range res {
		fmt.Print(res[i], " ")
	}
}
