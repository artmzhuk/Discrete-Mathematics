/*priority queue from golang docs
https://pkg.go.dev/container/heap#example-package-PriorityQueue
*/

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Node, value int, priority int) {
	item.dest = value
	item.weight = priority
	heap.Fix(pq, item.index)
}

type Graph struct {
	adjList [][]Node
	nodeN   int
}

type Node struct {
	dest, weight, index int
}

func getNodes() Graph {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	nodes := make([][]Node, n)
	for i := 0; i < m; i++ {
		var from, to, weight int
		_, _ = fmt.Scan(&from, &to, &weight)
		node1 := Node{dest: to, weight: weight}
		node2 := Node{dest: from, weight: weight}
		nodes[from] = append(nodes[from], node1)
		nodes[to] = append(nodes[to], node2)
	}
	g := Graph{adjList: nodes, nodeN: n}
	return g
}

func mst(g Graph) int {
	pq := make(PriorityQueue, 1)
	sum := 0
	src := 0
	key := make([]int, len(g.adjList))
	inMST := make([]bool, len(g.adjList))
	for i := range key {
		key[i] = math.MaxInt
		inMST[i] = false
	}
	pq[0] = &Node{
		dest:   src,
		weight: 0,
	}
	heap.Init(&pq)
	key[src] = 0

	for pq.Len() > 0 {
		u := pq[0].dest
		heap.Pop(&pq)
		if !inMST[u] {
			for i := range g.adjList[u] {
				v := g.adjList[u][i].dest
				weight := g.adjList[u][i].weight
				if !inMST[v] && key[v] > weight {
					key[v] = weight
					heap.Push(&pq, &Node{
						dest:   v,
						weight: key[v],
					})
				}
			}
			inMST[u] = true
		}
	}
	for i := range key {
		sum += key[i]
	}
	return sum
}

func main() {
	g := getNodes()
	fmt.Println(mst(g))
}
