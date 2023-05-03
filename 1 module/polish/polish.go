package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	result := polishNotation(line)
	fmt.Println(result)
}

func pop2(stack []int) ([]int, int, int) {
	a := stack[len(stack)-1]
	b := stack[len(stack)-2]
	stack = stack[:len(stack)-2]
	return stack, a, b
}

func polishNotation(input string) int {
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")
	inputRune := ([]rune)(input)
	length := len(inputRune)
	for i := 0; i < length/2; i++ { // reverses input string
		inputRune[i], inputRune[length-i-1] = inputRune[length-i-1], inputRune[i]
	}
	stack := make([]int, 0, 10)

	for i := 0; i < length; i++ {
		switch {
		case inputRune[i] < '(' || inputRune[i] > '9':
			continue
		case inputRune[i] >= '0' && inputRune[i] <= '9':
			number, _ := strconv.Atoi(string(inputRune[i]))
			stack = append(stack, number)
		case inputRune[i] == '+':
			var a, b int
			stack, a, b = pop2(stack)
			stack = append(stack, a+b)
		case inputRune[i] == '-':
			var a, b int
			stack, a, b = pop2(stack)
			stack = append(stack, a-b)
		case inputRune[i] == '*':
			var a, b int
			stack, a, b = pop2(stack)
			stack = append(stack, a*b)
		case inputRune[i] == '/':
			var a, b int
			stack, a, b = pop2(stack)
			stack = append(stack, a/b)
		}
	}
	result := stack[0]
	return result
}
