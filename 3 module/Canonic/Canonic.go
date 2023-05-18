package main

import (
	"bufio"
	"fmt"
	"os"
)

type transition struct {
	dest int
	in   string
	out  string
}

type Machine struct {
	stateNum   int //number of machine's states
	inputsNum  int //number of possible inputs
	startState int //number of start state
	adjList    [][]transition
	visited    []bool
	canonical  []int
}

func getInput() Machine {
	reader := bufio.NewReader(os.Stdin)
	var n, m, q0 int
	fmt.Scan(&n, &m, &q0)

	adjList := make([][]transition, n)
	for i := range adjList {
		adjList[i] = make([]transition, m)
		for j := 0; j < len(adjList[i]); j++ {
			fmt.Fscan(reader, &adjList[i][j].dest)
		}
	}
	for i := range adjList {
		for j := 0; j < len(adjList[i]); j++ {
			fmt.Fscan(reader, &adjList[i][j].out)
			adjList[i][j].in = string(rune('a' + j))
		}
	}
	visited := make([]bool, n)
	canonical := make([]int, n)
	machine := Machine{
		stateNum:   n,
		inputsNum:  m,
		startState: q0,
		adjList:    adjList,
		visited:    visited,
		canonical:  canonical,
	}
	return machine
}

func dfs(m *Machine, at int, counter *int) {
	if counter == nil {
		x := 0
		counter = &x
	}
	m.visited[at] = true
	m.canonical[at] = *counter
	*counter++
	for i := range m.adjList[at] {
		next := m.adjList[at][i].dest
		if !m.visited[next] {
			dfs(m, next, counter)
		}
	}
}

func printOutput(m *Machine) {
	newAdjList := make([][]transition, m.stateNum)
	for i := range newAdjList {
		newAdjList[i] = make([]transition, m.inputsNum)
	}
	for i := range m.canonical {
		canonicOfI := m.canonical[i]
		copy(newAdjList[canonicOfI], m.adjList[i])
		for j := range newAdjList[i] {
			newAdjList[canonicOfI][j].dest = m.canonical[newAdjList[canonicOfI][j].dest]
		}
	}
	fmt.Println(m.stateNum)
	fmt.Println(m.inputsNum)
	fmt.Println("0")
	for i := range newAdjList {
		for j := range newAdjList[i] {
			fmt.Print(newAdjList[i][j].dest, " ")
		}
		fmt.Println()
	}
	for i := range newAdjList {
		for j := range newAdjList[i] {
			fmt.Print(newAdjList[i][j].out, " ")
		}
		fmt.Println()
	}
}

func main() {
	m := getInput()
	dfs(&m, m.startState, nil)
	printOutput(&m)
}
