package main

import "fmt"

type Queue struct {
	queue []int
}

func (a *Queue) Push(x int) {
	a.queue = append(a.queue, x)
}

func (a *Queue) Pop(x int) int {
	res := a.queue[0]
	a.queue = a.queue[1:]
	return res
}

func getInput() ([][]int, []int) {
	var N, M, K int
	fmt.Scan(&N, &M)
	edges := make([][]int, N)
	for i := 0; i < M; i++ {
		var from, to int
		fmt.Scan(&from, &to)
		edges[from] = append(edges[from], to)
		edges[to] = append(edges[to], from)
	}
	fmt.Scan(&K)
	bases := make([]int, K)
	for i := 0; i < K; i++ {
		var base int
		fmt.Scan(&base)
		bases[i] = base
	}
	return edges, bases
}

func main() {
	edges, bases := getInput()
	fmt.Println(edges)
	fmt.Println(bases)
	a := make([]int, 3)
	a[0] = 0
	a[1] = 1
	a[2] = 2
	fmt.Println(a)
	a = a[1:]
	fmt.Println(a)
}
