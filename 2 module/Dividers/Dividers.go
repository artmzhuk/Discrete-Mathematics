package main

import (
	"fmt"
	"math"
)

func printDividersGraph(x int) {
	if x == 1 {
		fmt.Println("graph G {\n\t1\n}")
		return
	}

	xDividers2 := make([]int, 0, 1)
	xDividers := make([]int, 0, 1)
	for i := 1; i <= int(math.Sqrt(float64(x))); i++ {
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
	fmt.Println("graph G {")
	for i := 0; i < len(xDividers); i++ {
		alreadyInGraph := make([]int, 0, 1)
		for j := i + 1; j < len(xDividers); j++ {
			if xDividers[i]%xDividers[j] == 0 {
				alreadyInGraph = append(alreadyInGraph, xDividers[j])
				isToPrint := true
				for k := 0; k < len(alreadyInGraph)-1; k++ {
					if alreadyInGraph[k]%xDividers[j] == 0 {
						isToPrint = false
					}
				}
				if isToPrint {
					fmt.Println("\t", xDividers[i], "--", xDividers[j])
				}
			}
		}
	}
	fmt.Println("}")
}

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	printDividersGraph(x)
}
