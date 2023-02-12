package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func polishNotation(input string) int {
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")
	inputRune := ([]rune)(input)
	length := len(inputRune)
	for i := 0; i < length/2; i++ {
		inputRune[i], inputRune[length-i-1] = inputRune[length-i-1], inputRune[i]
	}
	stack := make([]int, 0, 10)

	for i := 0; i < length; i++ {
		switch {
		case inputRune[i] >= '0' && inputRune[i] <= '9':
			stack = append(stack, int(inputRune[i]-'0'))
		case inputRune[i] == '+':
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case inputRune[i] == '-':
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case inputRune[i] == '*':
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case inputRune[i] == '/':
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		}
	}
	result := stack[0]
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	result := polishNotation(line)
	fmt.Println(result)
}
