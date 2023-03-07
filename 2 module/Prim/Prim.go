package main

import "fmt"

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

func swapMinHeapNode(a, b int, heap *MinHeap) {
	heap.pos
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
	if isEmpty(heap) {
		return nil
	}
	heap.array[0] = heap.array[(len(heap.array) - 1)]
	heap.array = heap.array[:len(heap.array)-1]
	return heap.array[0]
}

/*func decreaseKey(heap *MinHeap, v,key int){
	heap.array[v].key = key
	while()
}*/

func mst(g Graph) {
	V := g.nodeN
	parent := make([]int, V)
	key := make([]int, V)
	heap := MinHeap{
		size:     0,
		capacity: V,
		pos:      nil,
		array:    make([]MinHeapNode, V),
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

}
