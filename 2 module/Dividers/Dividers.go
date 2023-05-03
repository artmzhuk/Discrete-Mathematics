package main

import (
	"fmt"
	"os"
)

func printDividersGraph(x int) {

	file, err := os.Create("D:\\Projects\\Go\\Discrete-Mathematics\\2 module\\Dividers\\out.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//writer := bufio.NewWriter(file)

	if x == 1 {
		fmt.Fprintf(file, "graph G {\n\t1\n}")
		return
	}

	fmt.Fprintf(file, "graph G {\n\t\\n}")
	xDividers2 := make([]int, 0, 1)
	xDividers := make([]int, 0, 1)
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			xDividers2 = append(xDividers2, i)
			if x/i != i {
				xDividers = append(xDividers, x/i)
			}
		}
	}
	for i := 0; i < len(xDividers2); i++ {
		xDividers = append(xDividers, xDividers2[len(xDividers2)-1-i])
	}
	fmt.Fprintf(file, "\n")
	for i := 0; i < len(xDividers); i++ {
		alreadyInGraph := make([]int, 0, 1)
		for j := i + 1; j < len(xDividers); j++ {
			if xDividers[i]%xDividers[j] == 0 {
				alreadyInGraph = append(alreadyInGraph, xDividers[j])
				isToPrint := true
				for k := 0; k < len(alreadyInGraph)-1; k++ {
					if alreadyInGraph[k]%xDividers[j] == 0 {
						isToPrint = false
						break
					}
				}
				if isToPrint {
					fmt.Println(xDividers[i], "--", xDividers[j])
					fmt.Fprintf(file, "\\t%d -- %d\\n", xDividers[i], xDividers[j])
				}
			}
		}
	}
	fmt.Fprintf(file, "}")
}

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	printDividersGraph(x)
}
