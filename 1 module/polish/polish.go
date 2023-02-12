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
			numberLength := 1
			for ; inputRune[i+numberLength] >= '0' && inputRune[i+numberLength] <= '9'; numberLength++ {
			} // calculates number length
			for k := 0; k < numberLength/2; k++ {
				inputRune[i+k], inputRune[i+numberLength-1-k] = inputRune[i+numberLength-1-k], inputRune[i+k]
			} //reverses number to its normal state
			number, _ := strconv.Atoi(string(inputRune[i : i+numberLength]))
			i += numberLength - 1
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
