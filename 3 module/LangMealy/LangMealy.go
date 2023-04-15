package main

import (
	"fmt"
	"sort"
)

type Machine struct {
	statesNumber int //number of machine's states
	maxWordLen   int //number of possible inputs
	startState   int //number of start state
	stateMatrix  [][]int
	outputMatrix [][]rune
}

func createMachineFromInput() Machine {
	var n, m, q0 int

	fmt.Scan(&n)

	states := make([][]int, n)
	out := make([][]rune, n)

	for i := range states {
		states[i] = make([]int, 2)
		for j := range states[i] {
			fmt.Scan(&states[i][j])
		}
	}
	for i := range out {
		out[i] = make([]rune, 2)
		for j := range out[i] {
			var scanned string
			fmt.Scan(&scanned)
			out[i][j] = ([]rune(scanned))[0]
		}
	}
	fmt.Scan(&q0, &m)
	machine := Machine{
		statesNumber: n,
		maxWordLen:   m,
		startState:   q0,
		stateMatrix:  states,
		outputMatrix: out,
	}
	return machine
}

func DFS(machine Machine, index int, wordsArray *[]string, currentWord string, previousUsagesN *[]int) {
	if len(currentWord) < machine.maxWordLen {
		for i := range machine.stateMatrix[index] { //goes through neighbours of node
			currentWordChanged := false
			if machine.outputMatrix[index][i] != '-' {
				currentWord += string(machine.outputMatrix[index][i])
				currentWordChanged = true
			}
			/* this condition checks that number of letters in current word changed since previous
			   visit of node (in order to avoid infinite loop on "lambda" transition)*/
			if len(currentWord) != (*previousUsagesN)[machine.stateMatrix[index][i]] &&
				!(machine.outputMatrix[index][i] == '-' && index == machine.stateMatrix[index][i]) {

				(*previousUsagesN)[index]++
				DFS(machine, machine.stateMatrix[index][i], wordsArray, currentWord, previousUsagesN)
			}
			if currentWordChanged {
				*wordsArray = append(*wordsArray, currentWord)
				currentWord = string([]rune(currentWord)[:len([]rune(currentWord))-1])
			}
		}
	}
}

func startDFS(machine Machine) []string {
	result := make([]string, 0)
	prev := make([]int, machine.statesNumber)
	for i := range prev {
		prev[i] = -1
	}
	DFS(machine, 0, &result, "", &prev)
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	for i := range result {
		if i > 0 && result[i-1] == result[i] {
			result[i-1] = ""
		}
	}
	return result
}

func main() {
	result := startDFS(createMachineFromInput())
	for i := range result {
		if result[i] != "" {
			fmt.Print(result[i] + " ")
		}
	}
}
