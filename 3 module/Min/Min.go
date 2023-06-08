/* реализация минимизации взята из лекций по ДМ,
реализация лес непересекающихся мн-в -- из лекций по АиСД
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

type transition struct {
	dest int //bi
	in   string
	out  string //phi
}

type state struct {
	parent int
	depth  int
}

type Machine struct {
	Q          int //number of machine's states
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
		Q:          n,
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

func canonic(m *Machine) {
	dfs(m, m.startState, nil)
	canonizedQ := 0
	for i := range m.visited {
		if m.visited[i] {
			canonizedQ++
		}
	}
	m.Q = canonizedQ
	newAdjList := make([][]transition, m.Q)
	for i := range newAdjList {
		newAdjList[i] = make([]transition, m.inputsNum)
	}
	for i := range m.canonical {
		if m.visited[i] {
			canonicOfI := m.canonical[i]
			copy(newAdjList[canonicOfI], m.adjList[i])
			for j := range newAdjList[canonicOfI] {
				newAdjList[canonicOfI][j].dest = m.canonical[newAdjList[canonicOfI][j].dest]
			}
		}
	}
	m.adjList = make([][]transition, len(newAdjList))
	m.startState = 0
	copy(m.adjList, newAdjList)
}

func printOutput(m *Machine) {
	fmt.Print("digraph {\n    rankdir = LR")
	for i := 0; i < m.Q; i++ {
		for j := 0; j < m.inputsNum; j++ {
			if len(m.adjList[i][j].out) != 0 {
				fmt.Printf("\n    %d -> %d [label = \"%c(%s)\"]",
					i, m.adjList[i][j].dest, 'a'+j, m.adjList[i][j].out)
			}
		}
	}
	fmt.Print("\n}")
}

func find(set *[]state, x int) int {
	if (*set)[x].parent == x {
		return x
	} else {
		(*set)[x].parent = find(set, (*set)[x].parent)
		return (*set)[x].parent
	}
}

func union(set *[]state, x, y int) {
	rootX := find(set, x)
	rootY := find(set, y)
	if (*set)[rootX].depth < (*set)[y].depth {
		(*set)[rootX].parent = rootY
	} else {
		(*set)[rootY].parent = (*set)[rootX].parent
		if (*set)[rootX].depth == (*set)[rootY].depth && rootX != rootY {
			(*set)[rootX].depth += 1
		}
	}
}

func split1(mach *Machine) (m int, p []int) {
	m = mach.Q
	set := make([]state, m)
	p = make([]int, m)
	for q := range set {
		set[q].parent = q
		set[q].depth = 0
	}
	for q1 := 0; q1 < mach.Q; q1++ {
		for q2 := q1 + 1; q2 < mach.Q; q2++ {
			if find(&set, q1) != find(&set, q2) {
				eq := true
				for l := 0; l < mach.inputsNum; l++ {
					if mach.adjList[q1][l].out != mach.adjList[q2][l].out {
						eq = false
						break
					}
				}
				if eq {
					union(&set, q1, q2)
					m -= 1
				}
			}
		}
	}
	for q := range set {
		p[q] = find(&set, q)
	}
	return m, p
}

func split(mach *Machine, p *[]int) int {
	m := mach.Q
	set := make([]state, m)
	for q := range set {
		set[q].parent = q
		set[q].depth = 0
	}
	for q1 := 0; q1 < mach.Q; q1++ {
		for q2 := q1 + 1; q2 < mach.Q; q2++ {
			if (*p)[q1] == (*p)[q2] && find(&set, q1) != find(&set, q2) {
				eq := true
				for x := 0; x < mach.inputsNum; x++ {
					w1 := mach.adjList[q1][x].dest
					w2 := mach.adjList[q2][x].dest
					if (*p)[w1] != (*p)[w2] {
						eq = false
						break
					}
				}
				if eq {
					union(&set, q1, q2)
					m -= 1
				}
			}
		}
	}
	for q := range set {
		(*p)[q] = find(&set, q)
	}
	return m
}

func AufenkampHohn(A *Machine) Machine {
	m, p := split1(A)
	for true {
		m1 := split(A, &p)
		if m == m1 {
			break
		}
		m = m1
	}
	alreadyInA2 := make(map[int]bool)
	adjList1 := make([][]transition, A.Q)
	for i := range adjList1 {
		adjList1[i] = make([]transition, A.inputsNum)
	}
	visited := make([]bool, A.Q)
	canonical := make([]int, A.Q)
	A1 := Machine{
		Q:          A.Q,
		inputsNum:  A.inputsNum,
		startState: A.startState,
		adjList:    adjList1,
		visited:    visited,
		canonical:  canonical,
	}
	for q := range A.adjList {
		q1 := p[q]
		i, _ := alreadyInA2[q1]
		if !i {
			for j := range A1.adjList[q] {
				A1.adjList[q][j].dest = p[A.adjList[q][j].dest]
				A1.adjList[q][j].out = A.adjList[q][j].out
			}
			alreadyInA2[q1] = true
		}
	}
	return A1
}

func main() {
	m := getInput()
	canonic(&m)
	m2 := AufenkampHohn(&m)
	canonic(&m2)
	printOutput(&m2)
}
