package main

import "fmt"

type Machine struct {
	stateNum   int //number of machine's states
	inputsNum  int //number of possible inputs
	startState int //number of start state
	stateTrans [][]int
	outTrans   [][]rune
}

func getInput() Machine {
	var n, m, q0 int
	fmt.Scan(&n, &m, &q0)

	states := make([][]int, n)
	out := make([][]rune, n)
	for i := range states {
		states[i] = make([]int, m)
		for j := range states[i] {
			fmt.Scan(&states[i][j])
		}
	}
	for i := range out {
		out[i] = make([]rune, m)
		for j := range out[i] {
			var scanned string
			fmt.Scan(&scanned)
			out[i][j] = ([]rune(scanned))[0]
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

func createGraphVizFromMachine(m Machine) string {
	result := "digraph {\n\trankdir = LR"
	for i := 0; i < m.stateNum; i++ {
		for j := 0; j < m.inputsNum; j++ {
			result += fmt.Sprintf("\n\t%d -> %d [label = \"%c(%c)\"]",
				i, m.stateTrans[i][j], 'a'+j, m.outTrans[i][j])
		}
	}
	result += "\n}"
	return result
}

func main() {
	machine := getInput()
	res := createGraphVizFromMachine(machine)
	fmt.Println(res)
}
