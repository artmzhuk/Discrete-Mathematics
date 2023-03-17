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

func getInput() Graph {
	var n int
	fmt.Scan(&n)
	edges := make([][]Node, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			currentNode := Node{
				dest:   i*n + j,
				weight: 0,
			}
			fmt.Scan(&currentNode.weight)
			if i > 0 {
				edges[(i-1)*n+j] = append(edges[(i-1)*n+j], currentNode)
			}
			if i < n-1 {
				edges[(i+1)*n+j] = append(edges[(i+1)*n+j], currentNode)
			}
			if j > 0 {
				edges[i*n+j-1] = append(edges[i*n+j-1], currentNode)
			}
			if j < n-1 {
				edges[i*n+j+1] = append(edges[i*n+j+1], currentNode)
			}
		}
	}
	g := Graph{
		adjList: edges,
		nodeN:   n * n,
	}
	return g
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

func isEmpty(heap *MinHeap) bool { return heap.size == 0 }

func extractMin(heap *MinHeap) MinHeapNode {
	root := heap.array[0]
	last := heap.array[heap.size-1]
	heap.pos[root.v] = heap.size - 1
	heap.pos[last.v] = 0

	heap.array[0] = last
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

func isInMinHeap(heap *MinHeap, v int) bool { return heap.pos[v] < heap.size }

func mst(g Graph) int {
	res := 0
	V := g.nodeN
	parent := make([]int, V)
	key := make([]int, V)
	heap := MinHeap{
		size:     V,
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
				res += key[v]
				decreaseKey(&heap, v, key[v])
			}
		}
	}
	return res
}

func main() {
	fmt.Println("Hi!")
	getInput()
}
