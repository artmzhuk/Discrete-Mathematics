package main

import "fmt"

type Machine struct {
	stateNum   int //number of machine's states
	inputsNum  int //number of possible inputs
	startState int //number of start state
	stateTrans [][]int
	outTrans   [][]string
}

func getInput() Machine {
	var n, m, q0 int
	fmt.Scan(&n, &m, &q0)

	states := make([][]int, n)
	out := make([][]string, n)
	for i := range states {
		states[i] = make([]int, m)
		for j := range states[i] {
			fmt.Scan(&states[i][j])
		}
	}
	for i := range out {
		out[i] = make([]string, m)
		for j := range out[i] {
			fmt.Scan(&out[i][j])
		}
	}

	machine := Machine{
		stateNum:   n,
		inputsNum:  m,
		startState: q0,
		stateTrans: states,
		outTrans:   out,
	}
	return machine
}

func createGraphVizFromMachine(m Machine) {
	fmt.Print("digraph {\n\trankdir = LR")
	for i := 0; i < m.stateNum; i++ {
		for j := 0; j < m.inputsNum; j++ {
			fmt.Printf("\n    %d -> %d [label = \"%c(%s)\"]",
				i, m.stateTrans[i][j], 'a'+j, m.outTrans[i][j])
		}
	}
	fmt.Print("\n}")
}

func main() {
	machine := getInput()
	createGraphVizFromMachine(machine)
}
