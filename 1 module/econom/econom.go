package main

import "fmt"

func ops(input string) int {
	inputRune := []rune(input)
	startIndexes := make([]int, 0, len(inputRune)/2) // stack of '(' indexes in inputRune slice
	expressions := make(map[string]bool, 1)
	for i := 0; i < len(inputRune); i++ {
		if inputRune[i] == '(' && (inputRune[i+1] == '@' ||
			inputRune[i+1] == '#' ||
			inputRune[i+1] == '$') {
			startIndexes = append(startIndexes, i)
		}
		if inputRune[i] == ')' {
			index := startIndexes[len(startIndexes)-1] // pops '(' index from stack,
			startIndexes = startIndexes[:len(startIndexes)-1]
			expressions[string(inputRune[index:i+1])] = true // adds "()" expression to the "expressions" map
		}
	}
	return len(expressions)
}

func main() {
	var input string
	_, _ = fmt.Scan(&input)
	fmt.Println(ops(input))
}
