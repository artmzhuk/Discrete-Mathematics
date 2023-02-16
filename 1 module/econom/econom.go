package main

import "fmt"

func getOperation(input []rune) []rune {
	openedBrackets := 1
	var i int
	for i = 0; openedBrackets != 0 && i < len(input); i++ {
		if input[i] == '(' {
			openedBrackets++
		}
		if input[i] == ')' {
			openedBrackets--
		}
	}
	return input[:i-1]
}

func econom(input string) int {
	inputRune := []rune(input)
	operations := make(map[string]bool, 1)
	for i := 0; i < len(inputRune); i++ {
		if inputRune[i] == '(' {
			operation := getOperation(inputRune[i+1:])
			operations[string(operation)] = true
		}
	}
	return len(operations)
}

func main() {
	var input string
	_, _ = fmt.Scan(&input)
	fmt.Println(econom(input))
}
