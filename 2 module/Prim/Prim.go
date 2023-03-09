package main

import (
	"fmt"
	"math"
)

type Graph struct {
	adjList [][]Node
	nodeN   int
}

type Node struct {
	dest, weight int
}

type MinHeapNode struct {
	v, key int
}

type MinHeap struct {
	size, capacity int
	pos            []int
	array          []MinHeapNode
}

func newMinHeapNode(v, key int) MinHeapNode {
	res := MinHeapNode{
		v:   v,
		key: key,
	}
	return res
}

func minHeapify(minHeap *MinHeap, idx int) {
	smallest := idx
	left := 2*idx + 1
	right := 2*idx + 2
	if left < minHeap.size &&
		minHeap.array[left].key < minHeap.array[smallest].key {
		smallest = left
	}
	if right < minHeap.size &&
		minHeap.array[right].key < minHeap.array[smallest].key {
		smallest = left
	}
	if smallest != idx {
		minHeap.pos[minHeap.array[smallest].v] = idx
		minHeap.pos[minHeap.array[idx].v] = smallest
		minHeap.array[smallest], minHeap.array[idx] =
			minHeap.array[idx], minHeap.array[smallest]
		minHeapify(minHeap, smallest)
	}
}

func isEmpty(heap *MinHeap) bool {
	if heap.size == 0 {
		return true
	} else {
		return false
	}
}

func extractMin(heap *MinHeap) MinHeapNode {
	root := heap.array[0]
	last := heap.array[heap.size-1]
	heap.pos[root.v] = heap.size - 1
	heap.pos[last.v] = 0

	heap.array[0] = last
	//heap.array[len(heap.array)-1] = root
	heap.size--
	minHeapify(heap, 0)
	return root
}

func decreaseKey(heap *MinHeap, v, key int) {
	i := heap.pos[v]
	heap.array[i].key = key
	for i != 0 && heap.array[i].key < heap.array[(i-1)/2].key {
		heap.pos[heap.array[i].v] = (i - 1) / 2
		heap.pos[heap.array[(i-1)/2].v] = i
		heap.array[i], heap.array[(i-1)/2] = heap.array[(i-1)/2], heap.array[i]
		i = (i - 1) / 2
	}
}

func isInMinHeap(heap *MinHeap, v int) bool {
	if heap.pos[v] < heap.size {
		return true
	}
	return false
}

func mst(g Graph) {
	res := 0
	V := g.nodeN
	parent := make([]int, V)
	key := make([]int, V)
	heap := MinHeap{
		size:     0,
		capacity: V,
		pos:      make([]int, V),
		array:    make([]MinHeapNode, V),
	}
	for v := 1; v < V; v++ {
		parent[v] = -1
		key[v] = math.MaxInt
		heap.array[v] = newMinHeapNode(v, key[v])
		heap.pos[v] = v
	}
	//key[0] = 0
	heap.array[0] = newMinHeapNode(0, key[0])
	heap.pos[0] = 0
	heap.size = V
	for !isEmpty(&heap) {
		minHeapNode := extractMin(&heap)
		u := minHeapNode.v
		for i := 0; i < len(g.adjList[u]); i++ {
			pCrawl := g.adjList[u][i]
			v := g.adjList[u][i].dest
			if isInMinHeap(&heap, v) && pCrawl.weight < key[v] {
				if key[v] != math.MaxInt {
					res -= key[v]
				}
				key[v] = pCrawl.weight
				parent[v] = u
				res += pCrawl.weight
				decreaseKey(&heap, v, key[v])
			}
		}
	}
}

func getNodes() Graph {
	var n, m int
	fmt.Scan(&n, &m)
	nodes := make([][]Node, n)
	for i := 0; i < m; i++ {
		var from, to, weight int
		fmt.Scan(&from, &to, &weight)
		node1 := Node{dest: to, weight: weight}
		node2 := Node{dest: from, weight: weight}
		nodes[from] = append(nodes[from], node1)
		nodes[to] = append(nodes[to], node2)
	}
	g := Graph{
		adjList: nodes,
		nodeN:   n,
	}
	return g
}

func main() {
	g := getNodes()
	mst(g)
}
